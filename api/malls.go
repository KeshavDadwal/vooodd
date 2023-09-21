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


type Mall struct {
	ID               int32       `json:"id"`
	ParentID         pgtype.Int4 `json:"parent_id"`
	Name             string      `json:"name"`
	MallNo           pgtype.Text `json:"mall_no"`
	Address          string      `json:"address"`
	City             string      `json:"city"`
	State            string      `json:"state"`
	CountryID        int32       `json:"country_id"`
	Zip              string      `json:"zip"`
	PhoneNo          pgtype.Text `json:"phone_no"`
	Latitude         pgtype.Text `json:"latitude"`
	Longitude        pgtype.Text `json:"longitude"`
	BannerPath       pgtype.Text `json:"banner_path"`
	Isdeleted        pgtype.Bool `json:"isdeleted"`
	Isblocked        pgtype.Bool `json:"isblocked"`
	Timezone         pgtype.Text `json:"timezone"`
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
	UserID           pgtype.Int4 `json:"user_id"`
	Map              pgtype.Text `json:"map"`
	IsHoliday        pgtype.Bool `json:"is_holiday"`
	HolidayStartDate pgtype.Date `json:"holiday_start_date"`
	HolidayEndDate   pgtype.Date `json:"holiday_end_date"`
	HolidayStartTime pgtype.Time `json:"holiday_start_time"`
	HolidayEndTime   pgtype.Time `json:"holiday_end_time"`
	ZoneID           pgtype.Int4 `json:"zone_id"`
	AdLink           pgtype.Text `json:"ad_link"`
	CreatedAt        time.Time   `json:"created_at"`
}
func (server *Server) createMall (w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	mall := Mall{}
	err := json.NewDecoder(r.Body).Decode(&mall)

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

	arg := db.CreateMallParams{
		ParentID: mall.ParentID,
		Name: mall.Name,
		MallNo: mall.MallNo,
		Address: mall.Address,
		City: mall.City,
		State: mall.State,
		CountryID: mall.CountryID,
		Zip: mall.Zip,
		PhoneNo: mall.PhoneNo,
		Latitude: mall.Latitude,
		Longitude: mall.Longitude,
		BannerPath: mall.BannerPath,
		Isdeleted: mall.Isdeleted,
		Isblocked: mall.Isblocked,
		Timezone: mall.Timezone,
		MondayStart: mall.MondayStart,
		MondayEnd: mall.MondayEnd,
		TuesdayStart: mall.TuesdayStart,
		TuesdayEnd: mall.TuesdayEnd,
		WednesdayStart: mall.WednesdayStart,
		WednesdayEnd: mall.WednesdayEnd,
		ThursdayStart: mall.ThursdayStart,
		ThursdayEnd: mall.ThursdayEnd,
		FridayStart: mall.FridayStart,
		FridayEnd: mall.FridayEnd,
		SaturdayStart: mall.SaturdayStart,
		SaturdayEnd: mall.SaturdayEnd,
		SundayStart: mall.SaturdayStart,
		SundayEnd: mall.SaturdayEnd,
		UserID: mall.UserID,
		Map: mall.Map,
		IsHoliday: mall.IsHoliday,
		HolidayStartDate: mall.HolidayStartDate,
		HolidayEndDate: mall.HolidayEndDate,
		HolidayStartTime: mall.HolidayStartTime,
		HolidayEndTime: mall.HolidayEndTime,
		ZoneID: mall.ZoneID,
		AdLink: mall.AdLink,
	}

	mallInfo, err := server.store.CreateMall(ctx,arg)
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
	json.NewEncoder(w).Encode(mallInfo)

}
