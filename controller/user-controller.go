package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rajesh4295/user-service-go/model"
	"github.com/rajesh4295/user-service-go/service"
)

/*
 *	User controller layer to accept request from exposed API and pass it user service layer
**/

var (
	userSVC service.UserService = service.NewUserService()
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Server up and running")
}

func Signup(w http.ResponseWriter, r *http.Request) {
	var u model.Signup
	var err error
	var user *model.User

	if err = json.NewDecoder(r.Body).Decode(&u); err != nil {
		fmt.Println("errr ", err)
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	if user, err = userSVC.Signup(&u); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var u model.Login
	var err error
	var user *model.User

	if err = json.NewDecoder(r.Body).Decode(&u); err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	if user, err = userSVC.Login(&u); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, user)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["id"]
	var user *model.User
	var err error

	if user, err = userSVC.GetUserById(userId); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, user)

}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
