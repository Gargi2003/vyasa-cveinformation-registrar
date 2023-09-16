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

func TestGetApidoc(t *testing.T) {
	router := gin.Default()
	router.GET("/api/v1/api-docs", handlers.GetApidoc)
	tests := []struct {
		name             string
		expectedCode     int
		expectedResponse utilities.APIDoc
	}{
		{
			name:         "Get App Info",
			expectedCode: http.StatusOK,
			expectedResponse: utilities.APIDoc{
				Link: "https://stoplight.dell.com/docs/cveinformation-1/zhn1e72hisfsn-cve-info",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, "/api/v1/api-docs", nil)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)
			assert.Equal(t, tt.expectedCode, resp.Code)
			var actualResponse utilities.APIDoc
			if err := json.Unmarshal(resp.Body.Bytes(), &actualResponse); err != nil {
				return
			}
			assert.Equal(t, tt.expectedResponse, actualResponse)
		})
	}
}
