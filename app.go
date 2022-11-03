package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"task-1/internal/models"
	"task-1/internal/service"
)

type AppData struct {
	SlackUsername string `json:"slackUsername"`
	Backend       bool   `json:"backend"`
	Age           int    `json:"age"`
	Bio           string `json:"bio"`
}

type app struct {
	srv service.ArithmeticSrv
}

func NewApp(srv service.ArithmeticSrv) *app {
	return &app{srv: srv}
}

func (*app) GetData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := &AppData{
		SlackUsername: "Oreoluwa10",
		Backend:       false,
		Age:           34,
		Bio:           "talk to me nice....",
	}
	json.NewEncoder(w).Encode(data)
}

func (a *app) PostData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var data models.ArithmeticReq
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid Body Request: %v", err), http.StatusBadRequest)
		return
	}
	ops, result := a.srv.Compute(data.OperationType, data.X, data.Y)
	if ops == "" && result == 0 {
		http.Error(w, fmt.Sprintf("Error Computing: %v", err), http.StatusInternalServerError)
		return
	}

	res := &models.ArithmeticRes{
		SlackUsername: "Oreoluwa10",
		Result:        result,
		OperationType: ops,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)

}
