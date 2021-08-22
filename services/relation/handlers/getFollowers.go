package handlers

func (h *handlers) GetFollowers(userId uint64) []uint64 {
	return h.srv.GetFollowers(userId)
}
