package main

import (
	"gmaps-location-test/model/response"
	"gmaps-location-test/pkg/gmaps"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Location(req gmaps.GeocodeReverseRequest) *response.GeneralResponse {
	resp, err := gmaps.GeocodeReverse(req)
	if err != nil {
		log.Printf("[Location] Error getting location data, cause %v", err)
		return response.Error(500, "Internal Server Error.")
	}

	return response.Success(200, resp)
}

func LocationWithFilter(req gmaps.GeocodeReverseRequest) *response.GeneralResponse {
	resp, err := gmaps.GeocodeReverseWithFilter(req)
	if err != nil {
		log.Printf("[Location] Error getting location data, cause %v", err)
		return response.Error(500, "Internal Server Error.")
	}

	return response.Success(200, resp)
}

func LocationUser(c echo.Context) error {
	body := new(gmaps.GeocodeReverseRequest)
	c.Bind(body)

	response := Location(*body)
	return c.JSON(http.StatusOK, response)
}

func LocationFiltered(c echo.Context) error {
	body := new(gmaps.GeocodeReverseRequest)
	c.Bind(body)

	response := LocationWithFilter(*body)
	return c.JSON(http.StatusOK, response)
}

func main() {
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)
	e := echo.New()

	e.POST("/location", LocationUser)
	e.POST("/location/filter", LocationFiltered)

	e.Logger.Fatal(e.Start(":8080"))
}
