package views

import (
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
)

func MaintenanceHandler(w http.ResponseWriter, r *http.Request) {
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

	tmplList := []string{}

	path := strings.TrimPrefix(r.URL.Path, "/")
	path = strings.TrimSuffix(path, "/")

	if path == "maintenance" {
		path = "maintenance/dashboard"
	}

	path = "templates/" + path + ".html"
	tmplList = append(tmplList, "templates/maintenance/base.html")
	tmplList = append(tmplList, path)

	uadmin.Trail(uadmin.DEBUG, "TMPL: %v", tmplList)

	uadmin.RenderMultiHTML(w, r, tmplList, data)
}
