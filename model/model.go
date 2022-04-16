package model

import (
	"net"

	"github.com/google/uuid"
)

type User struct {
	ID      uuid.UUID
	Name    string
	Session *Session
}

type Session struct {
	Id   string
	Conn net.Conn
	User *User
}

func (s *Session) WriteLine(str string) error {
	_, err := s.Conn.Write([]byte(str + "\r\n"))
	return err
}

func (s *Session) Write(str string) error {
	_, err := s.Conn.Write([]byte(str))
	return err
}

func (s *Session) SessionId() string {
	return s.Id
}

type Dungeon struct {
	Users []*User
}
