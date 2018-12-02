package models

import (
	"fmt"
	"sync"
)

// read sessions
func ReadAllSessions() (*sync.Map, error) {
	var sessions []Session

	if err = db.Find(&sessions).Error; err != nil {
		return nil, err
	}

	fmt.Println(sessions)

	var m *sync.Map
	for _, session := range sessions {
		m.Store(session.SessionID, 1)
	}

	return m, nil
}

func CreateSession(sid string, ttl int64, username string) (*Session, error) {
	var se Session
	se = Session{
		SessionID: sid,
		TTL:       ttl,
		Username:  username,
	}

	if err := db.Create(&se).Error; err != nil {
		return &se, err
	}

	return &se, nil
}

func ReadSession(sid string) (*Session, error) {
	se := Session{}
	if err := db.Where("session_id = ?", sid).First(&se).Error; err != nil {
		return &Session{}, err
	}
	return &se, nil
}

func DeleteSession(sid string) error {
	var se Session
	if err := db.Where("session_id", sid).Delete(&se).Error; err != nil {
		return err
	}
	return nil
}
