package user

import(
	"net/http"
	"path"
	"html/template"
)

func RegistrationPage(w http.ResponseWriter, r *http.Request){
	var filepath = path.Join("frontend","user","registration.html")

	var tmpl, err = template.ParseFiles(filepath)
	if err != nil{
		http.Error(w, "File not Found ! Keep Calm", http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"text":"Registration | Bagindo Aziz Chan Padang Youth Center",
	}

	err = tmpl.Execute(w, data)
	if err != nil{
		http.Error(w, "Failed to Execute", http.StatusInternalServerError)
	}
}