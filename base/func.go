package base

import (
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
)

const (
	TIME_OUT = 30
)

func Timeout() {
	time.AfterFunc(TIME_OUT*time.Second, func() {
		log.Error("Execute timeout")
		os.Exit(1)
	})
}

func CheckIsDir(sdir string) {
	finfo, err := os.Stat(sdir)
	if err != nil {
		log.WithField("dir path", sdir).Fatalf("mysql socket file dir does not exists: %v", err)
	}
	if ret := finfo.IsDir(); !ret {
		log.WithField("dir path", sdir).Fatalf("the input path [%v] is not a dir,please check it.", sdir)
	}
}

func InitLog(cfg *Cfg) {
	// Init log file
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.Level(cfg.LogLevel))

	if cfg.LogFile != "" {
		f, err := os.OpenFile(cfg.LogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err == nil {
			log.SetOutput(f)
			return
		}
	}
	log.SetOutput(os.Stderr)
}
