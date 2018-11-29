package models

import "sync"

func CreateSession(sid string, ttl int64, uname string) error {
	var se Session
	se = Session{
		SessionID: sid,
		TTL:       ttl,
		Username:  uname,
	}

	if err := db.Create(&se).Error; err != nil {
		return err
	}

	return nil
}

func ReadSession(sid string) (*Session, error) {
	se := Session{}
	if err := db.Where("session_id = ?", sid).First(&se).Error; err != nil {
		return &Session{}, err
	}
	return &se, nil
}

func ReadAllSessions() (*sync.Map, error) {
	var ses []*Session
	if err := db.Find(&ses).Error; err != nil {
		return nil, nil
	}

	var m *sync.Map
	for _, item := range ses {
		m.Store(item.SessionID, item)
	}
	return m, nil
}

func DeleteSession(sid string) error {
	var se Session
	if err := db.Where("session_id", sid).Delete(&se).Error; err != nil {
		return err
	}
	return nil
}
