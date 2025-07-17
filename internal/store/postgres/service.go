package postgres

import "context"

type ServiceStore struct {
	storage *Store
}

func (s *ServiceStore) Execute(ctx context.Context, query string, args ...interface{}) (result interface{}, err error) {
	// Example implementation, replace with actual logic
	conn, err := s.storage.conn.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	result, err = conn.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *ServiceStore) Array(ctx context.Context, query string, args ...interface{}) ([]interface{}, error) {
	// Example implementation, replace with actual logic
	conn, err := s.storage.conn.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	rows, err := conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var results []interface{}
	for rows.Next() {
		var row interface{}
		if err := rows.Scan(&row); err != nil {
			return nil, err
		}
		results = append(results, row)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

func NewServiceStore(storage *Store) *ServiceStore {
	return &ServiceStore{storage: storage}
}
