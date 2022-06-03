package server

// Routes contains the server routes
func (s *server) Routes() {
	s.router.Handle("/api/v1/", s.handleIndex())
	s.router.Handle("/api/v1/add", s.handleAddition())
}
