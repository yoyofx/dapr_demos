package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	daprPort      = "3500"
	stateStoreURL = "http://localhost:" + daprPort + "/v1.0/state/redis"
)

type stateRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func main() {
	r := gin.Default()

	r.POST("/state/value", func(c *gin.Context) {
		var reqBody struct {
			Value string `json:"value"`
		}
		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		req := []stateRequest{{
			Key:   "myway",
			Value: reqBody.Value,
		},
		}

		reqBytes, err := json.Marshal(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		client := &http.Client{Timeout: 5 * time.Second}
		resp, err := client.Post(stateStoreURL, "application/json", bytes.NewBuffer(reqBytes))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode >= http.StatusBadRequest {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to save state: %s", resp.Status)})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "State saved successfully"})
	})

	if err := r.Run(":3000"); err != nil {
		fmt.Printf("error starting server: %v\n", err)
		os.Exit(1)
	}
}
