package session

import (
	"fmt"
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

	//LoadSessionsFromDB()
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
		sessionMap.Store(key, ss)
		return true
	})
}

func GenerateNewSessionId(uname string) string {
	sid, _ := uuid.NewV4()
	sidStr := sid.String()
	ct := nowInMilli()
	ttl := ct + 1*60*1000

	se, err := models.CreateSession(sidStr, ttl, uname)
	if err != nil {
		return ""
	}
	sessionMap.Store(sidStr, se)

	return sidStr
}

func deleteIsExpiredSession(sid string) error {
	if err := models.DeleteSession(sid); err != nil {
		return err
	}
	sessionMap.Delete(sid)
	return nil
}

func IsSessionExpired(sid string) (string, bool) {
	if s, ok := sessionMap.Load(sid); ok {
		ct := nowInMilli()
		if s.(*models.Session).TTL < ct {
			deleteIsExpiredSession(sid)
			return "", true
		}

		return s.(*models.Session).Username, false
	}
	s, _ := sessionMap.Load(sid)
	fmt.Println("sssssssssssssssssss", s)
	return "", true
}
