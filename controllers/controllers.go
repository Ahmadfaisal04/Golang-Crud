package controllers

import (
	"html/template"
	"net/http"
	"pasien/entities"
	"pasien/models"
	"strconv"
)

// Read Data
func Index(w http.ResponseWriter, r *http.Request) {
	pasiens := models.GetAll()

	data := map[string]any{
		"pasiens": pasiens,
	}

	temp, err := template.ParseFiles("views/index.html")

	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/create.html")

		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var pasien entities.User

		pasien.NoIdentitas = r.FormValue("noIdentitas")
		pasien.Nama = r.FormValue("nama")
		pasien.JenisKelamin = r.FormValue("jenisKelamin")
		pasien.Alamat = r.FormValue("alamat")
		pasien.NoHP = r.FormValue("noHP")

		if ok := models.Create(pasien); !ok {
			temp, err := template.ParseFiles("views/create.html")

			if err != nil {
				panic(err)
			}

			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)

		if err != nil {
			panic(err)
		}

		pasien := models.Details(id)

		data := map[string]any{
			"pasien": pasien,
		}

		temp, err := template.ParseFiles("views/edit.html")

		if err != nil {
			panic(err)
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		//idString := r.URL.Query().Get("id")
		//ids := r.FormValue("id")

		//fmt.Print("idStr = ", ids)
		id, err := strconv.Atoi(r.FormValue("id"))
		//fmt.Print("id = ", reflect.TypeOf(ids))

		if err != nil {
			panic(err)
		}

		var pasien entities.User

		pasien.NoIdentitas = r.FormValue("noIdentitas")
		pasien.Nama = r.FormValue("nama")
		pasien.JenisKelamin = r.FormValue("jenisKelamin")
		pasien.Alamat = r.FormValue("alamat")
		pasien.NoHP = r.FormValue("noHP")

		if ok := models.Update(id, pasien); !ok {
			temp, err := template.ParseFiles("views/update.html")

			if err != nil {
				panic(err)
			}

			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		panic(err)
	}

	if err := models.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/create", http.StatusSeeOther)
}
