package tcp

import (
	"codeburg.com/da-br/metql/internal/database"
)

type Server interface {
	ListenAndServe(url string) error
	Stop() error
}

type TcpServer struct{}

func NewServer(db database.DataBase) Server {
	return nil
}

func (s *TcpServer) ListenAndServe(url string) error {
	return nil
}
