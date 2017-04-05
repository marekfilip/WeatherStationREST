package Brightness

import (
	"net/http"
	"time"

	"filip/WeatherStationREST/Models/Brightness/Signed"
	util "filip/WeatherStationREST/REST/Utilities"

	"github.com/ant0ine/go-json-rest/rest"
)

func Range(w rest.ResponseWriter, req *rest.Request) {
	from := req.PathParam("from")
	to := req.PathParam("to")

	fromTime, err := time.Parse("20060102", from)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	toTime, err := time.Parse("20060102", to)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if util.IsLimitExceeded(fromTime, toTime) {
		rest.Error(w, "Requested range too long", http.StatusForbidden)
		return
	}

	toTime = toTime.Add(time.Duration(24)*time.Hour - time.Nanosecond)

	requestedBrightnessSet, err := Signed.Find(fromTime, toTime)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteJson(requestedBrightnessSet)
}

func Last(w rest.ResponseWriter, req *rest.Request) {
	requestedBrightness, err := Signed.GetLast()
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteJson(requestedBrightness)
}
