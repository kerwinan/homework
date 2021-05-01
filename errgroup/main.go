package main

/*
	1. 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
*/

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/martian/log"
	"golang.org/x/sync/errgroup"
	"homework/errgroup/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c := make(chan struct{})
	g, ctx := errgroup.WithContext(context.Background())
	srv, err := service.NewService(c)
	if err != nil {
		fmt.Println(err)
		return
	}
	/* start go 1 */
	g.Go(func() error {
		return srv.Start()
	})
	/* start go 2 */
	g.Go(func() error {
		select {
		case <-ctx.Done():
			log.Errorf("errgroup exit")
		case <-c:
			log.Errorf("http server shutdown")
		}
		return srv.Stop()
	},
	)
	/* start go 3 */
	ch := make(chan os.Signal, 0)
	g.Go(func() error {
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
		select {
		case <-ctx.Done():
			return ctx.Err()
		case sig := <-ch:
			return errors.New(fmt.Sprintf("receive exit signal: %v", sig))
		}
	})
	fmt.Println("main exit: ", g.Wait())
}
