package random

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/cruzj6/gosomewhere/routes"
)

type randomRouter struct {}

type LatLon struct {
	Lat, Lon float32
}

type RandomPointRequest struct {
	MinLatLon, MaxLatLon LatLon
}

const BASE_ROUTE = "/random/"

func RegisterRoute() randomRouter{
	rr := randomRouter{}

	baseRoute := routes.MakeRoute(BASE_ROUTE, "GET", rr.randomPointHandler)
	distanceRoute := routes.MakeRoute("/distance/", "GET", rr.randomDistanceHandler)

	baseRoute.AddSubRoute(distanceRoute)
	routes.RegisterRoute(baseRoute)

	return rr
}

func (rr *randomRouter) randomDistanceHandler(w http.ResponseWriter, r *http.Request) {
	distance := rr.randomDistance()
	json.NewEncoder(w).Encode(distance)
}

func (rr *randomRouter) randomPointHandler(w http.ResponseWriter, r *http.Request) {
	var latlon LatLon

	if(r.Body == nil) {
		latlon = rr.randomPoint()
	} else {
		var randomPointRequest RandomPointRequest
		err := json.NewDecoder(r.Body).Decode(&randomPointRequest)

		if(err != nil) {
			fmt.Fprintf(w, err.Error())
		}
	}

	json.NewEncoder(w).Encode(latlon)
}
