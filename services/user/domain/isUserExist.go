package domain

func (s *service) IsUserExist(userId uint64) bool {
	return s.repo.Read.IsUserExistById(userId)
}
