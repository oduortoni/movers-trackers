package server

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"farmers/datatypes"
	"farmers/db"
)

func StartServer(port int) {
	addr := ":" + strconv.Itoa(port)
	log.Println("Server running on http://localhost" + addr)

	http.HandleFunc("/", home)
	http.HandleFunc("/register", register)
	http.HandleFunc("/users", users)
	http.HandleFunc("/updatecharges", update_charges)
	http.HandleFunc("/viewcharges", view_charges)
	http.HandleFunc("/usercharges", user_charges)
	http.ListenAndServe(addr, nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/home.html"))
	tmpl.Execute(w, nil)
}

func register(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/register.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	member := datatypes.Member{
		Group:    r.FormValue("membership"),
		Location: r.FormValue("location"),
		Produce:  r.FormValue("produce"),
	}

	// all farmers in a group to be registered
	var farmers []datatypes.Farmer

	// large scale producer is not a group
	if r.FormValue("firmsize") == "single" {
		farmer := datatypes.Farmer{
			First:  r.FormValue("fname"),
			Second: r.FormValue("sname"),
		}
		farmers = append(farmers, farmer)
	} else {
		// implement group insertion
	}
	db.InsertMember(member, farmers)

	tmpl.Execute(w, struct{ Success bool }{true})
}

func users(w http.ResponseWriter, r *http.Request) {
	users, _ := db.GetUsers()
	tmpl := template.Must(template.ParseFiles("views/users.html"))
	tmpl.Execute(w, users)
}
