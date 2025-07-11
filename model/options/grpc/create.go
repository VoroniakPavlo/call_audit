package grpc

import (
	"context"
	"errors"
	"time"

	"github.com/VoroniakPavlo/call_audit/auth"
	"github.com/VoroniakPavlo/call_audit/model"
	"github.com/VoroniakPavlo/call_audit/model/options"
	"github.com/VoroniakPavlo/call_audit/model/options/grpc/shared"
	"github.com/VoroniakPavlo/call_audit/util"
)

type CreateOption func(*CreateOptions) error

func WithCreateFields(
	fielder shared.Fielder,
	md model.ObjectMetadatter,
	fieldsModifiers ...func(fields []string) []string,
) CreateOption {
	return func(o *CreateOptions) error {
		if requestedFields := fielder.GetFields(); len(requestedFields) == 0 {
			o.Fields = md.GetDefaultFields()
		} else {
			o.Fields = util.DeduplicateFields(util.FieldsFunc(
				requestedFields, util.InlineFields,
			))
		}
		o.Fields, o.UnknownFields = util.SplitKnownAndUnknownFields(o.Fields, md.GetAllFields())
		for _, f := range fieldsModifiers {
			o.Fields = f(o.Fields)
		}
		return nil
	}
}

func WithCreateIDs(ids []int64) CreateOption {
	return func(o *CreateOptions) error {
		o.IDs = ids
		return nil
	}
}

func WithCreateParentID(id int64) CreateOption {
	return func(o *CreateOptions) error {
		o.ParentID = id
		return nil
	}
}

func WithCreateChildID(id int64) CreateOption {
	return func(o *CreateOptions) error {
		o.ChildID = id
		return nil
	}
}

type CreateOptions struct {
	context.Context
	Time              time.Time
	Fields            []string
	DerivedSearchOpts map[string]*options.SearchOptions
	UnknownFields     []string
	IDs               []int64
	ParentID          int64
	ChildID           int64
	Auth              auth.Auther
}

func (s *CreateOptions) SetAuthOpts(a auth.Auther) *CreateOptions {
	s.Auth = a
	return s
}

func (s *CreateOptions) RequestTime() time.Time { return s.Time }
func (s *CreateOptions) GetAuthOpts() auth.Auther {
	return s.Auth
}
func (s *CreateOptions) GetIDs() []int64     { return s.IDs }
func (s *CreateOptions) GetParentID() int64  { return s.ParentID }
func (s *CreateOptions) GetFields() []string { return s.Fields }
func (s *CreateOptions) GetDerivedSearchOpts() map[string]*options.SearchOptions {
	return s.DerivedSearchOpts
}
func (s *CreateOptions) GetUnknownFields() []string { return s.UnknownFields }
func (s *CreateOptions) GetChildID() int64          { return s.ChildID }

type Creator interface {
	GetFields() []string
}

func (s *CreateOptions) CurrentTime() time.Time {
	ts := s.Time
	if ts.IsZero() {
		ts = time.Now().UTC()
		s.Time = ts
	}
	return ts
}

func NewCreateOptions(ctx context.Context, opts ...CreateOption) (*CreateOptions, error) {
	createOpts := &CreateOptions{
		Context:           ctx,
		Time:              time.Now().UTC(),
		DerivedSearchOpts: make(map[string]*options.SearchOptions),
	}

	// Set authentication
	if err := setCreateAuthOptions(ctx, createOpts); err != nil {
		return nil, err
	}

	for _, opt := range opts {
		err := opt(createOpts)
		if err != nil {
			return nil, err
		}
	}
	return createOpts, nil
}

// setUpdateAuthOptions extracts authentication from context and sets it in options
func setCreateAuthOptions(ctx context.Context, options *CreateOptions) error {
	if sess := model.GetAutherOutOfContext(ctx); sess != nil {
		options.Auth = sess
		return nil
	}
	return errors.New("can't authorize user")
}
