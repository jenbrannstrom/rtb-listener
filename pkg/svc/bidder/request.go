package bidder

import (
	"bytes"
	"encoding/json"
	"google-rtb/config"
	"google-rtb/model"
	"google-rtb/pkg/logger"
	"io/ioutil"
	"net/http"
)

// SendBidRequest sends bid request from google to bid url
func SendBidRequest(requestBody model.RequestBody) {
	jsonContent, err := json.Marshal(requestBody)
	if err != nil {
		params := &logger.LogParams{}
		params.Add("reason:", err)
		params.Add("requestBody:", requestBody)
		logger.ErrorP("unable to parse requestBody:", params)
		return
	}

	resp, err := http.Post(config.Cfg.BidURL, "application/json", bytes.NewBuffer(jsonContent))
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

	params := &logger.LogParams{}
	params.Add("response:", string(body))
	logger.InfoP("getting response", params)

	return
}
