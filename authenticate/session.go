package authenticate

import (
	"QuizPortal/database"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

//SessionStore ......
var SessionStore *sessions.FilesystemStore

//CreateUserSession ......
func CreateUserSession(u database.User, sessionID string, w http.ResponseWriter, r *http.Request) error {

	gfSession, err := SessionStore.Get(r, "QuizPortal")

	if err != nil {
		log.Print(err)
	}

	gfSession.Values["sessionID"] = sessionID
	gfSession.Values["username"] = u.Username
	gfSession.Values["email"] = u.Email
	gfSession.Values["branch"] = u.Branch
	gfSession.Values["year"] = u.Year
	gfSession.Values["uuid"] = u.UUID
	gfSession.Values["college"] = u.College
	gfSession.Values["contact"] = u.Contact


	err = gfSession.Save(r, w)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

//ExpireUserSession .......
func ExpireUserSession(w http.ResponseWriter, r *http.Request) {
	gfSession, err := SessionStore.Get(r, "QuizPortal")

	if err != nil {
		log.Print(err)
	}

	gfSession.Options.MaxAge = -1
	gfSession.Save(r, w)

}

func init() {
	if _, err := os.Stat("/tmp/QuizPortal"); os.IsNotExist(err) {
		os.Mkdir("/tmp/QuizPortal", 711)
	}
	SessionStore = sessions.NewFilesystemStore("", []byte(os.Getenv("hashkey")))
}
