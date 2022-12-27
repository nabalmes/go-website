package main

import (
	"net/http"

	"github.com/jmcerezo/corporate-america/api"
	"github.com/jmcerezo/corporate-america/models"
	"github.com/jmcerezo/corporate-america/views"
	"github.com/uadmin/uadmin"
)

func main() {
	DatabaseConfig()
	RegisterModels()
	SampleData()
	RegisterHandlers()
	StartServer()
}

func RegisterModels() {
	uadmin.Trail(uadmin.INFO, "Register Models")
	uadmin.Register(
		models.Tag{},
		models.TagType{},
		models.Maintenance{},
	)
}

func DatabaseConfig() {
	uadmin.Trail(uadmin.INFO, "Database Config")
	uadmin.Database = &uadmin.DBSettings{
		Type:     "mysql",
		Host:     "localhost",
		Name:     "corporate_america",
		User:     "root",
		Password: "Allen is Great 200%",
		Port:     3306,
	}
}

func RegisterHandlers() {
	uadmin.Trail(uadmin.INFO, "Register Handlers")
	http.HandleFunc("/", uadmin.Handler(views.IndexHandler))
	http.HandleFunc("/api/", uadmin.Handler(api.APIHandler))
	http.HandleFunc("/logout", uadmin.Handler(views.LogoutHandler))
	http.HandleFunc("/dashboard", uadmin.Handler(views.DashboardHandler))
	http.HandleFunc("/console/", uadmin.Handler(views.ConsoleHandler))
	http.HandleFunc("/maintenance/", uadmin.Handler(views.MaintenanceHandler))
	http.HandleFunc("/bankaccount/", uadmin.Handler(views.BankAccountHandler))
}

func StartServer() {
	uadmin.Port = 2930
	uadmin.RootURL = "/admin/"
	uadmin.StartServer()
}
