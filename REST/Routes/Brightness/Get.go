package Brightness

/*
import (
	"net/http"
	"time"

	bModel "filip/WeatherStationREST/Models/Brightness"

	"github.com/ant0ine/go-json-rest/rest"
)

func JsonResponse(w rest.ResponseWriter, req *rest.Request) {
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
	toTime = toTime.Add(time.Duration(24)*time.Hour - time.Nanosecond)

	requestedBrightnessSet := new(bModel.BrightnessSet)
	err = requestedBrightnessSet.Find(fromTime, toTime)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteJson(requestedBrightnessSet)
}
*/
