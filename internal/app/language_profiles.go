package app

import (
	"context"

	pb "github.com/VoroniakPavlo/call_audit/api/call_audit"
)

type LanguageProfilesService struct {
	app *App // Define App type below or import from the correct package
	pb.UnimplementedLanguageProfileServiceServer
}

func (s *LanguageProfilesService) CreateLanguageProfile(ctx context.Context, req *pb.CreateLanguageProfileRequest) (*pb.CreateLanguageProfileResponse, error) {
	// TODO: implement your logic here
	return &pb.CreateLanguageProfileResponse{}, nil
}

func (s *LanguageProfilesService) GetLanguageProfile(ctx context.Context, req *pb.GetLanguageProfileRequest) (*pb.GetLanguageProfileResponse, error) {
	// TODO: implement your logic here
	return &pb.GetLanguageProfileResponse{}, nil
}

func (s *LanguageProfilesService) UpdateLanguageProfile(ctx context.Context, req *pb.UpdateLanguageProfileRequest) (*pb.UpdateLanguageProfileResponse, error) {
	// TODO: implement your logic here
	return &pb.UpdateLanguageProfileResponse{}, nil
}

func (s *LanguageProfilesService) DeleteLanguageProfile(ctx context.Context, req *pb.DeleteLanguageProfileRequest) (*pb.DeleteLanguageProfileResponse, error) {
	// TODO: implement your logic here
	return &pb.DeleteLanguageProfileResponse{}, nil
}

func (s *LanguageProfilesService) ListLanguageProfiles(ctx context.Context, req *pb.ListLanguageProfilesRequest) (*pb.ListLanguageProfilesResponse, error) {
	// TODO: implement your logic here
	return &pb.ListLanguageProfilesResponse{}, nil
}

func NewLanguageProfileService(app *App) (*LanguageProfilesService, error) {

	service := &LanguageProfilesService{
		app: app,
	}

	return service, nil
}
