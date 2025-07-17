package postgres

import (
	"context"

	lp "github.com/VoroniakPavlo/call_audit/api/protos/storage"
)

// CallQuestionnaireRule provides methods to interact with call questionnaire rules in the database.
type CallQuestionnaireRuleStore struct {
	storage *Store
}

// Create implements store.CallQuestionnaireRuleStore.
func (c *CallQuestionnaireRuleStore) Create(ctx context.Context, rule *lp.CallQuestionnaireRule) (*lp.CallQuestionnaireRule, error) {
	panic("unimplemented")
}

// Delete implements store.CallQuestionnaireRuleStore.
func (c *CallQuestionnaireRuleStore) Delete(ctx context.Context, id int64) error {
	panic("unimplemented")
}

// Get implements store.CallQuestionnaireRuleStore.
func (c *CallQuestionnaireRuleStore) Get(ctx context.Context, id int64) (*lp.CallQuestionnaireRule, error) {
	panic("unimplemented")
}

// List implements store.CallQuestionnaireRuleStore.
func (c *CallQuestionnaireRuleStore) List(ctx context.Context) ([]*lp.CallQuestionnaireRule, error) {
	panic("unimplemented")
}

// Update implements store.CallQuestionnaireRuleStore.
func (c *CallQuestionnaireRuleStore) Update(ctx context.Context, rule *lp.CallQuestionnaireRule) (*lp.CallQuestionnaireRule, error) {
	panic("unimplemented")
}

// NewCallQuestionnaireRuleStore creates a new CallQuestionnaireRuleStore.
func NewCallQuestionnaireRuleStore(storage *Store) *CallQuestionnaireRuleStore {
	return &CallQuestionnaireRuleStore{storage: storage}
}
