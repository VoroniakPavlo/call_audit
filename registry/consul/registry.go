package consul

import (
	"fmt"
	"log/slog"
	"net"
	"strconv"
	"time"

	consulapi "github.com/hashicorp/consul/api"
	conf "github.com/VoroniakPavlo/call_audit/config"
	cerror "github.com/VoroniakPavlo/call_audit/internal/errors"
	"github.com/VoroniakPavlo/call_audit/registry"
)

type ConsulRegistry struct {
	registrationConfig *consulapi.AgentServiceRegistration
	client             *consulapi.Client
	stop               chan any
	checkId            string
}

// NewConsulRegistry creates a new Consul registry instance.
func NewConsulRegistry(config *conf.ConsulConfig) (*ConsulRegistry, error) {
	var err error
	entity := ConsulRegistry{}
	if config.Id == "" {
		return nil, cerror.NewInternalError("consul.registry.new_consul.check_args.service_id", "service id is empty! (set it by '-id' flag)")
	}
	ip, port, err := net.SplitHostPort(config.PublicAddress)
	if err != nil {
		return nil, cerror.NewInternalError("consul.registry.new_consul.parse_address.error", "unable to parse address")
	}
	parsedPort, err := strconv.Atoi(port)
	if err != nil {
		return nil, cerror.NewInternalError("consul.registry.new_consul.parse_ip.error", "unable to parse ip")
	}

	consulConfig := consulapi.DefaultConfig()
	consulConfig.Address = config.Address
	entity.client, err = consulapi.NewClient(consulConfig)
	if err != nil {
		return nil, cerror.NewInternalError("consul.registry.new_consul_registry.consulapi_creation.error", err.Error())
	}

	entity.registrationConfig = &consulapi.AgentServiceRegistration{
		ID:      config.Id,
		Name:    registry.ServiceName,
		Port:    parsedPort,
		Address: ip,
		Check: &consulapi.AgentServiceCheck{
			DeregisterCriticalServiceAfter: registry.DeregisterCriticalServiceAfter.String(),
			//CheckID:                        config.Id,
			TTL: registry.CheckInterval.String(),
		},
	}
	entity.stop = make(chan any)

	return &entity, nil
}

// Register registers the service with Consul.
func (c *ConsulRegistry) Register() error {
	err := c.client.Agent().ServiceRegister(c.registrationConfig)
	if err != nil {
		return cerror.NewInternalError("consul.registry.consul.register.error", err.Error())
	}
	var checks map[string]*consulapi.AgentCheck
	if checks, err = c.client.Agent().Checks(); err != nil {
		return cerror.NewInternalError("consul.registry.consul.register.get_checks.error", err.Error())
	}

	var serviceCheck *consulapi.AgentCheck
	for _, check := range checks {
		if check.ServiceID == c.registrationConfig.ID {
			serviceCheck = check
		}
	}

	if serviceCheck == nil {
		return cerror.NewInternalError("consul.registry.consul.register.error", err.Error())
	}
	c.checkId = serviceCheck.CheckID
	go c.RunServiceCheck()
	slog.Info(fmtConsulLog("service was registered"))
	return nil
}

func (c *ConsulRegistry) Deregister() error {
	err := c.client.Agent().ServiceDeregister(c.registrationConfig.ID)
	if err != nil {
		return cerror.NewInternalError("consul.registry.consul.deregister.error", err.Error())
	}
	c.stop <- true
	slog.Info(fmtConsulLog("service was deregistered"))
	return nil
}

func (c *ConsulRegistry) RunServiceCheck() error {
	defer slog.Info(fmtConsulLog("stopped service checker"))
	slog.Info(fmtConsulLog("started service checker"))
	ticker := time.NewTicker(registry.CheckInterval / 2)
	for {
		select {
		case <-c.stop:
			// gracefull stop
			return nil
		case <-ticker.C:
			err := c.client.Agent().UpdateTTL(c.checkId, "success", "pass")
			if err != nil {
				slog.Error(fmtConsulLog(err.Error()))
			}
			// TODO: seems that connection is lost, reconnect?
		}
	}
}

func fmtConsulLog(s string) string {
	return fmt.Sprintf("consul: %s", s)
}
