// TODO WTEL-7091
//go:debug rsa1024min=0

package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand"
	"time"

	"github.com/VoroniakPavlo/call_audit/internal/app"
	"github.com/VoroniakPavlo/call_audit/model"
	"github.com/webitel/storage/pool"
	_ "github.com/webitel/storage/stt"
	_ "github.com/webitel/storage/synchronizer"
	_ "github.com/webitel/storage/uploader"

	"net/http"
	_ "net/http/pprof"
)

type Anc struct {
	pool pool.Pool
}

type anyT struct {
	name int
}

var (
	source = rand.NewSource(time.Now().UnixNano())
	r      = rand.New(source)
)

func (t *anyT) Execute() {
	// id, file
	// tran
	//
	s := time.Duration(r.Intn(3)+1) * time.Second
	fmt.Println(t.name, "Execute: ", s)
	time.Sleep(s)
	fmt.Println(t.name, "Execute complete: ", s)

	// set task state = 3
}

func drop_finished_jobs(app *app.App) {
	_, err := app.Store.ServiceStore().Execute(context.Background(), `with del as (
		    delete from call_audit.jobs where state = 3
		    returning *
		), j as (
		    select rule_id, max(call_stored_at) max_call_stored_at
		    from del
		    group by 1
		)
		update call_center.cc_call_questionnaire_rule r
		    set last_stored_at = max_call_stored_at
		from j
		where r.id = j.rule_id
	`)
	if err != nil {
		slog.Error("Failed to truncate jobs table")
		return
	}
}

func create_jobs(app *app.App, last int, limit int) any {
	_, err := app.Store.ServiceStore().Execute(context.Background(),
		`insert into call_audit.jobs(rule_id, type, params)
		select 1, 2,row_to_json( (h.id, f.id , h.stored_at, row_number() over (order by stored_at) ))
				from call_center.cc_calls_history h
					 inner join lateral (
						 select  f.id
						 from storage.files f
						 where f.domain_id = h.domain_id and f.uuid = h.id::text
						 limit 1
					) f on true
				where h.domain_id = :domain_id
					and h.parent_id isnull
					and stored_at > '2025-07-08 17:13:42.448274 +03:00'
					and h.payload->('vvv') isnull
					and (:call_direction isnull  or h.direction = :call_direction)
				order by stored_at
				limit 100 - :active;`)
	if err != nil {
		slog.Error("Failed to truncate jobs table")
		return err
	}

	return nil
}

func getRules(app *app.App, limit int) ([]model.CallQuestionnaireRule, error) {
	rules, err := app.Store.ServiceStore().Array(context.Background(),
		`select
			coalesce(last_stored_at, 'from') last,
			r.id,
			r.domain_id,
			(select count(*) as l from call_audit.jobs j where j.rule_id = r.id) active
		from call_audit.call_questionnaire_rule r
		where enabled
			and  (select count(*) as l from call_audit.jobs j where j.rule_id = r.id) < :limit
		order by coalesce(last_stored_at, 'from') ;
		`)
	if err != nil {
		slog.Error("Failed to truncate jobs table")
		return nil, err
	}
	if len(rules) == 0 {
		slog.Info("No active rules found")
		return nil, nil
	}
	// process rules

	var processedRules []model.CallQuestionnaireRule
	for _, ruleMap := range rules {
		ruleData, ok := ruleMap.(map[string]interface{})
		if !ok {
			slog.Error("ruleMap is not a map[string]interface{}")
			continue
		}
		lastStr, ok := ruleData["last"].(string)
		if !ok {
			slog.Error("last field is not a string")
			continue
		}
		lastTime, err := time.Parse("2006-01-02 15:04:05", lastStr)
		if err != nil {
			slog.Error("failed to parse last field as time.Time", slog.String("value", lastStr))
			continue
		}
		rule := model.CallQuestionnaireRule{
			Last: lastTime,
			//	Id:            ruleData["id"].(int64),
			//	DomainId:      ruleData["domain_id"].(int64),
			//	Active:        ruleData["active"].(int64),
			//	CallDirection: ruleData["call_direction"].(string),
		}
		processedRules = append(processedRules, rule)
	}
	return processedRules, nil
}

func get_tasks(app *app.App) ([]any, error) {
	tasks, err := app.Store.ServiceStore().Array(context.Background(), `update call_audit.jobs jj
		set state = 1 -- active
		from (select *
		      from call_audit.jobs
		      where state = 0
		      order by id
		      for update
		      limit 10
		     ) j
		where j.id = jj.id
		returning *;
		`)
	if err != nil {
		slog.Error("Failed to truncate jobs table")
		return nil, err
	}
	// process tasks
	return tasks, nil
}

func StartJobs(app *app.App) {
	_, err := app.Store.ServiceStore().Execute(context.Background(), "TRUNCATE TABLE call_audit.jobs;")
	if err != nil {
		slog.Error("Failed to truncate jobs table")
		return
	}

	go func() {
		profs, err := getRules(app, 100)
		if err != nil {
			slog.Error("Failed to get profiles")
			return
		}
		for _, v := range profs {
			create_jobs(app, v.Last.Day(), int(v.Limit))
			slog.Info("Created jobs for rule", slog.Int64("rule_id", int64(v.Id)), slog.Int64("active", int64(v.Active)), slog.String("last_stored_at", v.Last.String()))
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			drop_finished_jobs(app)
			time.Sleep(time.Second)
		}
	}()

	p := pool.NewPool(100, 10)

	/*
		for _, v := range get_tasks() {
			p.Exec(v)
		}
	*/

	go func() {
		for i := 0; i <= 10000; i++ {
			p.Exec(&anyT{name: i})
			fmt.Println("start job ", i)
		}
	}()

}

func setDebug() {
	//debug.SetGCPercent(-1)

	go func() {
		slog.Info("Start debug server on http://localhost:8090/debug/pprof/")
		err := http.ListenAndServe(":8090", nil)
		if err != nil {
			slog.Error(err.Error())
		}
	}()

}
