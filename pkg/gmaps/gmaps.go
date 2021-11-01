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

	FindPlaceRequest struct {
		Location string `json:"location"`
	}
)

type (
	ResponseGeocodeReverse struct {
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

type (
	ResponseFindPlace struct {
		Candidates []Candidates `json:"candidates"`
		Status     string       `json:"status"`
	}

	LocationFindPlace struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	}

	NortheastFindPlace struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	}

	SouthwestFindPlace struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	}

	ViewportFindPlace struct {
		Northeast NortheastFindPlace `json:"northeast"`
		Southwest SouthwestFindPlace `json:"southwest"`
	}

	GeometryFindPlace struct {
		Location LocationFindPlace `json:"location"`
		Viewport ViewportFindPlace `json:"viewport"`
	}

	OpeningHours struct {
		OpenNow bool `json:"open_now"`
	}

	Photos struct {
		Height           int      `json:"height"`
		HTMLAttributions []string `json:"html_attributions"`
		PhotoReference   string   `json:"photo_reference"`
		Width            int      `json:"width"`
	}

	PlusCodeFindPlace struct {
		CompoundCode string `json:"compound_code"`
		GlobalCode   string `json:"global_code"`
	}

	Candidates struct {
		BusinessStatus      string            `json:"business_status"`
		FormattedAddress    string            `json:"formatted_address"`
		Geometry            GeometryFindPlace `json:"geometry"`
		Icon                string            `json:"icon"`
		IconBackgroundColor string            `json:"icon_background_color"`
		IconMaskBaseURI     string            `json:"icon_mask_base_uri"`
		Name                string            `json:"name"`
		OpeningHours        OpeningHours      `json:"opening_hours"`
		Photos              []Photos          `json:"photos"`
		PlaceID             string            `json:"place_id"`
		PlusCode            PlusCodeFindPlace `json:"plus_code"`
		Rating              float64           `json:"rating"`
		Reference           string            `json:"reference"`
		Types               []string          `json:"types"`
		UserRatingsTotal    int               `json:"user_ratings_total"`
	}
)

func GeocodeReverse(req GeocodeReverseRequest) (*ResponseGeocodeReverse, error) {
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

	var respBody ResponseGeocodeReverse
	if err := json.Unmarshal(body, &respBody); err != nil {
		log.Printf("[GeocodeReverse] Error unmarshaling data, cause: %v", err)
		return nil, err
	}

	return &respBody, nil
}

func GeocodeReverseWithFilter(req GeocodeReverseRequest) (*ResponseGeocodeReverse, error) {
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

	var respBody ResponseGeocodeReverse
	if err := json.Unmarshal(body, &respBody); err != nil {
		log.Printf("[GeocodeReverse] Error unmarshaling data, cause: %v", err)
		return nil, err
	}

	return &respBody, nil
}

// url = https://maps.googleapis.com/maps/api/place/findplacefromtext/json?input=Rumah%20Sakit%20Dirgahayu&inputtype=textquery&fields=place_id,formatted_address,name,geometry&key=AIzaSyDQYavEeONw69aOFzYOe2l7XZlPnNnY7vE

func FindPlace(req string) (*ResponseFindPlace, error) {

	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/findplacefromtext/json?input=%v&inputtype=textquery&fields=business_status,place_id,icon,icon_background_color,icon_mask_base_uri,opening_hours,reference,formatted_address,user_ratings_total,types,plus_code,price_level,rating,name,geometry,photos&key=%v", req, config.GMAPS_API)
	method := "GET"

	client := &http.Client{}
	getReq, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	resp, err := client.Do(getReq)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var respBody ResponseFindPlace
	if err := json.Unmarshal(body, &respBody); err != nil {
		log.Printf("[GeocodeReverse] Error unmarshaling data, cause: %v", err)
		return nil, err
	}

	return &respBody, nil
}
