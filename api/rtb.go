package api

import (
	"encoding/json"
	"fmt"
	"google-rtb/config"
	"google-rtb/model"
	"google-rtb/pkg/logger"
	"google-rtb/pkg/svc/bidder"
	"google-rtb/pkg/svc/streamer"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// RtbListener listen all request coming
func RtbListener(c *gin.Context) {
	var requestBody model.RequestBody
	var res *model.RequestBody
	c.Writer.Header().Set("Connection", "close")
	defer c.Request.Body.Close()

	err := c.BindJSON(&requestBody)
	if err != nil {
		params := &logger.LogParams{}
		params.Add("reason:", err)
		params.Add("requestBody:", requestBody)
		logger.ErrorP("unable to parse requestBody:", params)
		c.JSON(http.StatusOK, "ignore")
		return
	}

	params := &logger.LogParams{}
	params.Add("requestBody:", requestBody)
	logger.ErrorP("testing request comming:", params)

	if config.Cfg.S3Stream == true {
		go func() {
			streamer.ProcessRequestBody(requestBody)
		}()
	}

	for _, v := range config.Cfg.BidURL {
		result := bidder.SendBidRequest(v.URL, requestBody)
		if result != nil {
			res = result
		}
	}

	if res != nil {
		c.Header("Content-Type", "application/octet-stream")
		c.JSON(http.StatusOK, res)
		return
	}
	c.JSON(http.StatusOK, "successs")
}

// RtbListenCheck test send request
func RtbListenCheck(c *gin.Context) {
	var requestBody model.RequestBody
	err := c.BindJSON(&requestBody)
	if err != nil {
		params := &logger.LogParams{}
		params.Add("reason:", err)
		params.Add("requestBody:", requestBody)
		logger.ErrorP("unable to parse requestBody:", params)

		return
	}
	b, _ := json.Marshal(requestBody)
	fmt.Println("request body", time.Now().String(), string(b))
	c.JSON(http.StatusOK, requestBody)
}
