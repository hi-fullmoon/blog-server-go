package models

// read sessions
func ReadAllSessions() ([]*Session, error) {
	var sessions []*Session

	if err = db.Find(&sessions).Error; err != nil {
		return nil, err
	}
	return sessions, nil
}

func CreateSession(sid string, ttl int64, uid uint) (*Session, error) {
	var session Session
	session = Session{
		SessionID: sid,
		TTL:       ttl,
		UserID:    uid,
	}

	if err = db.Create(&session).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func DeleteSession(uid uint) error {
	var se Session
	if err = db.Where("user_id = ?", uid).Unscoped().Delete(&se).Error; err != nil {
		return err
	}
	return nil
}

func DeleteSessionsByTTL(ttl int64) error {
	var ses []Session
	if err = db.Where("ttl < ?", ttl).Unscoped().Delete(&ses).Error; err != nil {
		return err
	}
	return nil
}
