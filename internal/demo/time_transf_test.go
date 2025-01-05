package demo

import (
	"fmt"
	"time"
)

func testTime() {
	// 获取当前时间
	currentTime := time.Now()
	fmt.Println("当前时间 (UTC):", currentTime)

	// 加载指定时区（例如，Asia/Shanghai）
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("加载时区失败:", err)
		return
	}

	// 转换到指定时区
	shanghaiTime := currentTime.In(location)
	fmt.Println("上海时间:", shanghaiTime)
}
