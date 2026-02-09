// cmd/server/main.go
func newUserRepository(cfg *config.Config, pool *pgxpool.Pool) domain.UserRepository {
	if cfg.Database.Provider == "memory" {
		return repository.NewInMemoryUserRepository()
	}
	return repository.NewPostgresUserRepository(pool)
}
