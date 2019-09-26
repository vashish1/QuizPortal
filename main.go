package main

import (
	// "QuizPortal/authenticate"
	"QuizPortal/database"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var cl1, cl2, cl3 *mongo.Collection
var c *mongo.Client
var e database.Event
var eve Events
var count int

//Events ....
type Events struct {
	Index int
	Elist []database.Event
}

func main() {
	r := NewRouter()

	// r.HandleFunc("/QuizPortal/", About)
	// r.HandleFunc("/QuizPortal/signup", signuphandler).Methods("GET", "POST")
	// r.HandleFunc("/QuizPortal/login", loginhandler).Methods("GET", "POST")
	// r.HandleFunc("/QuizPortal/orglogin", loginhandler).Methods("GET", "POST")
	// r.HandleFunc("/QuizPortal/login/dashboard", dashboard).Methods("GET")
	// r.HandleFunc("/QuizPortal/contact", Contacthandler)
	r.HandleFunc("/QuizPortal/events", eventhandler).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}

//NewRouter .....
func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	return r
}

// func signuphandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("yahaan aagya")

// 	switch r.Method {

// 	case "GET":
// 		{

// 			fmt.Println("yeh chlra hai")
// 			t, err := template.ParseFiles("C:/Users/yashi/go/src/QuizPortal/templates/signup.html")
// 			if err != nil {
// 				log.Fatal("Could not parse template files\n")
// 			}
// 			er := t.Execute(w, "")
// 			if er != nil {
// 				log.Fatal("could not execute the files\n")
// 			}
// 		}
// 		log.Print("working")
// 	case "POST":
// 		{
// 			fmt.Println(" lets see if it works ")
// 			a := r.FormValue("username")

// 			b := r.FormValue("email")
// 			c := r.FormValue("branch")
// 			d := r.FormValue("year")
// 			e := r.FormValue("college")
// 			f := r.FormValue("contact")
// 			g := r.FormValue("password")

// 			u := database.Newuser(a, b, c, d, e, f, g)
// 			user := *u
// 			processsignupform(cl1, user, w, r)

// 		}
// 	}
// }

// func loginhandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("yahaan aagya")
// 	switch r.Method {

// 	case "GET":
// 		{

// 			fmt.Println("yeh chlra hai")
// 			t, err := template.ParseFiles("C:/Users/yashi/go/src/QuizPortal/templates/login.html")
// 			if err != nil {
// 				log.Fatal("Could not parse template files\n")
// 			}
// 			er := t.Execute(w, "")
// 			if er != nil {
// 				log.Fatal("could not execute the files\n")
// 			}
// 		}
// 		log.Print("working")
// 	case "POST":
// 		{
// 			fmt.Println(" lets see if it works ")
// 			a := r.FormValue("username")
// 			user := database.Finddb(cl1, a)
// 			processloginform(cl1, user, w, r)

// 		}

// 	}
// }

// func organizerhandler(w http.ResponseWriter, r *http.Request) {

// 	switch r.Method {

// 	case "GET":
// 		{

// 			fmt.Println("yeh chlra hai")
// 			t, err := template.ParseFiles("C:/Users/yashi/go/src/QuizPortal/templates/orglogin.html")
// 			if err != nil {
// 				log.Fatal("Could not parse template files\n")
// 			}
// 			er := t.Execute(w, "")
// 			if er != nil {
// 				log.Fatal("could not execute the files\n")
// 			}
// 		}
// 		log.Print("working")
// 	case "POST":
// 		{
// 			fmt.Println(" lets see if it works ")
// 			a := r.FormValue("username")
// 			user := database.Finddb(cl1, a)
// 			processloginform(cl1, user, w, r)

// 		}
// 	}
// }

// //dashboard ....
// func dashboard(w http.ResponseWriter, r *http.Request) {
// 	value := authenticate.Getcookievalues(w, r)
// 	fmt.Println(value)
// 	fmt.Println("yeh chlra hai")
// 	t, err := template.ParseFiles("C:/Users/yashi/go/src/QuizPortal/templates/dashboard.html")
// 	if err != nil {
// 		log.Fatal("Could not parse template files\n")
// 	}
// 	er := t.Execute(w, "")
// 	if er != nil {

// 		log.Fatal("could not execute the files\n")
// 	}
// }

// //About shows the information about the portal
// func About(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("about chlra hai")
// 	t, err := template.ParseFiles("C:/Users/yashi/go/src/QuizPortal/templates/dashboard.html")
// 	if err != nil {
// 		log.Fatal("Could not parse template files\n")
// 	}
// 	er := t.Execute(w, "")
// 	if er != nil {
// 		log.Fatal("could not execute the files\n")
// 	}
// }

// //Contacthandler .....
// func Contacthandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("contact chlra hai")
// 	t, err := template.ParseFiles("C:/Users/yashi/go/src/QuizPortal/templates/dashboard.html")
// 	if err != nil {
// 		log.Fatal("Could not parse template files\n")
// 	}
// 	er := t.Execute(w, "")
// 	if er != nil {
// 		log.Fatal("could not execute the files\n")
// 	}
// }

// func processsignupform(cl *mongo.Collection, U database.User, w http.ResponseWriter, r *http.Request) {
// 	t := database.Findfromuserdb(cl, U.Username)
// 	if t == true {
// 		processloginform(cl, U, w, r)
// 	}
// }

// func processloginform(cl *mongo.Collection, U database.User, w http.ResponseWriter, r *http.Request) {
// 	sessionID := database.GenerateUUID()
// 	rr := authenticate.CreateSecureCookie(U, sessionID, w, r)
// 	if rr != nil {
// 		log.Fatal("error occured while making a secure cookie:", rr)
// 	}
// 	er := authenticate.CreateUserSession(U, sessionID, w, r)
// 	if er != nil {
// 		log.Fatal("error occured while making a session")
// 	}
// 	http.Redirect(w, r, "/QuizPortal/login/dashboard", 302)

// }

//EVENTHANDLER HANDLES THE TEMPLATE FOR EVENTS SELECTION
func eventhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("event chlra hai")

	t, err := template.ParseFiles("C:/Users/yashi/go/src/QuizPortal/templates/events.html")
	if err != nil {
		log.Fatal("Could not parse template files\n")
	}
	er := t.Execute(w, eve)
	if er != nil {
		log.Fatal("could not execute the files\n:", er)
	}
}
func init() {
	cl1, cl2, cl3, c = database.Createdb()
	count = 0
	count++
	e = database.Event{
		I:                count,
		Eventsname:       "firstevent",
		Eventdescription: "a basic idea to quiz",
		Eventstartdate:   "Mon Sep 22",
		Timenow:          "05:00:00",
		Eventstarttime:   "22:00:00",
		Eventenddate:     "Tue Sep 23",
		Eventendtime:     "00:00:00",
		Datenow:          "",
	}
	database.Insertintoeventdb(cl3, e)
	eve = Events{
		Index: 1,
		Elist: database.Findfromeventdb(cl3),
	}

}
