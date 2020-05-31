package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/nagahshi/bankApi/helpers"
	"github.com/nagahshi/bankApi/users"

	"github.com/gorilla/mux"
)

type Login struct {
	Username string
	Password string
}

type ErrResponse struct {
	Message string
}

func login(w http.ResponseWriter, r *http.Request)  {
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)

	var formatedBody Login
	err = json.Unmarshal(body, &formatedBody)
	helpers.HandleErr(err)

	login := users.Login(formatedBody.Username, formatedBody.Password)

	if login["message"] == "tudo certo!" {
		resp := login
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := ErrResponse{Message: "Credenciais inválidas"}
		json.NewEncoder(w).Encode(resp)
	}
}

func StartupApi() {
	router := mux.NewRouter()

	router.HandleFunc("/login", login).Methods("POST")

	fmt.Println("API is working on port :8882")
	log.Fatal(http.ListenAndServe(":8882", router))
}