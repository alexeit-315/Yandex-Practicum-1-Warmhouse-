package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"io"
)

// TemperatureService handles fetching temperature data from external API
type TemperatureService struct {
	BaseURL    string
	HTTPClient *http.Client
}

// TemperatureResponse represents the response from the temperature API
type TemperatureResponse struct {
	Value       float64   `json:"value"`
	Unit        string    `json:"unit"`
	Timestamp   time.Time `json:"timestamp"`
	Location    string    `json:"location"`
	Status      string    `json:"status"`
	SensorID    string    `json:"sensor_id"`
	SensorType  string    `json:"sensor_type"`
	Description string    `json:"description"`
}

// NewTemperatureService creates a new temperature service
func NewTemperatureService(baseURL string) *TemperatureService {
	return &TemperatureService{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetTemperature fetches temperature data for a specific location
func (s *TemperatureService) GetTemperature(location string) (*TemperatureResponse, error) {
	url := fmt.Sprintf("%s/temperature?location=%s", s.BaseURL, location)

	resp, err := s.HTTPClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching temperature data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var temperatureResp TemperatureResponse
	if err := json.NewDecoder(resp.Body).Decode(&temperatureResp); err != nil {
		return nil, fmt.Errorf("error decoding temperature response: %w", err)
	}

	return &temperatureResp, nil
}

// GetTemperatureByID fetches temperature data for a specific sensor ID
func (s *TemperatureService) GetTemperatureByID(sensorID string) (*TemperatureResponse, error) {
    url := fmt.Sprintf("%s/temperature/%s", s.BaseURL, sensorID)

    resp, err := s.HTTPClient.Get(url)
    if err != nil {
        return nil, fmt.Errorf("error fetching temperature data: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
    }

    // Читаем тело ответа один раз
    bodyBytes, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("error reading response body: %w", err)
    }

    // Выводим сырой JSON для отладки
    fmt.Println("Raw JSON response:")
    fmt.Println(string(bodyBytes))

    var temperatureResp TemperatureResponse
    if err := json.Unmarshal(bodyBytes, &temperatureResp); err != nil {
        return nil, fmt.Errorf("error decoding temperature response: %w", err)
    }

    // Выводим структурированные данные для отладки
    fmt.Println("Structured data:")
    fmt.Printf("%+v\n", temperatureResp)

    return &temperatureResp, nil
}
