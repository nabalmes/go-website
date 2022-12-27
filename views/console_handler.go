package views

import (
	"net/http"
	"strings"

	"github.com/jmcerezo/corporate-america/models"
	"github.com/uadmin/uadmin"
)

func ConsoleHandler(w http.ResponseWriter, r *http.Request) {
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

	if path == "console" {
		path = "console/dashboard"
	}

	path = "templates/" + path + ".html"
	tmplList = append(tmplList, "templates/console/base.html")
	tmplList = append(tmplList, path)

	uadmin.Trail(uadmin.DEBUG, "TMPL: %v", tmplList)

	// * Tag
	tag := []models.Tag{}
	uadmin.All(&tag)

	for i := range tag {
		uadmin.Preload(&tag[i])
	}

	data["Tag"] = tag

	// * Maintenance
	type Maintenance struct {
		ID        uint
		Tag       models.Tag
		TagID     uint
		Item      string
		Type      string
		TimeStart string
		TimeEnd   string
	}

	m := []Maintenance{}
	maintenance := []models.Maintenance{}
	uadmin.All(&maintenance)
	time_start, time_end := "", ""

	for i := range maintenance {
		uadmin.Preload(&maintenance[i])

		if maintenance[i].TimeStart != nil {
			time_start = maintenance[i].TimeStart.Format("01/02/2006 03:04 PM")
		}
		if maintenance[i].TimeEnd != nil {
			time_end = maintenance[i].TimeEnd.Format("01/02/2006 03:04 PM")
		}

		m = append(m, Maintenance{
			ID:        maintenance[i].ID,
			Tag:       maintenance[i].Tag,
			TagID:     maintenance[i].TagID,
			Item:      maintenance[i].Item,
			Type:      maintenance[i].Type,
			TimeStart: time_start,
			TimeEnd:   time_end,
		})
	}

	data["Maintenance"] = m

	uadmin.RenderMultiHTML(w, r, tmplList, data)
}
