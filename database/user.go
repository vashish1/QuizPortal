package database

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/google/uuid"
)

var i int

//User ......
type User struct {
	UUID              string
	Username          string
	Email             string
	Branch            string
	Year              string
	College           string
	Contact           string
	PasswordHash      string
	Timestampcreated  int64
	Timestampmodified int64
	Eventids          []string
}

//Newuser .....
func Newuser(username string, email string, branch string, year string, college string, contact string, password string) *User {

	Password := SHA256ofstring(password)
	now := time.Now()
	Unixtimestamp := now.Unix()
	U := User{UUID: GenerateUUID(), Username: username, Email: email, Branch: branch, Year: year, College: college, Contact: contact, PasswordHash: Password, Timestampcreated: Unixtimestamp, Timestampmodified: Unixtimestamp, Eventids: []string{}}
	return &U
}

//SHA256ofstring is a function which takes a string a reurns its sha256 hashed form
func SHA256ofstring(p string) string {
	h := sha1.New()
	h.Write([]byte(p))
	hash := hex.EncodeToString(h.Sum(nil))
	return hash
}

//GenerateUUID generates a unique id for every user.
func GenerateUUID() string {

	sd := uuid.New()
	return (sd.String())

}

//Organizer data
type Organizer struct {
	UUID              string
	Username          string
	PasswordHash      string
	Timestampcreated  int64
	Timestampmodified int64
	Events            []string
}

//NewEvent ........
func NewEvent(a string, b string, c string, d string, e string, f string) Event {
	var eve Event
	eve.I = i
	eve.Eventsname = a
	eve.Eventdescription = b
	start := c + " at " + d
	end := e + " at " + f
	fmt.Println(start, end)
	t := time.Now()
	t1 := t.Format("2006-Jan-02 at 03:04pm")
	t2, _ := time.Parse("2006-Jan-02 at 03:04pm", t1)
	eve.Timenow = t2
	eve.Starttime, _ = time.Parse("2006-Jan-02 at 03:04pm", start)
	eve.Endtime, _ = time.Parse("2006-Jan-02 at 03:04pm", end)
	i++
	return eve
}

//After Compares the time .....
func (e Event) After() bool {
	return e.Endtime.After(e.Timenow)
}

//Before ....
func (e Event) Before() bool {
	return e.Starttime.Before(e.Timenow)
}