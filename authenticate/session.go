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

	Session, err := SessionStore.Get(r, "QuizPortal")

	if err != nil {
		log.Print(err)
	}

	Session.Values["sessionID"] = sessionID
	Session.Values["username"] = u.Username
	Session.Values["email"] = u.Email
	Session.Values["branch"] = u.Branch
	Session.Values["year"] = u.Year
	Session.Values["uuid"] = u.UUID
	Session.Values["college"] = u.College
	Session.Values["contact"] = u.Contact


	err = Session.Save(r, w)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

//ExpireUserSession .......
func ExpireUserSession(w http.ResponseWriter, r *http.Request) {
	Session, err := SessionStore.Get(r, "QuizPortal")

	if err != nil {
		log.Print(err)
	}

	Session.Options.MaxAge = -1
	 Session.Save(r, w)

}

func init() {
	if _, err := os.Stat("/tmp"); os.IsNotExist(err) {
		os.Mkdir("/tmp", 711)
	}
	SessionStore = sessions.NewFilesystemStore("/tmp", []byte(os.Getenv("hashkey")))
}