package gmaps

import (
	"encoding/json"
	"fmt"
	"gmaps-location-test/config"
	"io/ioutil"
	"log"
	"net/http"
)

type (
	GeocodeReverseRequest struct {
		Lat  string `json:"lat"`
		Long string `json:"long"`
	}

	Response struct {
		PlusCode PlusCode  `json:"plus_code"`
		Results  []Results `json:"results"`
		Status   string    `json:"status"`
	}

	PlusCode struct {
		CompoundCode string `json:"compound_code"`
		GlobalCode   string `json:"global_code"`
	}
	AddressComponents struct {
		LongName  string   `json:"long_name"`
		ShortName string   `json:"short_name"`
		Types     []string `json:"types"`
	}
	Location struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	}
	Northeast struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	}
	Southwest struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	}
	Viewport struct {
		Northeast Northeast `json:"northeast"`
		Southwest Southwest `json:"southwest"`
	}
	Bounds struct {
		Northeast Northeast `json:"northeast"`
		Southwest Southwest `json:"southwest"`
	}
	Geometry struct {
		Bounds       Bounds   `json:"bounds"`
		Location     Location `json:"location"`
		LocationType string   `json:"location_type"`
		Viewport     Viewport `json:"viewport"`
	}
	Results struct {
		AddressComponents []AddressComponents `json:"address_components"`
		FormattedAddress  string              `json:"formatted_address"`
		Geometry          Geometry            `json:"geometry,omitempty"`
		PlaceID           string              `json:"place_id"`
		PlusCode          PlusCode            `json:"plus_code,omitempty"`
		Types             []string            `json:"types"`
	}
)

func GeocodeReverse(req GeocodeReverseRequest) (*Response, error) {
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?latlng=%s,%s&key=%s", req.Lat, req.Long, config.GMAPS_API)
	method := "GET"

	client := &http.Client{}
	getReq, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Printf("[GeocodeReverse] Error create new request, cause: %v", err)
		return nil, err
	}

	resp, err := client.Do(getReq)
	if err != nil {
		log.Printf("[GeocodeReverse] Error do request, cause: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[GeocodeReverse] Error cause: %v", err)
		return nil, err
	}

	var respBody Response
	if err := json.Unmarshal(body, &respBody); err != nil {
		log.Printf("[GeocodeReverse] Error unmarshaling data, cause: %v", err)
		return nil, err
	}

	return &respBody, nil
}

func GeocodeReverseWithFilter(req GeocodeReverseRequest) (*Response, error) {
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?latlng=%s,%s&result_type=street_address&key=%s", req.Lat, req.Long, config.GMAPS_API)
	method := "GET"

	client := &http.Client{}
	getReq, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Printf("[GeocodeReverse] Error create new request, cause: %v", err)
		return nil, err
	}

	resp, err := client.Do(getReq)
	if err != nil {
		log.Printf("[GeocodeReverse] Error do request, cause: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[GeocodeReverse] Error cause: %v", err)
		return nil, err
	}

	var respBody Response
	if err := json.Unmarshal(body, &respBody); err != nil {
		log.Printf("[GeocodeReverse] Error unmarshaling data, cause: %v", err)
		return nil, err
	}

	return &respBody, nil
}
