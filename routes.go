package main

import (
	"QuizPortal/database"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var cl1, cl2, cl3, cl4 *mongo.Collection
var c *mongo.Client
var eve Events
var e empty

type empty struct {
	res http.ResponseWriter
	req *http.Request
}

//Events ....
type Events struct {
	res   http.ResponseWriter
	req   *http.Request
	Elist []database.Event
}

//Quiz .......
type Quiz struct {
	res       http.ResponseWriter
	req       *http.Request
	Index     int
	Eventname string
	Q         database.Quizz
}

//Myevents stores the events of a particular organizer
type Myevents struct {
	res      http.ResponseWriter
	req      *http.Request
	username string
	List     []database.Event
}

func main() {
	r := NewRouter()

	r.HandleFunc("/QuizPortal", About)
	r.HandleFunc("/QuizPortal/signup", signuphandler).Methods("GET", "POST")
	r.HandleFunc("/QuizPortal/login", loginhandler).Methods("GET", "POST")
	r.HandleFunc("/QuizPortal/organizer", organizerhandler).Methods("GET", "POST")
	r.HandleFunc("/QuizPortal/organizer/dashboard", orgdashboard).Methods("GET")
	r.HandleFunc("/QuizPortal/login/dashboard", dashboard).Methods("GET")
	r.HandleFunc("/QuizPortal/contact", Contacthandler)
	r.HandleFunc("/QuizPortal/events", eventhandler).Methods("GET", "POST")
	r.HandleFunc("/quiz/", quizhandler).Methods("GET", "POST")
	r.HandleFunc("/QuizPortal/update/", updateevent).Methods("GET", "POST")
	r.HandleFunc("/QuizPortal/deleteEvent/", deleteevent).Methods("GET", "POST")
	r.HandleFunc("/QuizPortal/addEvent", addevent).Methods("GET", "POST")
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}

//NewRouter .....
func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	return r
}
