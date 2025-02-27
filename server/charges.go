package server

import (
	"html/template"
	"net/http"
	"strconv"

	"farmers/datatypes"
	"farmers/db"
	"farmers/files"
	"farmers/geo"
)

func user_charges(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/usercharges.html"))

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	details := datatypes.UserCharges{
		Group: r.FormValue("group"),
		From:  r.FormValue("from"),
		To:    r.FormValue("to"),
	}

	/**
	*CHECK IF GROUP IS REGISTERED
	*exists_group = db.ExistsGroup(detaiils.Group)
	*if exists_group == false {
	*   tmpl.Execute(w, datatypes.ChargesGroup{true, nil, false})
	*  return
	*}
	 */

	load, _ := strconv.Atoi(r.FormValue("load"))
	details.Load = load

	if details.Load < 3 {
		details.Means = "Pick-up"
	} else if details.Load >= 3 && details.Load < 7 {
		details.Means = "Refrigerated truck"
	} else if details.Load > 7 && details.Load <= 10 {
		details.Means = "Lorry"
	} else {
		details.Means = "Trailer"
	}
	details.CostPerDistance = GetCostPerDistance(details.Means)
	details.Distance = geo.GetDistance(details.From, details.To)
	details.TotalCost = details.Distance * details.CostPerDistance

	tmpl.Execute(w, datatypes.ChargesGroup{true, details, true})
}

func GetCostPerDistance(means string) int {
	charges, err := db.GetCharges()
	files.CheckError(err)

	for _, charge := range charges {
		if charge.Means == means {
			return charge.Cost
		}
	}
	return 0
}

func view_charges(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/viewcharges.html"))
	charges, _ := db.GetCharges()
	tmpl.Execute(w, charges)
}

func update_charges(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/updatecharges.html"))

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	load, _ := strconv.Atoi(r.FormValue("load"))
	cost, _ := strconv.Atoi(r.FormValue("cost"))

	charges := datatypes.Charges{
		Means: r.FormValue("means"),
		Load:  load,
		Cost:  cost,
	}

	db.UpdateCharges(charges)

	tmpl.Execute(w, struct{ Success bool }{true})
}
