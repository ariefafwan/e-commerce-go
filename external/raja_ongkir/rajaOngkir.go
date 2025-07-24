package raja_ongkir

import (
	"e-commerce-go/pkg"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
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


