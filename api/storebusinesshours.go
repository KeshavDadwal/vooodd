package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/vod/db/sqlc"
	util "github.com/vod/utils"
)

type StoreBusinessHour struct {
	ID              int32       `json:"id"`
	StoreLocationID int32       `json:"store_location_id"`
	Day             string      `json:"day"`
	StartTime       pgtype.Date `json:"start_time"`
	EndTime         pgtype.Date `json:"end_time"`
	CreatedAt       time.Time   `json:"created_at"`
}


func (server *Server) createStoreBusinessHours(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	storebusinesshour := StoreBusinessHour{}
	err := json.NewDecoder(r.Body).Decode(&storebusinesshour)

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

	arg := db.CreateStoreBusinessHoursParams{
		StoreLocationID: storebusinesshour.StoreLocationID,
		Day: storebusinesshour.Day,
		StartTime: storebusinesshour.StartTime,
		EndTime: storebusinesshour.EndTime,
	}

	storeBusinessHourInfo, err := server.store.CreateStoreBusinessHours(ctx,arg)
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
	json.NewEncoder(w).Encode(storeBusinessHourInfo)

}



