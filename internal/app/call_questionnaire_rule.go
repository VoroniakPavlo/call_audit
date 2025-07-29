package app

import (
	"context"

	pb "github.com/VoroniakPavlo/call_audit/api/call_audit"
)

type CallQuestionnaireRuleService struct {
	app *App
	pb.UnimplementedCallQuestionnaireRuleServiceServer
}

func NewCallQuestionnaireRuleService(app *App) (*CallQuestionnaireRuleService, error) {
	service := &CallQuestionnaireRuleService{
		app: app,
	}
	return service, nil
}

func (s *CallQuestionnaireRuleService) List(ctx context.Context, req *pb.Empty) (*pb.CallQuestionnaireRuleList, error) {
	return &pb.CallQuestionnaireRuleList{
		Items: []*pb.CallQuestionnaireRule{},
	}, nil
}

func (s *CallQuestionnaireRuleService) Create(ctx context.Context, req *pb.UpsertCallQuestionnaireRuleRequest) (*pb.CallQuestionnaireRule, error) {
	return req.Rule, nil
}

func (s *CallQuestionnaireRuleService) Update(ctx context.Context, req *pb.UpsertCallQuestionnaireRuleRequest) (*pb.CallQuestionnaireRule, error) {
	// TODO: додати логіку оновлення
	return req.Rule, nil
}

func (s *CallQuestionnaireRuleService) Delete(ctx context.Context, req *pb.DeleteCallQuestionnaireRuleRequest) (*pb.CallQuestionnaireRule, error) {
	// TODO: додати логіку видалення
	return nil, nil
}

func (s *CallQuestionnaireRuleService) Get(ctx context.Context, req *pb.GetCallQuestionnaireRuleRequest) (*pb.CallQuestionnaireRule, error) {
	// TODO: додати логіку отримання
	return &pb.CallQuestionnaireRule{
		Id: req.Id,
	}, nil
}
