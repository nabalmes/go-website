package views

import (
	"net/http"

	"github.com/uadmin/uadmin"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}
	data["Title"] = "Corporate America"

	// session := &http.Cookie{}
	// for _, c := range r.Cookies() {
	// 	cookie := c
	// 	if strings.Contains(cookie.Name, "session") {
	// 		session = cookie
	// 	} else {
	// 		session = nil
	// 	}
	// }
	// if session == nil {
	// 	http.Redirect(w, r, "/", http.StatusSeeOther)
	// 	return
	// }

	uadmin.RenderHTML(w, r, "./templates/dashboard.html", data)
}
