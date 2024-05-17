package orbi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Client interface {
	GetMetrics() (map[string]*Metric, error)
	Host() string
	GetDevices() ([]Device, error)
}

type Metric struct {
	Name  string
	Value string
}

type realClient struct {
	url      string
	username string
	password string
}

func (r *realClient) GetMetrics() (map[string]*Metric, error) {
	req, err := http.NewRequest("GET", r.url+"/RST_statistic.htm", nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(r.username, r.password)
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code: %s", resp.Status)
	}

	return parseMetrics(resp.Body)
}

func (r *realClient) GetDevices() ([]Device, error) {
	req, err := http.NewRequest("GET", r.url+"/DEV_device_info.htm", nil)
	if err != nil {
		return nil, err
	}
	queryParams := req.URL.Query()
	queryParams.Set("ts", fmt.Sprintf("%d", time.Now().UnixMilli()))
	req.URL.RawQuery = queryParams.Encode()
	req.SetBasicAuth(r.username, r.password)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code: %s", resp.Status)
	}

	return parseDeviceMetrics(resp.Body)
}

func parseDeviceMetrics(body io.Reader) ([]Device, error) {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}
	_, jsonArrayString, ok := strings.Cut(string(bodyBytes), "device=")
	if !ok {
		fmt.Println(string(bodyBytes))
		return nil, fmt.Errorf("unexpected response format")
	}
	var devices []Device
	return devices, json.NewDecoder(strings.NewReader(jsonArrayString)).Decode(&devices)
}

func (r *realClient) Host() string {
	return r.url
}

func NewClient(url, username, password string) (Client, error) {
	return &realClient{
		url:      url,
		username: username,
		password: password,
	}, nil
}
