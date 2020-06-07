package api

import (
	"encoding/json"
	"fmt"
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

	err := c.BindJSON(&requestBody)

	if err != nil {
		params := &logger.LogParams{}
		params.Add("reason:", err)
		params.Add("requestBody:", requestBody)
		logger.ErrorP("unable to parse requestBody:", params)

		return
	}

	go streamer.ProcessRequestBody(requestBody)

	go bidder.SendBidRequest(requestBody)

	c.JSON(http.StatusOK, requestBody)
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
