package main

import (
	"context"
	"fmt"
	"os"
	"time"

	dapr "github.com/dapr/go-sdk/client"
)

func main() {
	// 创建 Dapr 客户端
	client, err := dapr.NewClient()
	if err != nil {
		fmt.Println("无法创建 Dapr 客户端:", err)
		os.Exit(1)
	}
	defer client.Close()

	// 准备 ServiceInvocation 请求
	for {
		// 调用 ServiceInvocation API
		resp, err := client.InvokeMethodWithContent(context.Background(),
			"myapp", "echo", "post",
			&dapr.DataContent{Data: []byte("hello"), ContentType: "text/plain"})
		if err != nil {
			fmt.Println("无法调用 Dapr ServiceInvocation API:", err)
		}
		// 解析响应
		fmt.Printf("service method invoked, response: %s\n", string(resp))
		time.Sleep(5 * time.Second)
	}

	fmt.Println("程序已退出")
}
