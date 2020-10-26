package bidder

import (
	"bytes"
	"encoding/json"
	"google-rtb/model"
	"google-rtb/pkg/logger"
	"io/ioutil"
	"net/http"
)

// SendBidRequest sends bid request from google to bid url
func SendBidRequest(url string, requestBody model.RequestBody) *model.RequestBody {
	var res *model.RequestBody
	jsonContent, _ := json.Marshal(requestBody)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonContent))
	if err != nil {
		params := &logger.LogParams{}
		params.Add("reason:", err)
		logger.ErrorP("unable to send requestBody:", params)
		return nil
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		params := &logger.LogParams{}
		params.Add("reason:", err)
		logger.ErrorP("unable to read response:", params)
		return nil
	}

	if len(body) > 0 {
		params := &logger.LogParams{}
		params.Add("response:", string(body))
		logger.ErrorP("getting response", params)

		_ = json.Unmarshal(body, &res)
		// go streamer.ProcessRequestBody(res)
		return res
	}
	return nil
}
