package orbi

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func fakeOrbiServer(user, pwd string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedUser, expectedPwd, ok := r.BasicAuth()

		if !ok || expectedUser != user || expectedPwd != pwd {
			w.WriteHeader(500)
			w.Write([]byte("Wrong username/password"))
			return
		}

		if r.URL.Path == "/RST_statistic.htm" {
			file, err := os.ReadFile("./testdata/RST_statistic.htm")
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte("couldn't find fixture file"))
			}
			w.WriteHeader(200)
			w.Write(file)
			return
		}

		w.WriteHeader(404)
	})
}

func TestHTTP(t *testing.T) {
	username := "admin"
	password := "correct-horse-battery-staple"

	testServer := httptest.NewServer(fakeOrbiServer(username, password))

	client, err := NewClient(
		testServer.URL, username, password)

	assert.NoError(t, err)

	metrics, err := client.GetMetrics()
	assert.NoError(t, err)

	assert.NotEmpty(t, metrics)
	assert.Len(t, metrics, len(allowList))
	assert.Equal(t, "3069182", metrics["sys_uptime"].Value)
}