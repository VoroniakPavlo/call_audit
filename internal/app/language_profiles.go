package app

import (
	"context"

	"github.com/VoroniakPavlo/call_audit/api/protos/storage"
)

type LanguageProfilesService struct {
	app *App
	storage.UnimplementedLanguageProfileServiceServer
}

func (s *LanguageProfilesService) CreateLanguageProfile(ctx context.Context, req *storage.CreateLanguageProfileRequest) (*storage.CreateLanguageProfileResponse, error) {
	// TODO: implement your logic here
	return &storage.CreateLanguageProfileResponse{}, nil
}

func (s *LanguageProfilesService) GetLanguageProfile(ctx context.Context, req *storage.GetLanguageProfileRequest) (*storage.GetLanguageProfileResponse, error) {
	// TODO: implement your logic here
	return &storage.GetLanguageProfileResponse{}, nil
}

func (s *LanguageProfilesService) UpdateLanguageProfile(ctx context.Context, req *storage.UpdateLanguageProfileRequest) (*storage.UpdateLanguageProfileResponse, error) {
	// TODO: implement your logic here
	return &storage.UpdateLanguageProfileResponse{}, nil
}

func (s *LanguageProfilesService) DeleteLanguageProfile(ctx context.Context, req *storage.DeleteLanguageProfileRequest) (*storage.DeleteLanguageProfileResponse, error) {
	// TODO: implement your logic here
	return &storage.DeleteLanguageProfileResponse{}, nil
}

func (s *LanguageProfilesService) ListLanguageProfiles(ctx context.Context, req *storage.ListLanguageProfilesRequest) (*storage.ListLanguageProfilesResponse, error) {
	// TODO: implement your logic here
	return &storage.ListLanguageProfilesResponse{}, nil
}