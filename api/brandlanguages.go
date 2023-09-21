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

type BrandsLanguage struct {
	ID         int32       `json:"id"`
	BrandID    pgtype.Int4 `json:"brand_id"`
	LanguageID pgtype.Int4 `json:"language_id"`
	Name       string      `json:"name"`
}

func (server *Server) createBrandLanguage(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	brandlanguage := BrandsLanguage{}
	err := json.NewDecoder(r.Body).Decode(&brandlanguage)

	if err != nil {
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "invalid JSON request",
			StatusCode: http.StatusTeapot,
		}
		util.WriteJSONResponse(w, http.StatusTeapot, jsonResponse)
		return
	}

	arg := db.CreateBrandLanguageParams{
		Name: brandlanguage.Name,
		BrandID: brandlanguage.BrandID,
		LanguageID: brandlanguage.LanguageID,
	}

	brandLanguageInfo, err := server.store.CreateBrandLanguage(ctx,arg)
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
	json.NewEncoder(w).Encode(brandLanguageInfo)
}

func (server *Server) getBrandLanguagesById(w http.ResponseWriter, r *http.Request) {
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

    brandLanguages, err := server.store.SelectBrandLanguages(ctx, int32(id))
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

    if err := json.NewEncoder(w).Encode(brandLanguages); err != nil {
        jsonResponse := JsonResponse{
            Status:     false,
            Message:    "Failed to encode response",
            StatusCode: http.StatusInternalServerError,
        }
        util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
        return
    }
}


func (server *Server) getAllBrandLanguages(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        errorResponse(w, http.StatusMethodNotAllowed, "Only GET requests are allowed")
        return
    }
    ctx := r.Context()

    brandLanguages, err := server.store.SelectALLBrandLanguages(ctx)
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

    if err := json.NewEncoder(w).Encode(brandLanguages); err != nil {
        jsonResponse := JsonResponse{
            Status:     false,
            Message:    "Failed to encode response",
            StatusCode: http.StatusInternalServerError,
        }
        util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
        return
    }
}






