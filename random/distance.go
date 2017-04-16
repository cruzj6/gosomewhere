package random

import "math"

type Distance struct {
	Startpoint, Endpoint LatLon
	Distance float32
	Units string
}

func (rr *randomRouter) randomDistance() Distance {
	startpoint := rr.randomPoint()
	endpoint := rr.randomPoint()
	dist := LatLonDistance(startpoint, endpoint)

	distance := Distance{
		Startpoint: startpoint,
		Endpoint: endpoint,
		Distance: dist,
		Units: "m",
	}

	return distance
}

func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

// Thanks to git gist: https://gist.github.com/cdipaolo/d3f8db3848278b49db68
func LatLonDistance(latlon1, latlon2 LatLon) float32 {

	// Must be float64 for math lib
	var la1, lo1, la2, lo2, r float64
	la1 = float64(latlon1.Lat) * math.Pi / 180
	lo1 = float64(latlon1.Lon) * math.Pi / 180
	la2 = float64(latlon2.Lat) * math.Pi / 180
	lo2 = float64(latlon2.Lon) * math.Pi / 180

	r = 6378100 // Earth radius in METERS

	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return float32(2 * r * math.Asin(math.Sqrt(h)))
}
