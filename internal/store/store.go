package store

import (
	"context"

	_go "github.com/VoroniakPavlo/call_audit/api/protos/storage"
	dberr "github.com/VoroniakPavlo/call_audit/internal/errors"
)

type Store interface {
	LanguageProfiles() LanguageProfileStore
	CallQuestionnaireRules() CallQuestionnaireRuleStore
	ServiceStore() ServiceStore
	// ------------ Database Management ------------ //
	Open() *dberr.DBError  // Return custom DB error
	Close() *dberr.DBError // Return custom DB error

}

// LanguageProfileStore defines the methods for managing language profiles.
type LanguageProfileStore interface {
	Create(ctx context.Context, profile *_go.LanguageProfile) (*_go.LanguageProfile, error)
	Get(ctx context.Context, id int64) (*_go.LanguageProfile, error)
	Update(ctx context.Context, profile *_go.LanguageProfile) (*_go.LanguageProfile, error)
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]*_go.LanguageProfile, error)
}

// CallQuestionnaireRuleStore defines the methods for managing call questionnaire rules.
type CallQuestionnaireRuleStore interface {
	Create(ctx context.Context, rule *_go.CallQuestionnaireRule) (*_go.CallQuestionnaireRule, error)
	Get(ctx context.Context, id int64) (*_go.CallQuestionnaireRule, error)
	Update(ctx context.Context, rule *_go.CallQuestionnaireRule) (*_go.CallQuestionnaireRule, error)
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) (*_go.CallQuestionnaireRuleList, error)
}

type ServiceStore interface {
	Execute(ctx context.Context, query string, args ...interface{}) (result interface{}, err error)
	Array(ctx context.Context, query string, args ...interface{}) ([]interface{}, error)
}