package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	ctx, stop := context.WithCancel(context.Background())

	go func() {
		defer wg.Done()
		watchDog(ctx, "[监控狗]")
	}()
	time.Sleep(5 * time.Second)
	stop()
	wg.Wait()
}

func watchDog(ctx context.Context, name string) {
	select {
	case <-ctx.Done():
		fmt.Println(name, "停止指令已收到,马上停止")
		return
	default:
		fmt.Println(name, "正在监控...")
	}
	time.Sleep(1 * time.Second)
}
