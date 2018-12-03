package session

import (
	"log"
	"sync"
	"time"
	"zhengbiwen/blog_management_system/models"

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
	res, err := models.ReadAllSessions()
	if err != nil {
		return
	}

	res.Range(func(key, value interface{}) bool {
		ss := value.(*models.Session)
		sessionMap.Store(key, ss.SessionID)
		return true
	})
}

func GenerateNewSessionId(uid uint) string {
	var sessionId string

	if session, ok := sessionMap.Load(uid); ok {
		sessionId = session.(*models.Session).SessionID
		log.Println("session id is exist: ", sessionId)
	} else {
		sid, _ := uuid.NewV4()
		sessionId = sid.String()

		ct := nowInMilli()
		ttl := ct + 1*60*1000

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
