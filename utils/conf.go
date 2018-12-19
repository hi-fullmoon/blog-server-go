package utils

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	ConfFile *ini.File

	HTTPPort string
	RespTimeout time.Duration
	SessionAge int64
	PageSize int
)

func init() {
	var err error
	ConfFile, err = ini.Load("config/config.ini")
	if err != nil {
		log.Fatal(2, "Fail to parse 'config/config.ini': %v", err)
	}

	LoadBase()
	LoadServer()
}

func LoadBase() {
	sec, err := ConfFile.GetSection("base")
	if err != nil {
		log.Fatal(2, "Fail to get section `base`': %v", err)
	}
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
	SessionAge = sec.Key("SESSION_AGE").MustInt64(100)
}

func LoadServer() {
	sec, err := ConfFile.GetSection("server")
	if err != nil {
		log.Fatal(2, "Fail to get section `server`': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustString("8080")
	RespTimeout = time.Duration(sec.Key("RESP_TIMEOUT").MustInt(60)) * time.Second
}