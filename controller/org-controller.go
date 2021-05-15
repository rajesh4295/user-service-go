package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rajesh4295/user-service-go/model"
	"github.com/rajesh4295/user-service-go/service"
)

/*
 *	Org controller layer to accept request from exposed API and pass it org service layer
**/

var (
	orgSVC service.OrgService = service.NewOrgService()
)

func GetOrgById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	orgId := params["id"]
	var org *model.Org
	var err error

	if org, err = orgSVC.GetOrgById(orgId); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, org)

}
