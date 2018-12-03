package models

import (
	"sync"
)

// read sessions
func ReadAllSessions() (*sync.Map, error) {
	var sessions []Session

	if err = db.Find(&sessions).Error; err != nil {
		return nil, err
	}

	var m *sync.Map
	for _, session := range sessions {
		m.Store(session.UserID, session.SessionID)
	}

	return m, nil
}

func CreateSession(sid string, ttl int64, uid uint) (*Session, error) {
	var session Session
	session = Session{
		SessionID: sid,
		TTL:       ttl,
		UserID:    uid,
	}

	if err := db.Create(&session).Error; err != nil {
		return nil, err
	}

	return &session, nil
}

func DeleteSession(uid uint) error {
	var se Session
	if err := db.Where("user_id = ?", uid).Unscoped().Delete(&se).Error; err != nil {
		return err
	}
	return nil
}
