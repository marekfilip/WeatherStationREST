package Staff

import (
	"io/ioutil"
	"net/http"

	compModel "filip/WeatherStationREST/Models/Set"

	"github.com/ant0ine/go-json-rest/rest"
)

func PutComposition(w rest.ResponseWriter, req *rest.Request) {
	var composition *compModel.Set

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	composition, err = compModel.CreateFromEncryptedBytes(body)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	composition.SaveAll()

	w.WriteJson(map[string]string{
		"Status": "Ok",
	})
}
