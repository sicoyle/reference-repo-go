package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleIndex(t *testing.T) {
	tests := []struct {
		name       string
		req        *http.Request
		wantStatus int
	}{
		{
			name:       "success - happy path",
			req:        httptest.NewRequest("GET", "/api/v1/", nil),
			wantStatus: http.StatusNoContent,
		},
	}
	for _, tc := range tests {
		srv := NewServer()
		srv.Routes()
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			rr := httptest.NewRecorder()
			srv.ServeHTTP(rr, tc.req)
			assert.Equal(t, tc.wantStatus, rr.Code)
		})
	}
}

func TestHandleAddition(t *testing.T) {
	tests := []struct {
		name        string
		method      string
		payloadBody string
		want        string
		wantStatus  int
	}{
		{
			name:        "success - happy path",
			method:      http.MethodPost,
			payloadBody: `{"numbers":[1,2,3]}`,
			want:        "6",
			wantStatus:  http.StatusOK,
		},
		{
			name:        "success - no nums",
			method:      http.MethodPost,
			payloadBody: `{"numbers":[]}`,
			want:        "0",
			wantStatus:  http.StatusOK,
		},
		{
			name:        "fail - no input",
			method:      http.MethodPost,
			payloadBody: ``,
			want:        "0",
			wantStatus:  http.StatusInternalServerError,
		},
		{
			name:        "fail - invalid input",
			method:      http.MethodPost,
			payloadBody: `{"numbers":[a,b,c]}`,
			want:        "0",
			wantStatus:  http.StatusInternalServerError,
		},
	}
	for _, tc := range tests {
		srv := NewServer()
		srv.Routes()
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			rr := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/api/v1/add", strings.NewReader(tc.payloadBody))
			srv.ServeHTTP(rr, req)
			assert.Equal(t, tc.wantStatus, rr.Code)
			assert.Equal(t, tc.want, strings.TrimSpace(rr.Body.String()))
		})
	}
}
