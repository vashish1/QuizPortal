package main

import (
	"QuizPortal/authenticate"
	"QuizPortal/database"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	r := NewRouter()

	// r.HandleFunc("/QuizPortal/", About)
	r.HandleFunc("/QuizPortal/signup", signuphandler).Methods("GET", "POST")
	r.HandleFunc("/QuizPortal/login", loginhandler).Methods("GET", "POST")
	r.HandleFunc("/QuizPortal/orglogin", loginhandler).Methods("GET", "POST")
	r.HandleFunc("/QuizPortal/login/dashboard", dashboard).Methods("GET")
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
	cl1, _, _ := database.Createdb()
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
	cl1, _, _ := database.Createdb()
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
			user := database.Finddb(cl1, a)
			processloginform(cl1, user, w, r)

		}

	}
}

func organizerhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("yahaan aagya")
	cl1, _, _ := database.Createdb()
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
			b := r.FormValue("password")

			fmt.Println(a, b)
			if database.Findfromuserdb(cl1, a) {
				fmt.Println("Access granted")
			} else {
				fmt.Println("Access Denied")
			}

		}

		http.Redirect(w, r, "/login/dashboard", 302)

	}
}

func dashboard(w http.ResponseWriter, r *http.Request) {
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

// //About shows the information about the portal
// func About(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("about chlra hai")

// }

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
