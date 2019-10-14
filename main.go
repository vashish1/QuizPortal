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

var cl1, cl2, cl3, cl4 *mongo.Collection
var c *mongo.Client
var eve Events
var e empty
var i int

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
	r.HandleFunc("/QuizPortal/organizer/dashboard", orgdashboard).Methods("GET", "POST")
	r.HandleFunc("/QuizPortal/login/dashboard", dashboard).Methods("GET")
	r.HandleFunc("/QuizPortal/contact", Contacthandler)
	r.HandleFunc("/QuizPortal/events", eventhandler).Methods("GET", "POST")
	r.HandleFunc("/quiz/", quizhandler).Methods("GET", "POST")
	r.HandleFunc("/QuizPortal/update/", updateevent).Methods("GET", "POST")
	r.HandleFunc("/QuizPortal/deleteEvent/", deleteevent).Methods("GET", "POST")
	r.HandleFunc("/QuizPortal/addEvent", addevent).Methods("GET", "POST")
	r.HandleFunc("/QuizPortal/logout", logout).Methods("POST")
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}

//NewRouter .....
func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	return r
}

func logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("logout chlra hai")
	authenticate.ExpireSecureCookie(w, r)
	authenticate.ExpireUserSession(w, r)
	http.Redirect(w, r, "/QuizPortal/login", 302)
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
	fmt.Println("login mein aagya")
	x, _ := authenticate.ReadSecureCookieValues(w, r)
	if x == nil {
		em := empty{
			res: w,
			req: r,
		}

		switch r.Method {

		case "GET":
			{

				fmt.Println("yeh chlra hai")
				t, err := template.ParseFiles("C:/Users/yashi/go/src/QuizPortal/templates/login.html")
				if err != nil {
					log.Fatal("Could not parse template files\n")
				}
				er := t.Execute(w, em)
				if er != nil {
					log.Fatal("could not execute the files\n")
				}
			}
			log.Print("working")
		case "POST":
			{
				fmt.Println(" lets see if it works ")
				a := r.FormValue("username")
				b := r.FormValue("password")
				fmt.Println("username", a)
				user := database.Finddb(cl1, a)
				if user.PasswordHash == database.SHA256ofstring(b) {
					processloginform(cl1, user, w, r)
				} else {
					http.Redirect(w, r, "/QuizPortal/login", 302)
				}

			}

		}

	} else {
		http.Redirect(w, r, "/QuizPortal/login/dashboard", 302)
	}

}

func organizerhandler(w http.ResponseWriter, r *http.Request) {

	e = empty{
		res: w,
		req: r,
	}

	switch r.Method {

	case "GET":
		{

			fmt.Println("yeh chlra hai")
			t, err := template.ParseFiles("C:/Users/yashi/go/src/QuizPortal/templates/orglogin.html")
			if err != nil {
				log.Fatal("Could not parse template files\n")
			}
			er := t.Execute(w, e)
			if er != nil {
				log.Fatal("could not execute the files\n")
			}
		}
		log.Print("working")
	case "POST":
		{
			fmt.Println(" lets see if it works ")
			a := r.FormValue("username")
			b := r.FormValue("password")
			user := database.Findorgdb(cl2, a)
			if user.PasswordHash == database.SHA256ofstring(b) {
				processorgloginform(cl1, user, w, r)
			} else {
				http.Redirect(w, r, "/QuizPortal/organizer", 302)
			}

		}
	}
}

func quizhandler(w http.ResponseWriter, r *http.Request) {
    
	
	fmt.Println("quiz chlra hai")
	var qu Quiz
	var Qlist []database.Quizz   
    var ans string
	switch r.Method {
	case "POST":
		fmt.Println("1")
		re := r.FormValue("eventname")
		fmt.Println("eventname:", re)
		fmt.Println("2")
		Qlist = database.Findfromquizdb(cl4, re)
		fmt.Println("3")
		qu = Quiz{
			res:       w,
			req:       r,
			Index:     0,
			Eventname: re,
		}
		fmt.Println("5")
		ans = r.FormValue("answer")
		fmt.Println(ans)
		fmt.Println("6")
	case "GET" :
		 qu.Q=Qlist[i]
		 t, err := template.ParseFiles("C:/Users/yashi/go/src/QuizPortal/templates/quiz.html")
			if err != nil {
				log.Fatal("Could not parse template files:", err)
			}
			er := t.Execute(w, qu)
			if er != nil {
				log.Fatal("could not execute the files\n:", er)
			}

		 if ans == qu.Q.Answer {
			i++
			fmt.Println(i)
			fmt.Println("7")
			http.Redirect(w, r, "/quiz/", 302)
			fmt.Println("8")
		}
}
}

//dashboard ....
func dashboard(w http.ResponseWriter, r *http.Request) {
	value := authenticate.Getcookievalues(w, r)
	fmt.Println(value)
	em := empty{
		res: w,
		req: r,
	}
	fmt.Println("yeh chlra hai")
	t, err := template.ParseFiles("C:/Users/yashi/go/src/QuizPortal/templates/dashboard.html")
	if err != nil {
		log.Fatal("Could not parse template files\n")
	}
	er := t.Execute(w, em)
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
	em := empty{
		res: w,
		req: r,
	}

	t, err := template.ParseFiles("C:/Users/yashi/go/src/QuizPortal/templates/contact.html")
	if err != nil {
		log.Fatal("Could not parse template files\n")
	}
	er := t.Execute(w, em)
	if er != nil {
		log.Fatal("could not execute the files\n")
	}
}

func processsignupform(cl *mongo.Collection, U database.User, w http.ResponseWriter, r *http.Request) {
	database.Insertintouserdb(cl1, U)
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
	er := authenticate.CreateorgSession(U, sessionID, w, r)
	if er != nil {
		log.Fatal("error occured while making a session")
	}
	http.Redirect(w, r, "/QuizPortal/organizer/dashboard", 302)

}

//EVENTHANDLER HANDLES THE TEMPLATE FOR EVENTS SELECTION
func eventhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("event chlra hai")
	eve = Events{
		res: w,
		req: r,

		Elist: database.Findfromeventdb(cl3),
	}
	fmt.Println(eve.Elist)
	t, err := template.ParseFiles("C:/Users/yashi/go/src/QuizPortal/templates/events.html")
	if err != nil {
		log.Fatal("Could not parse template files:", err)
	}
	er := t.Execute(w, eve)
	if er != nil {
		log.Fatal("could not execute the files\n:", er)
	}

}
func init() {
	cl1, cl2, cl3, cl4, c = database.Createdb()
}

//Checksession checks  if session is present or not
func (e Events) Checksession() bool {
	res, err := authenticate.ReadSecureCookieValues(e.res, e.req)
	if res != nil && err == nil {
		return true
	}
	return false
}

//Checksession checks  if session is present or not
func (e Myevents) Checksession() bool {
	res, err := authenticate.ReadSecureCookieValues(e.res, e.req)
	if res != nil && err == nil {
		return true
	}
	return false
}

//Usertype defines thetype of user
func (e Events) Usertype() bool {
	res, err := authenticate.ReadSecureCookieValues(e.res, e.req)
	fmt.Println(res)
	if err == nil {
		st := res["username"]
		fmt.Println(st)
		y := database.Findfromorganizerdb(cl2, st)

		return y
	}

	fmt.Println("error while secure cookie read is :", err)
	return false

}

//Checksession ....
func (q Quiz) Checksession() bool {
	res, err := authenticate.ReadSecureCookieValues(q.res, q.req)
	if res != nil && err == nil {
		return true
	}
	return false
}

//Checksession ....
func (em empty) Checksession() bool {
	res, err := authenticate.ReadSecureCookieValues(em.res, em.req)
	if res != nil && err == nil {
		return true
	}
	return false
}

//Usertype defines thetype of user
func (q Quiz) Usertype() bool {
	res, err := authenticate.ReadSecureCookieValues(q.res, q.req)
	if err == nil {
		st := res["username"]
		fmt.Println(st)
		y := database.Findfromorganizerdb(cl2, st)

		return y
	}
	return false
}

//Usertype defines thetype of user
func (em empty) Usertype() bool {
	res, err := authenticate.ReadSecureCookieValues(em.res, em.req)
	if err == nil {
		st := res["username"]
		fmt.Println(st)
		y := database.Findfromorganizerdb(cl2, st)

		return y
	}
	return false
}

// func rendertemplate(w http.ResponseWriter, name string, data Quiz) {
// 	t, err := template.ParseFiles(name)
// 	if err != nil {
// 		log.Fatal("Could not parse template files:", err)
// 	}
// 	er := t.Execute(w, data)
// 	if er != nil {
// 		log.Fatal("could not execute the files\n:", er)
// 	}
// }

func orgdashboard(w http.ResponseWriter, r *http.Request) {
	x := findusername(w, r)
	fmt.Println("x", x)
	my := Myevents{
		res:      w,
		req:      r,
		username: x,
		List:     findorgevents(x),
	}
	fmt.Println("my events list:", my.List)

	t, err := template.ParseFiles("C:/Users/yashi/go/src/QuizPortal/templates/orgdashboard.html")
	if err != nil {
		log.Fatal("Could not parse template files:", err)
	}
	er := t.Execute(w, my)
	if er != nil {
		log.Fatal("could not execute the files\n:", er)
	}
}

func findusername(w http.ResponseWriter, r *http.Request) string {

	res, err := authenticate.ReadSecureCookieValues(w, r)
	if err != nil {

		fmt.Println("the error while finding username:", err)
	}

	st := res["username"]
	return st
}

func findorgevents(s string) []database.Event {
	var result []database.Event
	org := database.Findorgdb(cl2, s)
	fmt.Println("org database:", org)

	var i int
	for i = 0; i < len(org.Events); i++ {
		eve := org.Events[i]
		fmt.Println("eve:", eve)
		e := database.Findevent(cl3, eve)
		fmt.Println("event:", e)
		if e.Eventdescription != "" {
			result = append(result, e)
		}

	}
	return result
}

func addevent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("event add")
	em := empty{
		res: w,
		req: r,
	}

	switch r.Method {

	case "GET":
		{

			fmt.Println("event adiition function  chlra hai")
			t, err := template.ParseFiles("C:/Users/yashi/go/src/QuizPortal/templates/addevent.html")
			if err != nil {
				log.Fatal("Could not parse template files:", err)
			}
			er := t.Execute(w, em)
			if er != nil {
				log.Fatal("could not execute the files\n")
			}
		}
		log.Print("working")
	case "POST":
		{
			fmt.Println(" lets see if it works ")
			a := r.FormValue("name")

			b := r.FormValue("description")
			c := r.FormValue("startdate")
			d := r.FormValue("enddate")
			e := r.FormValue("starttime")
			f := r.FormValue("endtime")
			fmt.Println(a, b, c, d, e, f)

			u := database.NewEvent(a, b, c, d, e, f)
			database.Insertintoeventdb(cl3, u)
			fmt.Println("Event inserted:", u)
			s := findusername(w, r)
			database.Updateorg(cl2, s, a)
			http.Redirect(w, r, "/QuizPortal/organizer/dashboard", 302)
		}
	}

}

func deleteevent(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		re := r.FormValue("eventname")
		fmt.Println("eventname:", re)
		database.Deleteevent(cl3, re)
	}

	http.Redirect(w, r, "/QuizPortal/organizer/dashboard", 302)
}

func updateevent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("question add")
	em := empty{
		res: w,
		req: r,
	}

	switch r.Method {

	case "GET":
		{

			fmt.Println("question adiition function  chlra hai")
			t, err := template.ParseFiles("C:/Users/yashi/go/src/QuizPortal/templates/addquestion.html")
			if err != nil {
				log.Fatal("Could not parse template files:", err)
			}
			er := t.Execute(w, em)
			if er != nil {
				log.Fatal("could not execute the files\n")
			}
		}
		log.Print("working")
	case "POST":
		{
			fmt.Println(" lets see if it works ")
			re := r.FormValue("eventname")
			fmt.Println("eventname:", re)
			t := r.FormValue("title")
			q := r.FormValue("question")
			a := r.FormValue("answer")
			Q := database.Quizz{
				Event:       re,
				Title:       t,
				Description: q,
				Answer:      a,
			}
			database.Insertintoquizdb(cl4, Q)
			http.Redirect(w, r, "/QuizPortal/update/", 302)
		}
	}
}
