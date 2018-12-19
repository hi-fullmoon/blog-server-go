package cron

import (
	"zhengbiwen/blog-server/session"

	"github.com/robfig/cron"
)

func Run() {
	c := cron.New()
	// two o'clock a day
	spec := "0 0 2 * * ?"
	c.AddFunc(spec, func() {
		session.DeleteExpiredSessions()
	})
	c.Start()
}
