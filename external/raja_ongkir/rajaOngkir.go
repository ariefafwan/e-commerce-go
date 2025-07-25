package raja_ongkir

import (
	"e-commerce-go/pkg"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type ProvinceResponse struct {
	Meta struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
		Status  string `json:"status"`
	} `json:"meta"`
	Data []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"data"`
}

type CityResponse struct {
	Meta struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
		Status  string `json:"status"`
	} `json:"meta"`
	Data []struct{
		ID   int `json:"id"`
		Name string `json:"name"`
	} `json:"data"`
}

type DistrictResponse struct {
	Meta struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
		Status  string `json:"status"`
	} `json:"meta"`
	Data []struct{
		ID   int `json:"id"`
		Name string `json:"name"`
	} `json:"data"`
}

func GetProvince() (ProvinceResponse, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	url := "https://rajaongkir.komerce.id/api/v1/destination/province"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ProvinceResponse{}, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("key", pkg.GetEnv("RAJA_ONGKIR_API_KEY", "1234567"))

	res, err := client.Do(req)
	if err != nil {
		return ProvinceResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ProvinceResponse{}, err
	}

	var parsed ProvinceResponse
	if err := json.Unmarshal(body, &parsed); err != nil {
		return ProvinceResponse{}, err
	}

	if parsed.Meta.Code != 200 {
		return ProvinceResponse{}, errors.New(parsed.Meta.Message)
	}

	return parsed, nil
}

func GetCity(provinceID string) (CityResponse, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	url := fmt.Sprintf("https://rajaongkir.komerce.id/api/v1/destination/city/%s", provinceID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return CityResponse{}, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("key", pkg.GetEnv("RAJA_ONGKIR_API_KEY", "1234567"))

	res, err := client.Do(req)
	if err != nil {
		return CityResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return CityResponse{}, err
	}

	var parsed CityResponse
	if err := json.Unmarshal(body, &parsed); err != nil {
		return CityResponse{}, err
	}

	if parsed.Meta.Code != 200 {
		return CityResponse{}, errors.New(parsed.Meta.Message)
	}

	return parsed, nil
}

func GetDistrict(cityID string) (DistrictResponse, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	url := fmt.Sprintf("https://rajaongkir.komerce.id/api/v1/destination/district/%s", cityID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return DistrictResponse{}, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("key", pkg.GetEnv("RAJA_ONGKIR_API_KEY", "1234567"))

	res, err := client.Do(req)
	if err != nil {
		return DistrictResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return DistrictResponse{}, err
	}

	var parsed DistrictResponse
	if err := json.Unmarshal(body, &parsed); err != nil {
		return DistrictResponse{}, err
	}

	if parsed.Meta.Code != 200 {
		return DistrictResponse{}, errors.New(parsed.Meta.Message)
	}

	return parsed, nil
}
type CalculateCostResponse struct {
	Meta struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
		Status  string `json:"status"`
	} `json:"meta"`
	Data []ShippingOption `json:"data"`
}

type ShippingOption struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	Service     string `json:"service"`
	Description string `json:"description"`
	Cost        int    `json:"cost"`
	Etd         string `json:"etd"`
}

func CalculateShippingCost(originID, destinationID string, weight int) (*CalculateCostResponse, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	url := "https://rajaongkir.komerce.id/api/v1/calculate/district/domestic-cost"

	courierList := "jne:sicepat:jnt:ninja:tiki:lion:pos"

	payload := fmt.Sprintf(
		"origin=%s&destination=%s&weight=%d&courier=%s&price=lowest",
		originID, destinationID, weight, courierList,
	)

	req, err := http.NewRequest("POST", url, strings.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("key", pkg.GetEnv("RAJA_ONGKIR_API_KEY", "1234567"))

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	
	var parsed CalculateCostResponse
	if err := json.Unmarshal(body, &parsed); err != nil {
		return nil, err
	}

	if parsed.Meta.Code != 200 {
		return nil, errors.New(parsed.Meta.Message)
	}

	return &parsed, nil
}