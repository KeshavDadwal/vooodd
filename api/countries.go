package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	//"github.com/jackc/pgx/v5/pgtype"
	"github.com/gorilla/mux"
	db "github.com/vod/db/sqlc"
	util "github.com/vod/utils"
)

type Country struct {
	ID          int32  `json:"id"`
	Iso2        string `json:"iso2"`
	ShortName   string `json:"short_name"`
	LongName    string `json:"long_name"`
	Iso3        string `json:"iso3"`
	Numcode     string `json:"numcode"`
	UnMember    string `json:"un_member"`
	CallingCode string `json:"calling_code"`
	Cctld       string `json:"cctld"`
}


func (server *Server) createCountry(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	country := Country{}
	err := json.NewDecoder(r.Body).Decode(&country)

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

	arg := db.CreateCountryParams{
		Iso2:   country.Iso2,
		ShortName: country.ShortName,
		LongName: country.LongName,
		Iso3: country.Iso3,
		Numcode: country.Numcode,
		UnMember: country.UnMember,
		CallingCode: country.CallingCode,
		Cctld: country.Cctld,	
	}

	countryInfo, err := server.store.CreateCountry(ctx,arg)
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
	json.NewEncoder(w).Encode(countryInfo)

}

func (server *Server) getCountryById(w http.ResponseWriter, r *http.Request) {
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

    country, err := server.store.GetCountry(ctx, int32(id))
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

    if err := json.NewEncoder(w).Encode(country); err != nil {
        jsonResponse := JsonResponse{
            Status:     false,
            Message:    "Failed to encode response",
            StatusCode: http.StatusInternalServerError,
        }
        util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
        return
    }
}


func (server *Server) getAllCountry(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        errorResponse(w, http.StatusMethodNotAllowed, "Only GET requests are allowed")
        return
    }
    ctx := r.Context()

    countries, err := server.store.GetAllCountry(ctx)
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

    if err := json.NewEncoder(w).Encode(countries); err != nil {
        jsonResponse := JsonResponse{
            Status:     false,
            Message:    "Failed to encode response",
            StatusCode: http.StatusInternalServerError,
        }
        util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
        return
    }
}




