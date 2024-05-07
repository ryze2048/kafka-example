package main

import (
	"context"
	"fmt"
	"github.com/ryze2048/kafka-example/global"
	"github.com/ryze2048/kafka-example/initialize"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var parentCtx = context.Background()
	var ctx, cancel = context.WithCancel(parentCtx)

	initialize.InitLog()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalCh
		fmt.Println("接收到终止信号，正在优雅地退出...")
		global.ZAPLOG.Info("接收到终止信号，正在优雅地退出...")
		cancel() // 取消操作
	}()

	<-ctx.Done()
	global.ZAPLOG.Info("程序结束")
}
