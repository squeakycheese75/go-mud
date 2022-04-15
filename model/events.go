package model

type MessageEvent struct {
	Msg string
}

type UserJoinedEvent struct {
}

type SessionEvent struct {
	Session *Session
	Event   interface{}
}

type SessionCreatedEvent struct {
}

type SessionInputEvent struct {
	Msg string
}

type SessionDisconnectedEvent struct {
}
