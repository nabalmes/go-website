package views

import (
	"net/http"

	"github.com/uadmin/uadmin"
)

func BankAccountHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}
	data["Title"] = "Corporate America"

	uadmin.RenderHTML(w, r, "./templates/bank.html", data)
}
