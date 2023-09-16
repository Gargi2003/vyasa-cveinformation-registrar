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

func TestGetAppInfo(t *testing.T) {
	router := gin.Default()
	router.GET("/api/v1/info", handlers.GetAppInfo)
	tests := []struct {
		name             string
		expectedCode     int
		expectedResponse utilities.CommonAppInformation
	}{
		{
			name:         "Get App Info",
			expectedCode: http.StatusOK,
			expectedResponse: utilities.CommonAppInformation{
				Id:          "12973",
				Name:        "CVEInformation",
				Identifier:  "cve-info",
				Description: "An app to retrieve CVE information",
				Version:     "2.0",
				BuildNumber: "54321",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, "/api/v1/info", nil)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)
			assert.Equal(t, tt.expectedCode, resp.Code)
			var actualResponse utilities.CommonAppInformation
			if err := json.Unmarshal(resp.Body.Bytes(), &actualResponse); err != nil {
				return
			}
			assert.Equal(t, tt.expectedResponse, actualResponse)
		})
	}
}
