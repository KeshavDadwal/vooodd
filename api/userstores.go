package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	//"github.com/jackc/pgx/v5/pgtype"
	db "github.com/vod/db/sqlc"
	util "github.com/vod/utils"
)

type UserStore struct {
	ID              int32     `json:"id"`
	UserID          int32     `json:"user_id"`
	StoreLocationID int32     `json:"store_location_id"`
	CreatedAt       time.Time `json:"created_at"`
}
func (server *Server) createUserStore(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	userstore := UserStore{}
	err := json.NewDecoder(r.Body).Decode(&userstore)

	if err != nil {
		fmt.Println("------error1------", err)

		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "invalid JSON request",
			StatusCode: http.StatusTeapot,
		}
		util.WriteJSONResponse(w, http.StatusTeapot, jsonResponse)
		return
	}

	arg := db.CreateUserStoreRelationshipParams{
		UserID:  userstore.UserID,
		StoreLocationID: userstore.StoreLocationID,
	}

	userStoreInfo, err := server.store.CreateUserStoreRelationship(ctx,arg)
	if err != nil {
		fmt.Println("------error2------", err)
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "invalid JSON request",
			StatusCode: http.StatusTeapot,
		}
		util.WriteJSONResponse(w, http.StatusTeapot, jsonResponse)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userStoreInfo)

}
