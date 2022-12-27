package views

import (
	"net/http"
	"strings"

	"github.com/jmcerezo/corporate-america/models"
	"github.com/uadmin/uadmin"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}
	data["Title"] = "Corporate America"

	session := uadmin.IsAuthenticated(r)
	sessionCookie, _ := r.Cookie("session")
	usernameCookie, _ := r.Cookie("username")
	user := uadmin.User{}
	tag := models.Tag{}

	if session != nil || sessionCookie != nil {
		if usernameCookie.Value == "admin" {
			http.Redirect(w, r, "/console/dashboard", http.StatusSeeOther)
		}

		uadmin.Get(&tag, "machine_id = ?", usernameCookie.Value)

		if tag.TagTypeID == 1 {
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		}

		if tag.TagTypeID == 2 {
			http.Redirect(w, r, "/maintenance/dashboard", http.StatusSeeOther)
		}
	}

	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		uadmin.Get(&user, "username = ?", strings.ToLower(username))
		newSession := user.Login(password, "")

		if newSession != nil {
			http.SetCookie(w, &http.Cookie{
				Path:  "/",
				Name:  "username",
				Value: username,
			})
			http.SetCookie(w, &http.Cookie{
				Path:  "/",
				Name:  "session",
				Value: newSession.Key,
			})

			if username == "admin" {
				http.Redirect(w, r, "/console/dashboard", http.StatusSeeOther)
			}

			t := models.Tag{}
			uadmin.Get(&t, "machine_id = ?", username)

			if t.TagTypeID == 1 {
				http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
			}

			if tag.TagTypeID == 2 {
				http.Redirect(w, r, "/maintenance/dashboard", http.StatusSeeOther)
			}
		}
	}

	uadmin.RenderHTML(w, r, "./templates/index.html", data)
}
