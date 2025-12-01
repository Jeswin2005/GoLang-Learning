package http

import (
	"linux_agent_framework/src/errors"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestSendRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()
	u, _ := url.Parse(server.URL)
	host, port, _ := net.SplitHostPort(u.Host)

	tests := []struct {
		name          string
		conn          HttpConnection
		expectedError bool
	}{
		{
			"SuccessfulSendRequest",
			HttpConnection{
				server.Client(),
				server.Client().Transport.(*http.Transport),
				nil,
				nil,
				&HttpConfiguration{
					"",
					"http",
					http.MethodGet,
					host,
					port,
					nil,
					nil,
					nil,
				},
			},
			false,
		},

		{
			"InvalidUrl",
			HttpConnection{
				&http.Client{},
				&http.Transport{},
				nil,
				nil,
				&HttpConfiguration{
					"",
					"http",
					http.MethodGet,
					"127.0.0.1",
					"1",
					nil,
					nil,
					nil,
				},
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := tt.conn.SendRequest()
			if tt.expectedError {
				if err == errors.ERROR_NO_ERROR {
					t.Fatalf("Expected error, got no error")
				}
			} else {
				if err != errors.ERROR_NO_ERROR {
					t.Fatalf("expected no error, got %v", err)
				}

			}
		})
	}
}
