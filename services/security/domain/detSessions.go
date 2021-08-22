package domain

import "freeSociety/services/security/models"

func (s *service) GetSessions(userId uint64) ([]*models.Session, error) {
	return s.postgresRepo.Read.GetSessions(userId)
}
