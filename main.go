package main

import (
	"QuizPortal/authenticate"
	"QuizPortal/database"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var cl1, cl2, cl3,cl4  *mongo.Collection
var c *mongo.Client
var e database.Event
var eve Events
var org database.Organizer

type functions interface {
}

//Events ....
type Events struct {
	res   http.ResponseWriter
	req   *http.Request
	Index int
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

func main() {
	r := NewRouter()

	r.HandleFunc("/QuizPortal", About)
	r.HandleFunc("/QuizPortal/signup", signuphandler).Methods("GET", "POST")
	r.HandleFunc("/QuizPortal/login", loginhandler).Methods("GET", "POST")
	r.HandleFunc("/QuizPortal/organizer", organizerhandler).Methods("GET", "POST")
	r.HandleFunc("/QuizPortal/login/dashboard", dashboard).Methods("GET")
	r.HandleFunc("/QuizPortal/contact", Contacthandler)
	r.HandleFunc("/QuizPortal/events", eventhandler).Methods("GET", "POST")
	r.HandleFunc("/quiz/", quizhandler).Methods("GET", "POST")
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}

//NewRouter .....
func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	return r
}

func signuphandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("yahaan aagya")

	switch r.Method {

	case "GET":
		{

			fmt.Println("yeh chlra hai")
			t, err := template.ParseFiles("C:/Users/yashi/go/src/QuizPortal/templates/signup.html")
			if err != nil {
				log.Fatal("Could not parse template files\n")
			}
			er := t.Execute(w, "")
			if er != nil {
				log.Fatal("could not execute the files\n")
			}
		}
		log.Print("working")
	case "POST":
		{
			fmt.Println(" lets see if it works ")
			a := r.FormValue("username")

			b := r.FormValue("email")
			c := r.FormValue("branch")
			d := r.FormValue("year")
			e := r.FormValue("college")
			f := r.FormValue("contact")
			g := r.FormValue("password")

			u := database.Newuser(a, b, c, d, e, f, g)
			user := *u
			processsignupform(cl1, user, w, r)

		}
	}
}

func loginhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("yahaan aagya")
	switch r.Method {

	case "GET":
		{

			fmt.Println("yeh chlra hai")
			t, err := template.ParseFiles("C:/Users/yashi/go/src/QuizPortal/templates/login.html")
			if err != nil {
				log.Fatal("Could not parse template files\n")
			}
			er := t.Execute(w, "")
			if er != nil {
				log.Fatal("could not execute the files\n")
			}
		}
		log.Print("working")
	case "POST":
		{
			fmt.Println(" lets see if it works ")
			a := r.FormValue("username")
			fmt.Println("username", a)
			user := database.Finddb(cl1, a)
			if user.Username != "" {
				processloginform(cl1, user, w, r)
			} else {
				http.Redirect(w, r, "/QuizPortal/login", 302)
			}

		}

	}
}

func organizerhandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		{

			fmt.Println("yeh chlra hai")
			t, err := template.ParseFiles("C:/Users/yashi/go/src/QuizPortal/templates/orglogin.html")
			if err != nil {
				log.Fatal("Could not parse template files\n")
			}
			er := t.Execute(w, "")
			if er != nil {
				log.Fatal("could not execute the files\n")
			}
		}
		log.Print("working")
	case "POST":
		{
			fmt.Println(" lets see if it works ")
			a := r.FormValue("username")
			user := database.Findorgdb(cl2, a)
			if user.Username != "" {
				processorgloginform(cl1, user, w, r)
			} else {
				http.Redirect(w, r, "/QuizPortal/organizer", 302)
			}

		}
	}
}

func quizhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("quiz chlra hai")
	name := r.FormValue("eventname")
	Qlist := database.Findfromquizdb(cl3, name)
	qu := Quiz{
		res:       w,
		req:       r,
		Index:     0,
		Eventname: name,
	}
	var i int
	i = 0
	for i < len(Qlist) {
		qu.Q = Qlist[i]
		rendertemplate(w, "C:/Users/yashi/go/src/Qu izPortal/templates/quiz.html", qu)
		ans := r.FormValue("answer")
		if ans == qu.Q.Answer {
			i++
		}
	}

}

//dashboard ....
func dashboard(w http.ResponseWriter, r *http.Request) {
	value := authenticate.Getcookievalues(w, r)
	fmt.Println(value)
	fmt.Println("yeh chlra hai")
	t, err := template.ParseFiles("C:/Users/yashi/go/src/QuizPortal/templates/dashboard.html")
	if err != nil {
		log.Fatal("Could not parse template files\n")
	}
	er := t.Execute(w, "")
	if er != nil {

		log.Fatal("could not execute the files\n")
	}
}

//About shows the information about the portal
func About(w http.ResponseWriter, r *http.Request) {
	fmt.Println("about chlra hai")
	t, err := template.ParseFiles("C:/Users/yashi/go/src/QuizPortal/templates/dashboard.html")
	if err != nil {
		log.Fatal("Could not parse template files\n")
	}
	er := t.Execute(w, "")
	if er != nil {
		log.Fatal("could not execute the files\n")
	}
}

//Contacthandler .....
func Contacthandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("contact chlra hai")
	t, err := template.ParseFiles("C:/Users/yashi/go/src/QuizPortal/templates/contact.html")
	if err != nil {
		log.Fatal("Could not parse template files\n")
	}
	er := t.Execute(w, "")
	if er != nil {
		log.Fatal("could not execute the files\n")
	}
}

func processsignupform(cl *mongo.Collection, U database.User, w http.ResponseWriter, r *http.Request) {
	t := database.Findfromuserdb(cl, U.Username)
	if t == true {
		processloginform(cl, U, w, r)
	}
}

func processloginform(cl *mongo.Collection, U database.User, w http.ResponseWriter, r *http.Request) {
	sessionID := database.GenerateUUID()
	rr := authenticate.CreateSecureCookie(U, sessionID, w, r)
	if rr != nil {
		log.Fatal("error occured while making a secure cookie:", rr)
	}
	er := authenticate.CreateUserSession(U, sessionID, w, r)
	if er != nil {
		log.Fatal("error occured while making a session")
	}
	http.Redirect(w, r, "/QuizPortal/login/dashboard", 302)

}
func processorgloginform(cl *mongo.Collection, U database.Organizer, w http.ResponseWriter, r *http.Request) {
	sessionID := database.GenerateUUID()
	rr := authenticate.CreateSecureorgCookie(U, sessionID, w, r)
	if rr != nil {
		log.Fatal("error occured while making a secure cookie:", rr)
	}
	// er := authenticate.Createorgsession(U, sessionID, w, r)
	// if er != nil {
	// 	log.Fatal("error occured while making a session")
	// }
	http.Redirect(w, r, "/QuizPortal/login/dashboard", 302)

}

//EVENTHANDLER HANDLES THE TEMPLATE FOR EVENTS SELECTION
func eventhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("event chlra hai")
	eve = Events{
		res:   w,
		req:   r,
		Index: 1,
		Elist: database.Findfromeventdb(cl3),
	}
	fmt.Println(eve.Elist)
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
	cl1, cl2, cl3,cl4, c = database.Createdb()
	q:=database.Quizz{
		Event: "firstname",
		Title: "Question 1",
		Description: "what is your name",
		Answer: "Yashi",
	}
	database.Insertintoquizdb(cl4,q)
	q1:=database.Quizz{
		Event: "firstname",
		Title: "Question 2",
		Description: "My favourite Book",
		Answer: "Me Before You",
	}
	database.Insertintoquizdb(cl4,q1)
}

//Checksession checks  if session is present or not
func (e Events) Checksession() bool {
	res, err := authenticate.ReadSecureCookieValues(e.res, e.req)
	if res != nil && err == nil {
		return true
	}
	return false
}

//Usertype defines thetype of user
func (e Events) Usertype() string {
	res, err := authenticate.ReadSecureCookieValues(e.res, e.req)
	if err != nil {
		st := res["username"]
		fmt.Println(st)
		x := database.Findfromuserdb(cl1, st)
		y := database.Findfromorganizerdb(cl2, st)
		if x == true {
			return "user"
		} else if y == true {
			return "organizer"
		} else {
			return "null"
		}
	}
	return "null"
}

//Checksession ....
func (q Quiz) Checksession() bool {
	res, err := authenticate.ReadSecureCookieValues(q.res, q.req)
	if res != nil && err == nil {
		return true
	}
	return false
}

//Usertype defines thetype of user
func (q Quiz) Usertype() string {
	res, err := authenticate.ReadSecureCookieValues(q.res, q.req)
	if err != nil {
		st := res["username"]
		fmt.Println(st)
		x := database.Findfromuserdb(cl1, st)
		y := database.Findfromorganizerdb(cl2, st)
		if x == true {
			return "user"
		} else if y == true {
			return "organizer"
		} else {
			return "null"
		}
	}
	return "null"
}

func rendertemplate(w http.ResponseWriter, name string, data interface{}) {
	t, err := template.ParseFiles(name)
	if err != nil {
		log.Fatal("Could not parse template files\n")
	}
	er := t.Execute(w, data)
	if er != nil {
		log.Fatal("could not execute the files\n:", er)
	}
}
