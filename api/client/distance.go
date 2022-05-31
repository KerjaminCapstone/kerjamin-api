package client

import (
	"context"
	"log"
	"net/http"

	"github.com/KerjaminCapstone/kerjamin-backend-v1/vendor/github.com/labstack/echo/v4"

	"googlemaps.github.io/maps"
)

func Distance(c echo.Context) error {

	d, err := maps.NewClient(maps.WithAPIKey("Insert-API-Key-Here"))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	r := &maps.DirectionsRequest{
		Origin:      "Sydney",
		Destination: "Perth",
	}
	route, _, err := d.Directions(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	return c.JSON(http.StatusOK, route)
}
