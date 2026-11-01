package main

import (
	"encoding/json"
	"fmt"

	geojson "github.com/paulmach/go.geojson"
)

func main1() {
	featureCollection := geojson.NewFeatureCollection()

	point1 := geojson.NewPointFeature([]float64{125.6, 10.1})
	point1.SetProperty("city", "Manila")

	point2 := geojson.NewPointFeature([]float64{120.98, 14.6})
	point2.SetProperty("city", "Quezon City")

	featureCollection.AddFeature(point1)
	featureCollection.AddFeature(point2)

	data, _ := json.MarshalIndent(featureCollection, "", "  ")
	fmt.Println(string(data))
}
