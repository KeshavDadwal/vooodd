package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	//"github.com/jackc/pgx/v5/pgtype"
	db "github.com/vod/db/sqlc"
	util "github.com/vod/utils"
)

type UnitLanguage struct {
	ID         int32 `json:"id"`
	UnitID     int32 `json:"unit_id"`
	LanguageID int32 `json:"language_id"`
}

func (server *Server) createUnitLanguage(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	unitlanguage := UnitLanguage{}
	err := json.NewDecoder(r.Body).Decode(&unitlanguage)

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

	arg := db.CreateUnitLanguageParams{
		UnitID: unitlanguage.UnitID,
		LanguageID: unitlanguage.LanguageID,
	}

	unitLanguageInfo, err := server.store.CreateUnitLanguage(ctx,arg)
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
	json.NewEncoder(w).Encode(unitLanguageInfo)

}
