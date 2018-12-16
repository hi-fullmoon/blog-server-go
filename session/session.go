package session

import (
	"sync"
	"time"
	"zhengbiwen/blog-server/models"

	"github.com/satori/go.uuid"
)

type SimpleSession struct {
	Username string
	TTL      int64
}

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func nowInMilli() int64 {
	return time.Now().UnixNano() / 1000000
}

func LoadSessionsFromDB() {
	sessions, err := models.ReadAllSessions()
	if err != nil {
		return
	}

	for _, session := range sessions {
		sessionMap.Store(session.UserID, session)
	}
}

func GenerateNewSessionId(uid uint) string {
	var sessionId string

	if session, ok := sessionMap.Load(uid); ok {
		sessionId = session.(*models.Session).SessionID
	} else {
		sid, _ := uuid.NewV4()
		sessionId = sid.String()

		ct := nowInMilli()
		ttl := ct + 60*60*1000

		s, err := models.CreateSession(sessionId, ttl, uid)
		if err != nil {
			return ""
		}
		sessionMap.Store(uid, s)
	}

	return sessionId
}

func IsSessionExpired(uid uint) (string, bool) {
	if s, ok := sessionMap.Load(uid); ok {
		ct := nowInMilli()
		if s.(*models.Session).TTL < ct {
			deleteIsExpiredSession(uid)
			return "", true
		}

		return s.(*models.Session).SessionID, false
	}
	return "", true
}

func deleteIsExpiredSession(uid uint) error {
	if err := models.DeleteSession(uid); err != nil {
		return err
	}
	sessionMap.Delete(uid)
	return nil
}
