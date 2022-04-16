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

type Story struct {
	Stages []Stage `json:"stages"`
}

type Stage struct {
	Page      int           `json:"page"`
	Narrative string        `json:"narrative"`
	Options   []Option      `json:"options"`
	Events    []interface{} `json:"events"`
	Action    string        `json:"action"`
}

type Option struct {
	Choice string `json:"choice"`
	Next   int    `json:"next"`
	Key    string `json:"key"`
}
