package api

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jmcerezo/corporate-america/models"
	"github.com/uadmin/uadmin"
)

func ScanLoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	user := uadmin.User{}
	tag := models.Tag{}

	uadmin.Get(&user, "username = ?", username)
	now := time.Now()
	user.LastLogin = &now
	uadmin.Save(&user)
	scanqrSession := uuid.NewString()

	http.SetCookie(w, &http.Cookie{
		Path:  "/",
		Name:  "username",
		Value: username,
	})
	http.SetCookie(w, &http.Cookie{
		Path:  "/",
		Name:  "session",
		Value: scanqrSession,
	})

	uadmin.Get(&tag, "machine_id = ?", username)

	if tag.TagTypeID == 1 {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}

	if tag.TagTypeID == 2 {
		http.Redirect(w, r, "/maintenance/dashboard", http.StatusSeeOther)
	}
}
