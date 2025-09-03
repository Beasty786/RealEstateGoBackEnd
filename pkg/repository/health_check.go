package repository

func (r repository) HealthCheck() error {
	return r.db.Ping()
}
