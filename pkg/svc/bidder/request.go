package bidder

import (
	"bytes"
	"encoding/json"
	"google-rtb/model"
	"google-rtb/pkg/logger"
	"google-rtb/pkg/svc/streamer"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SendBidRequest sends bid request from google to bid url
func SendBidRequest(c *gin.Context, url string, requestBody model.RequestBody) {
	if rand.Float32() > 0.5 {
		return
	}
	var res model.RequestBody
	jsonContent, err := json.Marshal(requestBody)
	if err != nil {
		params := &logger.LogParams{}
		params.Add("reason:", err)
		params.Add("requestBody:", requestBody)
		logger.ErrorP("unable to parse requestBody:", params)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonContent))
	if err != nil {
		params := &logger.LogParams{}
		params.Add("reason:", err)
		params.Add("requestBody:", requestBody)
		logger.ErrorP("unable to send requestBody:", params)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		params := &logger.LogParams{}
		params.Add("reason:", err)
		logger.ErrorP("unable to read response:", params)
	}

	if len(body) > 0 {
		_ = json.Unmarshal(body, &res)
		c.Header("Content-Type", "application/octet-stream")
		c.JSON(http.StatusOK, res)
		go streamer.ProcessRequestBody(res)

		params := &logger.LogParams{}
		params.Add("anurl:", url)
		params.Add("status:", 200)
		logger.InfoP("getting response", params)
	}
	return
}
