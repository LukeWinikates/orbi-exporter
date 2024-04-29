package orbi

import (
	"io"
	"net/http"
)

type Client interface {
	GetMetrics() ([]Metric, error)
}

type Metric struct {
}

type realClient struct {
	url      string
	username string
	password string
}

func (r realClient) GetMetrics() ([]Metric, error) {
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

	return parse(resp.Body)
}

func parse(body io.ReadCloser) ([]Metric, error) {
	_, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func NewClient(url, username, password string) (Client, error) {
	return &realClient{
		url:      url,
		username: username,
		password: password,
	}, nil
}
