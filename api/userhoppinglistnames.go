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
type UserShoppingListName struct {
	ID        int32       `json:"id"`
	UserID    int32       `json:"user_id"`
	Name      string      `json:"name"`
	HowObtain pgtype.Text `json:"how_obtain"`
	Frequency pgtype.Bool `json:"frequency"`
	Type      pgtype.Text `json:"type"`
	Isnotify  pgtype.Bool `json:"isnotify"`
	Created   time.Time   `json:"created"`
	Modified  time.Time   `json:"modified"`
}

func (server *Server) createUserShoppingListName(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	usershoppinglistname := UserShoppingListName{}
	err := json.NewDecoder(r.Body).Decode(&usershoppinglistname)

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

	arg := db.CreateUserShoppingListNameParams{
		Name: usershoppinglistname.Name,
		UserID: usershoppinglistname.UserID,
		HowObtain: usershoppinglistname.HowObtain,
		Frequency: usershoppinglistname.Frequency,
		Type: usershoppinglistname.Type,
		Isnotify: usershoppinglistname.Isnotify,
	}

	userShoppingListNameInfo, err := server.store.CreateUserShoppingListName(ctx,arg)
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
	json.NewEncoder(w).Encode(userShoppingListNameInfo)

}
