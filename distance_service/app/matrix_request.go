package app

import (
	"encoding/json"
	"fmt"
	"os"

	pb "../proto"
	resty "gopkg.in/resty.v1"
)

// DistanceElement : ...
type DistanceElement struct {
	Distance map[string]int64 `json:"distance"`
	Duration map[string]int64 `json:"duration"`
}

// DistanceRow : ...
type DistanceRow struct {
	Elements []DistanceElement `json:"elements"`
}

// DistanceMatrixResponse : ...
type DistanceMatrixResponse struct {
	Destinations []string      `json:"destination_addresses"`
	Origins      []string      `json:"origin_addresses"`
	Rows         []DistanceRow `json:"rows"`
}

// DistanceMatrixRequest : request to google distance matrix api for get ride distance/time information
func DistanceMatrixRequest(trip *pb.Trip) (*DistanceElement, error) {
	start := trip.GetStart()
	end := trip.GetEnd()

	origin := fmt.Sprintf("%v,%v", start.GetLat(), start.GetLong())
	destination := fmt.Sprintf("%v,%v", end.GetLat(), end.GetLong())

	// ########################

	GoogleAPIKey := os.Getenv("GOOGLE_API_KEY")

	resp, err := resty.R().
		SetQueryParams(map[string]string{
			"units":        "metric",
			"mode":         "driving",
			"key":          GoogleAPIKey,
			"origins":      origin,
			"destinations": destination,
		}).
		SetHeader("Accept", "application/json").
		Get("https://maps.googleapis.com/maps/api/distancematrix/json")

	if err != nil {
		return nil, err
	}

	jsonResponse := new(DistanceMatrixResponse)
	json.Unmarshal(resp.Body(), &jsonResponse)

	element := jsonResponse.Rows[0].Elements[0]

	return &element, nil
}
