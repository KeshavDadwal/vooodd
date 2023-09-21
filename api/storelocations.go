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

type StoreLocation struct {
	ID               int32       `json:"id"`
	StoreID          int32       `json:"store_id"`
	MallID           int32       `json:"mall_id"`
	Address          string      `json:"address"`
	City             string      `json:"city"`
	State            string      `json:"state"`
	CountryID        int32       `json:"country_id"`
	Zip              string      `json:"zip"`
	Latitude         pgtype.Text `json:"latitude"`
	Longitude        pgtype.Text `json:"longitude"`
	PhoneNo          pgtype.Text `json:"phone_no"`
	Timezone         pgtype.Text `json:"timezone"`
	Isdeleted        pgtype.Bool `json:"isdeleted"`
	Isblocked        pgtype.Bool `json:"isblocked"`
	MondayStart      pgtype.Time `json:"monday_start"`
	MondayEnd        pgtype.Time `json:"monday_end"`
	TuesdayStart     pgtype.Time `json:"tuesday_start"`
	TuesdayEnd       pgtype.Time `json:"tuesday_end"`
	WednesdayStart   pgtype.Time `json:"wednesday_start"`
	WednesdayEnd     pgtype.Time `json:"wednesday_end"`
	ThursdayStart    pgtype.Time `json:"thursday_start"`
	ThursdayEnd      pgtype.Time `json:"thursday_end"`
	FridayStart      pgtype.Time `json:"friday_start"`
	FridayEnd        pgtype.Time `json:"friday_end"`
	SaturdayStart    pgtype.Time `json:"saturday_start"`
	SaturdayEnd      pgtype.Time `json:"saturday_end"`
	SundayStart      pgtype.Time `json:"sunday_start"`
	SundayEnd        pgtype.Time `json:"sunday_end"`
	IpAddress        pgtype.Text `json:"ip_address"`
	Level            pgtype.Int4 `json:"level"`
	Floor            pgtype.Text `json:"floor"`
	Side             pgtype.Text `json:"side"`
	UserID           pgtype.Int4 `json:"user_id"`
	Map              pgtype.Text `json:"map"`
	IsHoliday        pgtype.Bool `json:"is_holiday"`
	HolidayStartDate pgtype.Date `json:"holiday_start_date"`
	HolidayEndDate   pgtype.Date `json:"holiday_end_date"`
	HolidayStartTime pgtype.Time `json:"holiday_start_time"`
	HolidayEndTime   pgtype.Time `json:"holiday_end_time"`
	CurrencyID       int32       `json:"currency_id"`
	CreatedAt        time.Time   `json:"created_at"`
}

func (server *Server) createStoreLocation(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	storelocation := StoreLocation{}
	err := json.NewDecoder(r.Body).Decode(&storelocation)

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

	arg := db.CreateStoreLocationParams{
		StoreID: storelocation.StoreID,
		MallID: storelocation.MallID,
		Address: storelocation.Address,
		City: storelocation.City,
		State: storelocation.State,
		CountryID: storelocation.CountryID,
		Zip: storelocation.Zip,
		Latitude: storelocation.Latitude,
		Longitude: storelocation.Longitude,
		PhoneNo: storelocation.PhoneNo,
		Timezone: storelocation.Timezone,
		Isdeleted: storelocation.Isdeleted,
		Isblocked: storelocation.Isblocked,
		MondayStart: storelocation.MondayStart,
		MondayEnd: storelocation.MondayEnd,
		TuesdayStart: storelocation.TuesdayStart,
		TuesdayEnd: storelocation.TuesdayEnd,
		WednesdayStart: storelocation.WednesdayStart,
		WednesdayEnd: storelocation.WednesdayEnd,
		ThursdayStart: storelocation.ThursdayStart,
		ThursdayEnd: storelocation.ThursdayEnd,
		FridayStart: storelocation.FridayStart,
		FridayEnd: storelocation.FridayEnd,
		SaturdayStart: storelocation.SaturdayStart,
		SaturdayEnd: storelocation.SaturdayEnd,
		SundayStart: storelocation.SundayStart,
		SundayEnd: storelocation.SundayEnd,
		IpAddress: storelocation.IpAddress,
		Level: storelocation.Level,
		Floor: storelocation.Floor,
		Side: storelocation.Side,
		UserID: storelocation.UserID,
		Map: storelocation.Map,
		IsHoliday: storelocation.IsHoliday,
		HolidayStartDate: storelocation.HolidayStartDate,
		HolidayEndDate: storelocation.HolidayEndDate,
		HolidayStartTime: storelocation.HolidayStartTime,
		HolidayEndTime: storelocation.HolidayEndTime,
		CurrencyID: storelocation.CurrencyID,
	}

	storeLocationInfo, err := server.store.CreateStoreLocation(ctx,arg)
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
	json.NewEncoder(w).Encode(storeLocationInfo)

}
