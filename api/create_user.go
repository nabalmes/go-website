package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/jmcerezo/corporate-america/models"
	"github.com/uadmin/uadmin"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	uid := r.FormValue("uid")
	username := r.FormValue("username")
	password := r.FormValue("password")
	first_name := r.FormValue("first_name")
	last_name := r.FormValue("last_name")
	contact, _ := strconv.Atoi(r.FormValue("contact"))

	user := uadmin.User{}
	tag := models.Tag{}
	existingTags := []models.Tag{}
	uadmin.All(&existingTags)
	isExisting := false

	uadmin.Trail(uadmin.DEBUG, "UID: %v", uid)
	resp, err := http.Get("https://repo.pudding.ws/tags/?serial=" + uid)
	if err != nil {
		uadmin.Trail(uadmin.ERROR, "Error fetching serial. %v", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		uadmin.Trail(uadmin.ERROR, "Error reading response body. %v", err)
	}
	x := map[string]string{}
	json.Unmarshal([]byte(body), &x)

	if x["Result"] != "" {
		// * Validation for existing tags
		for i := range existingTags {
			if existingTags[i].MachineID == tag.MachineID {
				isExisting = true
				break
			}
		}
		if !isExisting {
			user.Username = username
			user.Password = password
			user.FirstName = first_name
			user.LastName = last_name
			user.Active = true
			user.Save()

			tag.MachineID = uid
			tag.Enc = x["Result"]
			tag.UserID = user.ID
			tag.ContactNumber = contact
			tag.Save()

			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}
