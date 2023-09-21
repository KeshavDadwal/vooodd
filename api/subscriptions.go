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

type Subscription struct {
	ID        int32     `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	MallID    int32     `json:"mall_id"`
	Zip       string    `json:"zip"`
	CreatedAt time.Time `json:"created_at"`
}


func (server *Server) createSubscription(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	subscription := Subscription{}
	err := json.NewDecoder(r.Body).Decode(&subscription)

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

	arg := db.CreateSubscriptionParams{
		FirstName: subscription.FirstName,
		LastName: subscription.LastName,
		Email: subscription.Email,
		MallID: subscription.MallID,
		Zip: subscription.Zip,
	}

	subscriptionInfo, err := server.store.CreateSubscription(ctx,arg)
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
	json.NewEncoder(w).Encode(subscriptionInfo)

}



