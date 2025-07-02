package options

import (
	"context"
	"time"

	"github.com/VoroniakPavlo/cases/auth"
	"github.com/webitel/webitel-go-kit/etag"
)

type UpdateOptions interface {
	context.Context
	GetAuthOpts() auth.Auther
	GetFields() []string
	GetUnknownFields() []string
	GetDerivedSearchOpts() map[string]*SearchOptions
	RequestTime() time.Time
	GetMask() []string
	GetEtags() []*etag.Tid
	GetParentID() int64
	GetIDs() []int64
}
