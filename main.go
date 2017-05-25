package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/otherpirate/dbaas-metric-collector/collector"
	"github.com/otherpirate/dbaas-metric-collector/model"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("page")))
	http.HandleFunc("/loading", loading)
	http.HandleFunc("/database_count", database_count)
	http.HandleFunc("/team_count", team_count)
	http.HandleFunc("/project_count", project_count)
	http.HandleFunc("/environment_count", environment_count)

	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func loading(res http.ResponseWriter, req *http.Request) {
	collector.GetDatabases()
	fmt.Fprintln(res, "Loaded!")
}

func database_count(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	connection := model.GetConnection()
	defer connection.Session.Close()

	counters := model.DatabaseCounterGet(connection)
	content, err := json.Marshal(counters)
	if err != nil {
		panic(err)
	}

	res.Write(content)
}

func team_count(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	connection := model.GetConnection()
	defer connection.Session.Close()

	counters := model.TeamCounterGetLatest(connection)
	content, err := json.Marshal(counters)
	if err != nil {
		panic(err)
	}

	res.Write(content)
}

func project_count(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	connection := model.GetConnection()
	defer connection.Session.Close()

	counters := model.ProjectCounterGetLatest(connection)
	content, err := json.Marshal(counters)
	if err != nil {
		panic(err)
	}

	res.Write(content)
}

func environment_count(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	connection := model.GetConnection()
	defer connection.Session.Close()

	counters := model.EnvironmentCounterGetLatest(connection)
	content, err := json.Marshal(counters)
	if err != nil {
		panic(err)
	}

	res.Write(content)
}
