package random

import "math/rand"

func (rr *randomRouter) randomPoint() LatLon{
	num := rand.Float32()
	longitude := num * 180
	latitude := num * 90

	latlon := LatLon{ Lat: latitude, Lon: longitude }
	return latlon
}
