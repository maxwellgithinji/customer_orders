package auth

import (
	"encoding/gob"
	"os"

	"github.com/gorilla/sessions"
)

var (
	Store *sessions.FilesystemStore
)

func InitSession() error {
	sessionKey := os.Getenv("SESSION_KEY")
	Store = sessions.NewFilesystemStore("", []byte(sessionKey))
	gob.Register(map[string]interface{}{})
	return nil
}
