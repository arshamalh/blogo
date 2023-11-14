package session

import "strconv"

// Sessions should be stored in a database,
// and it's better to use Redis for session as it's fast,
// but for now, we want to keep it simple.

type Session struct {
	SessionID string `json:"session_id"`
	UserID    uint   `json:"user_id"`
	Valid     bool   `json:"valid"`
}

var sessions []Session

func Create(user_id uint) *Session {
	for _, session := range sessions {
		if session.UserID == user_id {
			return &session
		}
	}
	session := Session{
		SessionID: strconv.Itoa(len(sessions) + 1),
		UserID:    user_id,
		Valid:     true,
	}
	sessions = append(sessions, session)
	return &session
}

func Get(session_id string) *Session {
	for _, session := range sessions {
		if session.SessionID == session_id {
			return &session
		}
	}
	return nil
}

func Invalidate(session_id string) {
	for i, session := range sessions {
		if session.SessionID == session_id {
			sessions[i].Valid = false
			return
		}
	}
}
