package app

import (
	"context"

	pb "github.com/VoroniakPavlo/call_audit/api/protos/storage"
)

type CallQuestionnaireRuleService struct {
	app *App // Define App type below or import from the correct package
	pb.UnimplementedCallQuestionnaireRuleServiceServer
}

func (s *CallQuestionnaireRuleService) Create(ctx context.Context, req *pb.UpsertCallQuestionnaireRuleRequest) (*pb.CallQuestionnaireRule, error) {
	// TODO: implement
	return &pb.CallQuestionnaireRule{}, nil
}

func (s *CallQuestionnaireRuleService) Get(ctx context.Context, req *pb.GetCallQuestionnaireRuleRequest) (*pb.CallQuestionnaireRule, error) {
	// TODO: implement
	return &pb.CallQuestionnaireRule{}, nil
}

func (s *CallQuestionnaireRuleService) Update(ctx context.Context, req *pb.UpsertCallQuestionnaireRuleRequest) (*pb.CallQuestionnaireRule, error) {
	// TODO: implement
	return &pb.CallQuestionnaireRule{}, nil
}

func (s *CallQuestionnaireRuleService) Delete(ctx context.Context, req *pb.DeleteCallQuestionnaireRuleRequest) (*pb.CallQuestionnaireRule, error) {
	// TODO: implement
	return &pb.CallQuestionnaireRule{}, nil
}

func (s *CallQuestionnaireRuleService) List(ctx context.Context) (*pb.CallQuestionnaireRuleList, error) {
	// TODO: implement
	return &pb.CallQuestionnaireRuleList{}, nil
}
