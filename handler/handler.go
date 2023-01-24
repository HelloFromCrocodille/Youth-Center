package handler

import(
	"net/http"
	"path"
	"html/template"
)

func MainPage(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/"{
		http.NotFound(w, r)
		return
	}

	var filepath = path.Join("Frontend","index.html")

	var tmpl, err = template.ParseFiles(filepath)
	if err != nil{
		http.Error(w, "Unexpected Error ! Keep Calm", http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"title":"Bagindo Aziz Chan Padang Youth Center",
	}

	err = tmpl.Execute(w, data)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}