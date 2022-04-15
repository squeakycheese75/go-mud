package model

import (
	"net"
)

type User struct {
	ID      uint
	Name    string
	Session *Session
}

type Session struct {
	id   string
	Conn net.Conn
}

func (s *Session) WriteLine(str string) error {
	_, err := s.Conn.Write([]byte(str + "\r\n"))
	return err
}

func (s *Session) SessionId() string {
	return s.id
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
}

type Option struct {
	Choice string `json:"choice"`
	Next   int    `json:"next"`
	Key    string `json:"key"`
}
