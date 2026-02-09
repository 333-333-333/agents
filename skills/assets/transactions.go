// Use pgx transactions directly
func (r *PostgresUserRepository) WithTx(ctx context.Context, pool *pgxpool.Pool, fn func(q *db.Queries) error) error {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin tx: %w", err)
	}
	defer tx.Rollback(ctx)

	if err := fn(r.q.WithTx(tx)); err != nil {
		return err
	}

	return tx.Commit(ctx)
}
