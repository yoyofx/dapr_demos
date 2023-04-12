package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	dapr "github.com/dapr/go-sdk/client"
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
	log := log.Default()
	// 创建 Dapr 客户端
	client, err := dapr.NewClient()
	if err != nil {
		fmt.Println("无法创建 Dapr 客户端:", err)
		os.Exit(1)
	}
	defer client.Close()
	// gin http server
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

	// 绑定 cron ( input binding )
	r.Any("/demo-cron", func(c *gin.Context) {
		log.Printf("service method invoked app-id:%s,methodName:", "myapp", "echo")
		timeString := time.Now().Format("2006-01-02 15:04:05")

		// 在K8s中，调用 appID参数为 被调用服务定义的  dapr.io/app-id 标签 ，具体请查看部署yaml或PaaS中的部署定义。
		resp, err := client.InvokeMethodWithContent(context.Background(),
			"daprclient-daprdemos-kind-kind", "echo", "post",
			&dapr.DataContent{Data: []byte("hello " + timeString), ContentType: "text/plain"})
		if err != nil {
			fmt.Println("无法调用 Dapr ServiceInvocation API:", err)
		}
		// 解析响应
		log.Printf("service method invoked response:%s", string(resp))

		_ = client.SaveState(context.Background(), "redis", "myway", resp, nil)

		c.JSON(http.StatusOK, gin.H{"message": "State saved successfully"})
	})

	if err := r.Run(":3000"); err != nil {
		fmt.Printf("error starting server: %v\n", err)
		os.Exit(1)
	}
}
