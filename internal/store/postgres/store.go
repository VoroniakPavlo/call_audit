package postgres

import (
	"context"
	"log/slog"

	conf "github.com/VoroniakPavlo/call_audit/config"
	dberr "github.com/VoroniakPavlo/call_audit/internal/errors"

	"github.com/jackc/pgx/v5/pgxpool"
	otelpgx "github.com/webitel/webitel-go-kit/tracing/pgx"
)

// Store is the struct implementing the Store interface.
type Store struct {
	//------------cases stores ------------ ----//

	config *conf.DatabaseConfig
	conn   *pgxpool.Pool
}

// Database returns the database connection or a custom error if it is not opened.
func (s *Store) Database() (*pgxpool.Pool, *dberr.DBError) { // Return custom DB error
	if s.conn == nil {
		return nil, dberr.NewDBError("store.database.check.bad_arguments", "database connection is not opened")
	}
	return s.conn, nil
}

// Open establishes a connection to the database and returns a custom error if it fails.
func (s *Store) Open() *dberr.DBError {
	config, err := pgxpool.ParseConfig(s.config.Url)
	if err != nil {
		return dberr.NewDBError("store.open.parse_config.fail", err.Error())
	}

	// Attach the OpenTelemetry tracer for pgx
	config.ConnConfig.Tracer = otelpgx.NewTracer(otelpgx.WithTrimSQLInSpanName())

	conn, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return dberr.NewDBError("store.open.connect.fail", err.Error())
	}
	s.conn = conn
	slog.Debug("cases.store.connection_opened", slog.String("message", "postgres: connection opened"))
	return nil
}

// Close closes the database connection and returns a custom error if it fails.
func (s *Store) Close() *dberr.DBError {
	if s.conn != nil {
		s.conn.Close()
		slog.Debug("cases.store.connection_closed", slog.String("message", "postgres: connection closed"))
		s.conn = nil
	}
	return nil
}
