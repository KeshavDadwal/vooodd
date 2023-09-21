package api

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/vod/db/sqlc"
	util "github.com/vod/utils"
)
type ImportFlow struct {
	ID            int32       `json:"id"`
	FileName      string      `json:"file_name"`
	NoRows        pgtype.Int4 `json:"no_rows"`
	Created       time.Time   `json:"created"`
	Modified      time.Time   `json:"modified"`
	Status        string      `json:"status"`
	NoError       pgtype.Int4 `json:"no_error"`
	Note          string      `json:"note"`
	LogFile       pgtype.Text `json:"log_file"`
	CompletedFile pgtype.Text `json:"completed_file"`
	CreatedAt     time.Time   `json:"created_at"`
}

func (server *Server) createImportFlow(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	importflow := ImportFlow{}
	err := json.NewDecoder(r.Body).Decode(&importflow)

	if err != nil {

		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "invalid JSON request",
			StatusCode: http.StatusTeapot,
		}
		util.WriteJSONResponse(w, http.StatusTeapot, jsonResponse)
		return
	}

	arg := db.CreateImportFlowParams{
		FileName: importflow.FileName,
		NoRows: importflow.NoRows,
		Note: importflow.Note,
		Status: importflow.Status,
	}

	importFlowInfo, err := server.store.CreateImportFlow(ctx,arg)
	if err != nil {
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "invalid JSON request",
			StatusCode: http.StatusTeapot,
		}
		util.WriteJSONResponse(w, http.StatusTeapot, jsonResponse)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(importFlowInfo)

}
func (server *Server) getAllImportFlow(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        errorResponse(w, http.StatusMethodNotAllowed, "Only GET requests are allowed")
        return
    }
    ctx := r.Context()

    homeBanners, err := server.store.SelectAllImportFlowByID(ctx)
    if err != nil {
        jsonResponse := JsonResponse{
            Status:     false,
            Message:    "Failed to fetch brand languages",
            StatusCode: http.StatusInternalServerError,
        }
        util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
        return
    }

    w.Header().Set("Content-Type", "application/json")

    if err := json.NewEncoder(w).Encode(homeBanners); err != nil {
        jsonResponse := JsonResponse{
            Status:     false,
            Message:    "Failed to encode response",
            StatusCode: http.StatusInternalServerError,
        }
        util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
        return
    }
}
