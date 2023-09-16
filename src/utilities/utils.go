package utilities

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func CheckCVEExistence(req CVEInformation, c *gin.Context) bool {
	for _, cve := range CveData {
		if req.CVEID == cve.CVEID {
			log.Error().Msgf("CVE with cve_id %s already exists", req.CVEID)
			return true
		}
	}
	return false
}
