package service

func (s serviceImpl) Health() error {	
	return s.repo.HealthCheck()
}