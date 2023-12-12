package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type router struct{}

func (engine *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	time.Sleep(time.Second * 20)
	w.WriteHeader(200)
	w.Write([]byte("ok\n"))
}

func main() {
	engine := new(router)
	server := &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}

	// http服务
	go func() {
		log.Println("HTTP服务启动", "http://localhost"+server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println(err)
			os.Exit(1)
		}
		log.Println("HTTP服务关闭")
	}()

	// 监听信号 优雅退出http服务
	Watch(func() error {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		go func() {
			ticker := time.NewTicker(time.Second)
			defer ticker.Stop()

			acc := 1
			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					log.Printf("%ds\n", acc)
					acc++
				}
			}
		}()
		return server.Shutdown(ctx)
	})

	log.Println("程序退出")
}

// Watch 监听信号
func Watch(fns ...func() error) {
	// 程序无法捕获信号 SIGKILL 和 SIGSTOP （终止和暂停进程），因此 os/signal 包对这两个信号无效。
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)

	// 阻塞
	s := <-ch
	close(ch)
	log.Println("接收到信号", s.String(), "执行关闭函数")
	for i := range fns {
		if err := fns[i](); err != nil {
			log.Println("fn() error:", err)
		}
	}
	log.Println("关闭函数执行完成")
}
