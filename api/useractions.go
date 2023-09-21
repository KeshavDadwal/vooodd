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
type UserAction struct {
	ID              int32       `json:"id"`
	UserID          int32       `json:"user_id"`
	SaveType        string      `json:"save_type"`
	Tablename       string      `json:"tablename"`
	AutoincrementID pgtype.Int4 `json:"autoincrement_id"`
	Fieldname       string      `json:"fieldname"`
	Value           string      `json:"value"`
	OldValue        pgtype.Text `json:"old_value"`
	Note            string      `json:"note"`
	CreatedAt       time.Time   `json:"created_at"`
}



func (server *Server) createUserAction(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	useraction := UserAction{}
	err := json.NewDecoder(r.Body).Decode(&useraction)

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

	arg := db.CreateUserActionParams{
		UserID: useraction.UserID,
		SaveType: useraction.SaveType,
		Tablename: useraction.Tablename,
		AutoincrementID: useraction.AutoincrementID,
		Fieldname: useraction.Fieldname,
		Value: useraction.Value,
		OldValue: useraction.OldValue,
		Note: useraction.Note,
	}

	userActionInfo, err := server.store.CreateUserAction(ctx,arg)
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
	json.NewEncoder(w).Encode(userActionInfo)

}



