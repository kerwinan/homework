package main

/*
	1. 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
*/

import (
	"context"
	"errors"
	"fmt"
	"github.com/kerwinan/go/kit/log"
	"golang.org/x/sync/errgroup"
	"homework/errgroup/service"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func main() {
	// init log config
	err := log.Init("./config/log.yaml")
	if err != nil {
		fmt.Printf("init log failed: %v\n", err)
		return
	}
	c := make(chan struct{})
	eg, ctx := errgroup.WithContext(context.Background())
	srv, err := service.NewService(c)
	if err != nil {
		log.X.Errorf("err: %v", err)
		return
	}
	/* TODO 此处是否需要使用 errgroup 进行管理 */
	go report()
	/* start go 1 */
	eg.Go(func() error {
		log.X.Debugf("http start")
		return srv.Start()
	})
	/* start go 2 */
	eg.Go(func() error {
		<-ctx.Done()
		log.X.Warnf("http stop")
		return srv.Stop()
	},
	)
	/* start go 3 */
	ch := make(chan os.Signal, 0)
	eg.Go(func() error {
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case sig := <-ch:
				return errors.New(fmt.Sprintf("receive exit signal: %v", sig))
			}
		}
	})
	log.X.Debugf("main exit: ", eg.Wait())
}

func report() {
	ticker := time.NewTicker(time.Second * 60)
	for _ = range ticker.C {
		memStat := new(runtime.MemStats)
		runtime.ReadMemStats(memStat)
		info_msg := fmt.Sprintf("使用内存 %d KB, GO程数: %d.", memStat.Alloc/1024, runtime.NumGoroutine())
		fmt.Println(info_msg)
	}
}
