package session

import (
	"strconv"

	"go.uber.org/zap"
)

var Gl *zap.Logger

func init() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic("Failed to initialize logger")
	}
	Gl = logger
}

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
	Gl.Info("Session created",
		zap.String("session_id", session.SessionID),
		zap.Uint("user_id", session.UserID),
		zap.Bool("valid", session.Valid))
	return &session
}

func Get(session_id string) *Session {
	for _, session := range sessions {
		if session.SessionID == session_id {
			return &session
		}
	}
	Gl.Error("Session not found", zap.String("session_id", session_id))
	return nil
}

func Invalidate(session_id string) {
	for i, session := range sessions {
		if session.SessionID == session_id {
			sessions[i].Valid = false
			Gl.Info("Session invalidated", zap.String("session_id", session_id))
			return
		}
	}
	Gl.Error("Unable to invalidate session", zap.String("session_id", session_id))
}
