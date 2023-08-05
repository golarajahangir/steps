package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type team struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Counter int      `json:"counter"`
	Members []string `json:"members"`
}

type updateScoreRequest struct {
	Counter int `json:"counter"`
}

var teams = []team{
	{ID: "1", Name: "Team 1", Counter: 0, Members: []string{"Alice", "Bob", "Carol"}},
	{ID: "2", Name: "Team 2", Counter: 0, Members: []string{"Teo", "Tony"}},
	{ID: "3", Name: "Team 3", Counter: 0, Members: []string{"Alexander", "Sara"}},
}

func getAllTeams(q http.ResponseWriter, r *http.Request) {
	q.Header().Set("Content-Type", "application/json")
	json.NewEncoder(q).Encode(teams)
}
func getTeam(q http.ResponseWriter, r *http.Request) {
	var idparam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idparam)
	if err != nil {
		q.WriteHeader(400)
		q.Write([]byte("ID could not converted to integer"))
		return
	}

	if id >= len(teams) {
		q.WriteHeader(404)
		q.Write([]byte("NO Team find whit this id"))
	}
	team := teams[id]
	q.Header().Set("Content-Type", "application/json")
	json.NewEncoder(q).Encode(team)
}

func updateScore(q http.ResponseWriter, r *http.Request) {
	var idparam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idparam)
	if err != nil {
		q.WriteHeader(400)
		q.Write([]byte("ID could not converted to integer"))
		return
	}

	if id >= len(teams) {
		q.WriteHeader(404)
		q.Write([]byte("NO Team find whit this id"))
	}
	var updateScore updateScoreRequest
	json.NewDecoder(r.Body).Decode(&updateScore)
	teams[id].Counter = teams[id].Counter + updateScore.Counter
	q.Header().Set("Content-Type", "application/json")
	json.NewEncoder(q).Encode(updateScore)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/teams", getAllTeams).Methods("GET")
	router.HandleFunc("/teams/{id}", getTeam).Methods("GET")
	router.HandleFunc("/teams/{id}", updateScore).Methods("PUT")
	http.ListenAndServe(":5000", router)
}
