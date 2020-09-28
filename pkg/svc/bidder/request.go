package bidder

import (
	"bytes"
	"encoding/json"
	"google-rtb/model"
	"google-rtb/pkg/logger"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SendBidRequest sends bid request from google to bid url
func SendBidRequest(c *gin.Context, url string, requestBody model.RequestBody) {
	var res model.RequestBody
	jsonContent, err := json.Marshal(requestBody)
	if err != nil {
		params := &logger.LogParams{}
		params.Add("reason:", err)
		logger.ErrorP("unable to parse requestBody:", params)
		c.JSON(http.StatusOK, "ignore")
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonContent))
	if err != nil {
		params := &logger.LogParams{}
		params.Add("reason:", err)
		logger.ErrorP("unable to send requestBody:", params)
		c.JSON(http.StatusOK, "ignore")
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		params := &logger.LogParams{}
		params.Add("reason:", err)
		logger.ErrorP("unable to read response:", params)
		c.JSON(http.StatusOK, "ignore")
		return
	}

	if len(body) > 0 {
		_ = json.Unmarshal(body, &res)
		c.Header("Content-Type", "application/octet-stream")
		c.JSON(http.StatusOK, res)
		// go streamer.ProcessRequestBody(res)
	} else {
		c.Header("Content-Type", "application/octet-stream")
		c.JSON(http.StatusOK, "ignore")
	}
	return
}
