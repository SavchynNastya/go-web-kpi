package auth

import (
	"go-web-kpi/app/common"
	"html/template"
	"net/http"
)

func GetLoginPage(rw http.ResponseWriter, req *http.Request) {
	type Page struct {
		Title string
	}

	p := Page{
		Title: "login",
	}

	common.Templates = template.Must(template.ParseFiles("templates/auth/login.html", common.LayoutPath))
	err := common.Templates.ExecuteTemplate(rw, "base", p)
	common.CheckError(err, 2)
}
