package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	//"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/vod/db/sqlc"
	util "github.com/vod/utils"
)

type StoreLocationHoliday struct {
	ID              int32       `json:"id"`
	StoreLocationID int32       `json:"store_location_id"`
	Name            string      `json:"name"`
	StartDate       pgtype.Date `json:"start_date"`
	EndDate         pgtype.Date `json:"end_date"`
	CreatedAt       time.Time   `json:"created_at"`
}

func (server *Server) createStoreLocationHolidays(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed, "Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	storelocationholiday := StoreLocationHoliday{}
	err := json.NewDecoder(r.Body).Decode(&storelocationholiday)

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

	arg := db.CreateStoreLocationHolidayParams{
		StoreLocationID: storelocationholiday.StoreLocationID,
		Name:            storelocationholiday.Name,
		StartDate:       storelocationholiday.StartDate,
		EndDate:         storelocationholiday.EndDate,
	}

	storeLocationHolidayInfo, err := server.store.CreateStoreLocationHoliday(ctx, arg)
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
	json.NewEncoder(w).Encode(storeLocationHolidayInfo)

}
