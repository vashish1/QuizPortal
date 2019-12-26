package authenticate

import (
	"QuizPortal/database"
	"fmt"
	"net/http"

	"time"

	"github.com/gorilla/securecookie"
)

var hashKey []byte
var blockKey []byte
var s *securecookie.SecureCookie

//CreateSecureCookie .....
func CreateSecureCookie(u database.User, sessionID string, w http.ResponseWriter, r *http.Request) error {

	value := map[string]string{
		"username": u.Username,
		"sid":      sessionID,
	}

	if encoded, err := s.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:     "session",
			Value:    encoded,
			Path:     "/",
			Secure:   false,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	} else {
		return err
	}

	return nil

}
//CreateSecureorgCookie .....
func CreateSecureorgCookie(u database.Organizer, sessionID string, w http.ResponseWriter, r *http.Request) error {

	value := map[string]string{
		"username": u.Username,
		"sid":      sessionID,
	}

	if encoded, err := s.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:     "session",
			Value:    encoded,
			Path:     "/",
			Secure:   false,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	} else {
		return err
	}

	return nil

}

//ReadSecureCookieValues ....
func ReadSecureCookieValues(w http.ResponseWriter, r *http.Request) (map[string]string, error) {
	if cookie, err := r.Cookie("session"); err == nil {
		value := make(map[string]string)
		if err = s.Decode("session", cookie.Value, &value); err == nil {
			return value, nil
		}
		return nil, err
	}
	return nil, nil

}

//ExpireSecureCookie ........
func ExpireSecureCookie(w http.ResponseWriter, r *http.Request) {

	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
	w.Header().Set("Expires", time.Unix(0, 0).Format(http.TimeFormat))
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("X-Accel-Expires", "0")

	http.SetCookie(w, cookie)
}

func init() {

	hashKey = securecookie.GenerateRandomKey(32)
	blockKey = securecookie.GenerateRandomKey(32)

	s = securecookie.New(hashKey, blockKey)
}

//Values ....
type Values map[string]string

//Getcookievalues ....
func Getcookievalues(w http.ResponseWriter, r *http.Request) Values {
	c, err := r.Cookie("session")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return Values{}
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return Values{}
	}
	var value Values

	er := s.Decode("session", c.Value, &value)
	// if er == nil {
	// 	fmt.println(w, value)
	// }
	fmt.Print(er)
	return value
}
