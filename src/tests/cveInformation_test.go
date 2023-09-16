package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"eos2git.cec.lab.emc.com/Gargi-Bandyopadhyay/vyasa-cveinformation-registrar/src/handlers"
	"eos2git.cec.lab.emc.com/Gargi-Bandyopadhyay/vyasa-cveinformation-registrar/src/utilities"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	filePath       = "src/examples/CVEInformation.json"
	UpdatefilePath = "src/examples/CVEInformation-Update.json"
)

func unmarshalJsonData(jsonMessage string) ([]utilities.CVEInformation, error) {
	req := []utilities.CVEInformation{}
	if err := json.Unmarshal([]byte(jsonMessage), &req); err != nil {
		return nil, err
	}
	return req, nil
}
func readFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
func TestGetCVEInfo(t *testing.T) {
	router := gin.Default()
	router.GET("/api/v1/cveinfo", handlers.GetCVEInfo)
	jsonData, err := readFile(filePath)
	if err != nil {
		return
	}
	testData, err := unmarshalJsonData(jsonData)
	if err != nil {
		return
	}
	type args struct {
		c *gin.Context
	}
	testCases := []struct {
		name     string
		query    string
		expected int
	}{
		{
			name:     "Get CVE Info-Positive",
			query:    "cve_id=CVE-2023-12345",
			expected: http.StatusOK,
		},
		{
			name:     "Get CVE Info- Not found",
			query:    "cve_id=CVE-2023-12346",
			expected: http.StatusNotFound,
		},
	}

	utilities.CveData = testData
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, "/api/v1/cveinfo?"+tc.query, nil)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)
			assert.Equal(t, tc.expected, resp.Code)
			if tc.expected == http.StatusOK {
				var response utilities.CVEInformation
				err := json.Unmarshal(resp.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, "CVE-2023-12345", response.CVEID)
			}
		})
	}
}

func TestListCVEInfo(t *testing.T) {
	router := gin.Default()
	router.GET("/api/v1/cveinfo", handlers.ListCVEInfo)
	jsonData, err := readFile(filePath)
	if err != nil {
		return
	}
	testData, err := unmarshalJsonData(jsonData)
	if err != nil {
		return
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name     string
		expected int
	}{
		{
			name:     "List CVE Info-Positive",
			expected: http.StatusOK,
		},
	}
	utilities.CveData = testData
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, "/api/v1/cveinfo", nil)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)
			assert.Equal(t, tt.expected, resp.Code)
			if tt.expected == http.StatusOK {
				var response []utilities.CVEInformation
				err := json.Unmarshal(resp.Body.Bytes(), &response)
				assert.NoError(t, err)
			}
		})
	}
}

func TestCreateCVEInfo(t *testing.T) {
	router := gin.Default()
	router.POST("/api/v1/cveinfo", handlers.CreateCVEInfo)
	jsonData, err := readFile(filePath)
	if err != nil {
		return
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name     string
		data     string
		expected int
	}{
		{
			name:     "Create CVE Info- Success",
			data:     jsonData,
			expected: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodPost, "/api/v1/cveinfo", strings.NewReader(tt.data))
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)
			assert.Equal(t, tt.expected, resp.Code)
		})
	}
}

func TestUpdateCVEInfo(t *testing.T) {
	router := gin.Default()
	router.PUT("/api/v1/cveinfo", handlers.UpdateCVEInfo)
	jsonData, err := readFile(filePath)
	if err != nil {
		return
	}
	testData, err1 := unmarshalJsonData(jsonData)
	if err1 != nil {
		return
	}

	//update testData
	updateJsonData, err := readFile(UpdatefilePath)
	if err != nil {
		return
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name     string
		expected int
		query    string
		data     string
	}{
		{
			name:     "Update CVE Info-Positive",
			query:    "cve_id=CVE-2023-12345",
			data:     updateJsonData,
			expected: http.StatusOK,
		},
		{
			name:     "Update CVE Info- Not found",
			query:    "cve_id=CVE-2023-12346",
			data:     updateJsonData,
			expected: http.StatusNotFound,
		},
	}
	utilities.CveData = testData
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodPut, "/api/v1/cveinfo?"+tt.query, strings.NewReader(tt.data))
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)
			assert.Equal(t, tt.expected, resp.Code)
			if tt.expected == http.StatusOK {
				var response utilities.CVEInformation
				err := json.Unmarshal(resp.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, "CVE-2023-12345", response.CVEID)
			}
		})
	}
}

func TestDeleteCVEInfo(t *testing.T) {
	router := gin.Default()
	router.DELETE("/api/v1/cveinfo", handlers.DeleteCVEInfo)
	jsonData, err := readFile(filePath)
	if err != nil {
		return
	}
	testData, err := unmarshalJsonData(jsonData)
	if err != nil {
		return
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name     string
		query    string
		expected int
	}{
		{
			name:     "Delete CVE Info-Positive",
			query:    "cve_id=CVE-2023-12345",
			expected: http.StatusOK,
		},
		{
			name:     "Delete CVE Info- Not found",
			query:    "cve_id=CVE-2023-12346",
			expected: http.StatusNotFound,
		},
	}
	utilities.CveData = testData

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodDelete, "/api/v1/cveinfo?"+tc.query, nil)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)
			assert.Equal(t, tc.expected, resp.Code)
			if tc.expected == http.StatusOK {
				var response utilities.CVEInformation
				err := json.Unmarshal(resp.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, "CVE-2023-12345", response.CVEID)
			}
		})
	}
}
