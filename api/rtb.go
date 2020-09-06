package api

import (
	"encoding/json"
	"fmt"
	"google-rtb/config"
	"google-rtb/model"
	"google-rtb/pkg/logger"
	"google-rtb/pkg/svc/bidder"
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

	// go streamer.ProcessRequestBody(requestBody)
	// b, _ := json.Marshal(requestBody)
	// sr := string(b)

	for _, v := range config.Cfg.BidURL {
		// if strings.Contains(sr, v.BillingID) {
		// 	bidder.SendBidRequest(c, v.URL, requestBody)
		// }
		paramss := &logger.LogParams{}
		paramss.Add("url:", v.URL)
		paramss.Add("requestBody:", requestBody)
		logger.ErrorP("get receive", paramss)
		bidder.SendBidRequest(c, v.URL, requestBody)
	}
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
