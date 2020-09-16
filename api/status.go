package api

import (
	"encoding/json"
	"fmt"
	"google-rtb/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// StatusCheck tests health
func StatusCheck(c *gin.Context) {
	var requestBody model.RequestBody

	b, _ := json.Marshal(requestBody)
	fmt.Println("request body", time.Now().String(), string(b))
	c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
