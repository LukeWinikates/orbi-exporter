package orbi

import (
	"fmt"
	"net/http"
)

type Client interface {
	GetMetrics() (map[string]Metric, error)
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

func (r realClient) GetMetrics() (map[string]Metric, error) {
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

	return parse(resp.Body)
}

func NewClient(url, username, password string) (Client, error) {
	return &realClient{
		url:      url,
		username: username,
		password: password,
	}, nil
}
