package REST

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"filip/WeatherStationREST/REST/Routes/Staff"
	"filip/WeatherStationREST/REST/Routes/Temperature"

	"github.com/ant0ine/go-json-rest/rest"
)

const (
	DefaultRestPort string = "6589"
	StatusStopped   int    = 0
	StatusRunning   int    = 1

	DevStack  int = 0
	ProdStack int = 1
)

type WeatherStationREST struct {
	api         *rest.Api
	port        string
	status      int
	middlewares map[string]rest.Middleware
}

func New(environment int) (*WeatherStationREST, error) {
	var api *rest.Api = rest.NewApi()
	var middlewares map[string]rest.Middleware

	switch environment {
	case DevStack:
		api.Use(getDefaultDevMiddlewares()...)
		break
	case ProdStack:
		api.Use(getDefaultProdMiddlewares()...)
		break
	default:
		return nil, fmt.Errorf("You picked wrong environment!")
	}

	router, err := getRouter()
	if err != nil {
		return nil, err
	}

	api.SetApp(router)

	rest := WeatherStationREST{
		api:         api,
		port:        DefaultRestPort,
		status:      StatusStopped,
		middlewares: middlewares,
	}

	return &rest, nil
}

func (w *WeatherStationREST) Start() {
	w.status = StatusRunning
	log.Println("Server started at port", w.port)

	log.Fatal(http.ListenAndServe(":"+w.port, w.api.MakeHandler()))
}

func (w *WeatherStationREST) SetPort(port uint16) {
	w.port = strconv.Itoa(int(port))
}

func getDefaultDevMiddlewares() []rest.Middleware {
	return []rest.Middleware{
		&rest.AccessLogApacheMiddleware{},
		&rest.TimerMiddleware{},
		&rest.RecorderMiddleware{},
		&rest.PoweredByMiddleware{
			XPoweredBy: "WeatherStationREST by Filip",
		},
		&rest.RecoverMiddleware{
			EnableResponseStackTrace: true,
		},
		&rest.JsonIndentMiddleware{},
		&rest.ContentTypeCheckerMiddleware{},
	}
}

func getDefaultProdMiddlewares() []rest.Middleware {
	return []rest.Middleware{
		&rest.AccessLogApacheMiddleware{},
		&rest.TimerMiddleware{},
		&rest.RecorderMiddleware{},
		&rest.PoweredByMiddleware{
			XPoweredBy: "WeatherStationREST by Filip",
		},
		&rest.RecoverMiddleware{
			EnableResponseStackTrace: true,
		},
		&rest.JsonIndentMiddleware{},
		&rest.ContentTypeCheckerMiddleware{},
	}
}

func getRouter() (rest.App, error) {
	return rest.MakeRouter(
		rest.Get("/temperature/get/#from/#to", Temperature.JsonResponse),
		rest.Put("/staff/add", Staff.PutComposition),
	)
}
