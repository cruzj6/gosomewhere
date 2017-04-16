package random

import (
	"net/http"
	"encoding/json"

	"github.com/cruzj6/gosomewhere/routes"
)

type randomRouter struct {}

type LatLon struct {
	Lat, Lon float32
}

const BASE_ROUTE = "/random/"

func RegisterRoute() randomRouter{
	rr := randomRouter{}

	baseRoute := routes.MakeRoute(BASE_ROUTE, "GET", rr.randomHandler)
	distanceRoute := routes.MakeRoute("/distance/", "GET", rr.distanceHandler)

	baseRoute.AddSubRoute(distanceRoute)
	routes.RegisterRoute(baseRoute)

	return rr
}

func (rr *randomRouter) distanceHandler(w http.ResponseWriter, r *http.Request) {
	distance := rr.randomDistance()
	json.NewEncoder(w).Encode(distance)
}

func (rr *randomRouter) randomHandler(w http.ResponseWriter, r *http.Request) {
	latlon := rr.randomPoint()
	json.NewEncoder(w).Encode(latlon)
}
