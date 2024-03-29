package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/sumer312/Health-App-Backend/internal/database"
)

func (apiCfg *Api) InputHandler(w http.ResponseWriter, r *http.Request) {
	godotenv.Load()
	base_url := os.Getenv("BASE_URL")
	r.ParseForm()
	DesiredWeightIsEmpty := false
	TimeFrameIsEmpty := false
	cookieVal, err := r.Cookie("user-id")
	if err != nil {
		fmt.Println(err)
		w.Header().Add("HX-Trigger", `{ "errorToast" : "Could not parse data" }`)
		w.WriteHeader(500)
		return
	}
	userId, err := uuid.Parse(cookieVal.Value)
	if err != nil {
		fmt.Println(err)
		w.Header().Add("HX-Trigger", `{ "errorToast" : "Not logged in" }`)
		w.WriteHeader(500)
		return
	}
	height, err := strconv.ParseFloat(r.FormValue("height"), 10)
	if err != nil {
		fmt.Println(err)
		w.Header().Add("HX-Trigger", `{ "errorToast" : "Could not parse data" }`)
		w.WriteHeader(500)
		return
	}
	weight, err := strconv.ParseFloat(r.FormValue("weight"), 10)
	if err != nil {
		fmt.Println(err)
		w.Header().Add("HX-Trigger", `{ "errorToast" : "Could not parse data" }`)
		w.WriteHeader(500)
		return
	}
	desiredWeight, err := strconv.ParseFloat(r.FormValue("desired_weight"), 10)
	if err != nil && r.Form.Has("desired_weight") {
		fmt.Println(err)
		w.Header().Add("HX-Trigger", `{ "errorToast" : "Could not parse data" }`)
		w.WriteHeader(500)
		return
	} else if !r.Form.Has("desired_weight") {
		DesiredWeightIsEmpty = true
	}
	timeFrame, err := strconv.ParseFloat(r.FormValue("time_frame"), 10)
	if err != nil && r.Form.Has("time_frame") {
		fmt.Println(err)
		w.Header().Add("HX-Trigger", `{ "errorToast" : "Could not parse data" }`)
		w.WriteHeader(500)
		return
	} else if !r.Form.Has("time_frame") {
		TimeFrameIsEmpty = true
	}
	bmi := (float64(weight) * 10000) / (float64(height) * float64(height))
	currKcal := weight * 2.204 * 15
	program := r.FormValue("program")
	var sex string
	switch r.FormValue("sex") {
	case "Male":
		sex = sex_male
	case "Female":
		sex = sex_female
	default:
		sex = sex_none
	}
	if DesiredWeightIsEmpty == false && TimeFrameIsEmpty == false {
		if desiredWeight > weight {
			w.Header().Add("HX-Trigger", `{ "errorToast" : "Desired weight cannot be greater than weight" }`)
			w.WriteHeader(400)
			return
		}
		TempChan := make(chan sql.NullFloat64)
		go func(w int, dw int, tf int) {
			TempChan <- DeficitCalc(w, dw, tf)
		}(int(weight), int(desiredWeight), int(timeFrame))
		deficit := <-TempChan
		apiCfg.DB.CreateUserInput(r.Context(), database.CreateUserInputParams{
			ID:            uuid.New(),
			CreatedAt:     time.Now().UTC(),
			UpdatedAt:     time.Now().UTC(),
			UserID:        userId,
			Height:        int32(height),
			Weight:        int32(weight),
			DesiredWeight: sql.NullInt32{Int32: int32(desiredWeight), Valid: true},
			TimeFrame:     sql.NullInt32{Int32: int32(timeFrame), Valid: true},
			CurrKcal:      currKcal,
			Bmi:           bmi,
			Program:       program,
			Deficit:       deficit,
			Sex:           sex,
		})
	} else if DesiredWeightIsEmpty && TimeFrameIsEmpty == false {
		apiCfg.DB.CreateUserInput(r.Context(), database.CreateUserInputParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			UserID:    userId,
			Height:    int32(height),
			Weight:    int32(weight),
			TimeFrame: sql.NullInt32{Int32: int32(timeFrame), Valid: true},
			CurrKcal:  currKcal,
			Bmi:       bmi,
			Program:   program,
			Sex:       sex,
		})
	} else if DesiredWeightIsEmpty && TimeFrameIsEmpty {
		apiCfg.DB.CreateUserInput(r.Context(), database.CreateUserInputParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			UserID:    userId,
			Height:    int32(height),
			Weight:    int32(weight),
			CurrKcal:  currKcal,
			Bmi:       bmi,
			Program:   program,
			Sex:       sex,
		})
	}
	w.Header().Add("Hx-Redirect", base_url)
	w.WriteHeader(200)
}
