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

type MallView struct {
	ID         int32       `json:"id"`
	UserID     int32       `json:"user_id"`
	MallID     int32       `json:"mall_id"`
	StoreID    int32       `json:"store_id"`
	Visited    pgtype.Date `json:"visited"`
	DeviceID   pgtype.Text `json:"device_id"`
	FromDevice pgtype.Bool `json:"from_device"`
	IsPhysical pgtype.Bool `json:"is_physical"`
	CreatedAt  time.Time   `json:"created_at"`
}

func (server *Server) createMallViews(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	mallview := MallView{}
	err := json.NewDecoder(r.Body).Decode(&mallview)

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

	arg := db.CreateMallViewParams{
		UserID: mallview.MallID,
		MallID: mallview.MallID,
		StoreID: mallview.StoreID,
		Visited: mallview.Visited,
		DeviceID: mallview.DeviceID,
		FromDevice: mallview.FromDevice,
		IsPhysical: mallview.IsPhysical,
	}

	mallViewInfo, err := server.store.CreateMallView(ctx,arg)
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
	json.NewEncoder(w).Encode(mallViewInfo)

}
