package handlers

import (
	"fmt"
	"net/http"

	"eos2git.cec.lab.emc.com/Gargi-Bandyopadhyay/vyasa-cveinformation-registrar/src/utilities"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func GetCVEInfo(c *gin.Context) {
	id := c.Query("cve_id")
	for _, cve := range utilities.CveData {
		if id == cve.CVEID {
			log.Debug().Msgf("CVE Info:  %+v\n", cve)
			c.JSON(http.StatusOK, cve)
			return
		}
	}
	log.Error().Msgf("No CVE Info found with the cve_id: %s", id)
	c.JSON(http.StatusNotFound, "No CVE Info found with the given id")
}

func ListCVEInfo(c *gin.Context) {
	if len(utilities.CveData) > 0 {
		log.Debug().Msgf("CVE Info:  %+v\n", utilities.CveData)
		c.JSON(http.StatusOK, utilities.CveData)
		return
	}
	log.Error().Msgf("No CVE Info found")
	c.JSON(http.StatusNotFound, "No CVE Info found")
}

func CreateCVEInfo(c *gin.Context) {
	req := utilities.CVEInformation{}

	if err := c.BindJSON(&req); err != nil {
		log.Error().Msgf("Error binding request object. Error: ", err)
		c.JSON(http.StatusInternalServerError, "Error binding request object")
		return
	}
	if utilities.CheckCVEExistence(req, c) {
		cveExists := fmt.Sprintf("CVE with CVE ID %s already exists", req.CVEID)
		log.Error().Msgf(cveExists)
		c.JSON(http.StatusBadRequest, cveExists)
		return
	}
	utilities.CveData = append(utilities.CveData, req)
	c.JSON(http.StatusOK, "CVE Info creation successful!")
}

func UpdateCVEInfo(c *gin.Context) {
	id := c.Query("cve_id")
	index := -1
	for i, cve := range utilities.CveData {
		if id == cve.CVEID {
			index = i
		}
	}
	req := utilities.CVEInformation{}
	if err := c.BindJSON(&req); err != nil {
		log.Error().Msgf("Error binding request object. Error: ", err)
		c.JSON(http.StatusInternalServerError, "Error binding request object")
		return
	}
	if index == -1 {
		log.Error().Msgf("No CVE Info found with the cve_id: %s", id)
		c.JSON(http.StatusNotFound, "No CVE Info found with the given id")
		return
	}
	utilities.CveData[index] = req

	c.JSON(http.StatusOK, "CVE Info updated successfully!")
}

func DeleteCVEInfo(c *gin.Context) {
	id := c.Query("cve_id")
	for index, cve := range utilities.CveData {
		if id == cve.CVEID {
			utilities.CveData = append(utilities.CveData[:index], utilities.CveData[index+1:]...)
			c.JSON(http.StatusOK, "CVE Info deletion successful!")
			return
		}
	}
	log.Error().Msgf("No CVE Info found with the cve_id: %s", id)
	c.JSON(http.StatusNotFound, "No CVE Info found with the given id")
}
