package main

import (
	"context"
	"fmt"
	"go-micro.dev/v5"
	"go-micro.dev/v5/server"
	"log"
	"runtime"
)

func main() {
	service := micro.NewService(
		micro.Address(":8080"),
		micro.Name("helloworld"),
		micro.Handle(new(Helloworld)),
		micro.WrapHandler(Recover, HelloServer, LogWrapper),
	)

	service.Init()

	service.Run()

	panic("something went wrong") // cho nay bi panic, nhung go micro se tu dong recover
}

func LogWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		log.Print("logWrapper middleware")
		return fn(ctx, req, resp)
	}
}

func HelloServer(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		log.Print("helloServer")
		return fn(ctx, req, resp)
	}
}

func Recover(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		defer func() {
			if r := recover(); r != nil {
				buf := make([]byte, 1024) // Buffer để lưu stack trace
				n := runtime.Stack(buf, false)
				fmt.Printf("Panic recovered: %v\n", r)
				fmt.Printf("Stack trace:\n%s\n", buf[:n])
			}
		}()

		return fn(ctx, req, resp)
	}
}

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
}

type Helloworld struct{}

func (h *Helloworld) Greeting(ctx context.Context, req *Request, rsp *Response) error {
	panic("something went wrong") // cho nay bi panic, nhung go micro se tu dong recover
}

func (h *Helloworld) Bye(ctx context.Context, req *Request, rsp *Response) error {
	rsp.Message = "Bye " + req.Name
	return nil
}
