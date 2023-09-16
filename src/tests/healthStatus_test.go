package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"eos2git.cec.lab.emc.com/Gargi-Bandyopadhyay/vyasa-cveinformation-registrar/src/handlers"
	"eos2git.cec.lab.emc.com/Gargi-Bandyopadhyay/vyasa-cveinformation-registrar/src/utilities"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetHealthStatus(t *testing.T) {
	router := gin.Default()
	router.GET("/api/v1/health", handlers.GetHealthStatus)
	tests := []struct {
		name             string
		expectedCode     int
		expectedResponse utilities.CommonHealthCheckResponse
	}{
		{
			name:         "Get Health Info",
			expectedCode: http.StatusOK,
			expectedResponse: utilities.CommonHealthCheckResponse{
				ServiceName: "CVEInformation",
				Status:      "Up",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, "/api/v1/health", nil)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)
			assert.Equal(t, tt.expectedCode, resp.Code)
			var actualResponse utilities.CommonHealthCheckResponse
			if err := json.Unmarshal(resp.Body.Bytes(), &actualResponse); err != nil {
				return
			}
			assert.Equal(t, tt.expectedResponse, actualResponse)
		})
	}
}
