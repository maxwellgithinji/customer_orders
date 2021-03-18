package auth

import (
	"encoding/gob"
	"os"
	"time"

	"github.com/gorilla/sessions"
)

var (
	Store   *sessions.FilesystemStore
	Options *sessions.Options
)

func InitSession() error {
	sessionKey := os.Getenv("SESSION_KEY")
	Store = sessions.NewFilesystemStore("", []byte(sessionKey))
	Options = &sessions.Options{
		MaxAge:   int(time.Hour.Seconds()),
		HttpOnly: true,
	}
	gob.Register(map[string]interface{}{})
	return nil
}
