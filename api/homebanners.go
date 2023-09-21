package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	//"github.com/jackc/pgx/v5/pgtype"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/vod/db/sqlc"
	util "github.com/vod/utils"
)

type HomeBanner struct {
	ID          int32       `json:"id"`
	Name        string      `json:"name"`
	UrlType     string      `json:"url_type"`
	Description string      `json:"description"`
	Param       pgtype.Text `json:"param"`
	Image       string      `json:"image"`
	Isdefault   pgtype.Bool `json:"isdefault"`
	Sequence    pgtype.Int4 `json:"sequence"`
	StartDate   pgtype.Date `json:"start_date"`
	EndDate     pgtype.Date `json:"end_date"`
	Isdeleted   pgtype.Bool `json:"isdeleted"`
	Isblocked   pgtype.Bool `json:"isblocked"`
}

func (server *Server) createHomeBanner(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	homebanner := HomeBanner{}
	err := json.NewDecoder(r.Body).Decode(&homebanner)

	if err != nil {

		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "invalid JSON request",
			StatusCode: http.StatusTeapot,
		}
		util.WriteJSONResponse(w, http.StatusTeapot, jsonResponse)
		return
	}

	arg := db.CreateHomeBannerParams{
		Name: homebanner.Name,
		UrlType: homebanner.UrlType,
		Description: homebanner.Description,
		Param: homebanner.Param,
		Image: homebanner.Image,
		Isdefault: homebanner.Isdefault,
		Sequence: homebanner.Sequence,
		StartDate: homebanner.StartDate,
		EndDate: homebanner.EndDate,
	}

	homeBannerInfo, err := server.store.CreateHomeBanner(ctx,arg)
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
	json.NewEncoder(w).Encode(homeBannerInfo)

}
func (server *Server) getHomeBannerById(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        errorResponse(w, http.StatusMethodNotAllowed, "Only GET requests are allowed")
        return
    }
    ctx := r.Context()

    vars := mux.Vars(r)
    idParam, ok := vars["id"]
    if !ok {
        errorResponse(w, http.StatusBadRequest, "Missing 'id' URL parameter")
        return
    }

    id, err := strconv.Atoi(idParam)
    if err != nil {
        errorResponse(w, http.StatusBadRequest, "Invalid 'id' URL parameter")
        return
    }

    homeBanner, err := server.store.SelectHomeBanners(ctx, int32(id))
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

    if err := json.NewEncoder(w).Encode(homeBanner); err != nil {
        jsonResponse := JsonResponse{
            Status:     false,
            Message:    "Failed to encode response",
            StatusCode: http.StatusInternalServerError,
        }
        util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
        return
    }
}


func (server *Server) getAllHomeBanner(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        errorResponse(w, http.StatusMethodNotAllowed, "Only GET requests are allowed")
        return
    }
    ctx := r.Context()

    homeBanners, err := server.store.SelectAllHomeBanners(ctx)
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
