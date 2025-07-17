package postgres

import (
	"context"
	"github.com/VoroniakPavlo/call_audit/api/protos/storage"
)

// LanguageProfileStore provides methods to interact with language profiles in the database.
type LanguageProfileStore struct {
	storage *Store
}

// Create implements store.LanguageProfileStore.
func (l *LanguageProfileStore) Create(ctx context.Context, profile *storage.LanguageProfile) (*storage.LanguageProfile, error) {
	panic("unimplemented")
}

// Delete implements store.LanguageProfileStore.
func (l *LanguageProfileStore) Delete(ctx context.Context, id int64) error {
	panic("unimplemented")
}

// Get implements store.LanguageProfileStore.
func (l *LanguageProfileStore) Get(ctx context.Context, id int64) (*storage.LanguageProfile, error) {
	panic("unimplemented")
}

// List implements store.LanguageProfileStore.
func (l *LanguageProfileStore) List(ctx context.Context) ([]*storage.LanguageProfile, error) {
	panic("unimplemented")
}

// Update implements store.LanguageProfileStore.
func (l *LanguageProfileStore) Update(ctx context.Context, profile *storage.LanguageProfile) (*storage.LanguageProfile, error) {
	panic("unimplemented")
}

// NewLanguageProfileStore creates a new LanguageProfileStore.
func NewLanguageProfileStore(storage *Store) *LanguageProfileStore {
	return &LanguageProfileStore{storage: storage}
}
