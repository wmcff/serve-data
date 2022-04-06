package session

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/wmcff/serve-data/model"
)

const (
	// sessionStr represents a string of session key.
	sessionStr = "GSESSION"
	// User is the key of user data in the session.
	User = "User"
)

// Get returns a session for the current request.
func Get(c echo.Context) *sessions.Session {
	sess, _ := session.Get(sessionStr, c)
	return sess
}

// Save saves the current session.
func Save(c echo.Context) error {
	sess := Get(c)
	sess.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
	}
	return saveSession(c, sess)
}

// Delete the current session.
func Delete(c echo.Context) error {
	sess := Get(c)
	sess.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	}
	return saveSession(c, sess)
}

func saveSession(c echo.Context, sess *sessions.Session) error {
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return nil
}

// SetValue sets a key and a value.
func SetValue(c echo.Context, key string, value interface{}) error {
	sess := Get(c)
	bytes, err := json.Marshal(value)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	sess.Values[key] = string(bytes)
	return nil
}

// GetValue returns value of session.
func GetValue(c echo.Context, key string) string {
	sess := Get(c)
	if sess != nil {
		if v, ok := sess.Values[key]; ok {
			data, result := v.(string)
			if result && data != "null" {
				return data
			}
		}
	}
	return ""
}

// SetUser sets user data in session.
func SetUser(c echo.Context, user *model.User) error {
	return SetValue(c, User, user)
}

// GetUser returns user object of session.
func GetUser(c echo.Context) *model.User {
	if v := GetValue(c, User); v != "" {
		a := &model.User{}
		_ = json.Unmarshal([]byte(v), a)
		return a
	}
	return nil
}
