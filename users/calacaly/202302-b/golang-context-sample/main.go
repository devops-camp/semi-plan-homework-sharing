package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const LINE = "================"

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	rand.Seed(time.Now().UnixNano())
	for true {
		time.Sleep(500 * time.Millisecond)
		select {
		case <-ctx.Done():
			os.Exit(1)
		default:
			go Lubei(ctx, rand.Intn(100))
		}
	}
	cancel()
}

func Lubei(ctx context.Context, n int) {
	select {
	case <-ctx.Done():
		os.Exit(1)
	default:
		ctx = context.WithValue(ctx, "num", n)
		fmt.Printf("\n%s\n", LINE)
		fmt.Printf("刘备: 曹操来了%d万人\n", n)
		go Guanyu(ctx)
	}

}

func Guanyu(ctx context.Context) {
	select {
	case <-ctx.Done():
		os.Exit(1)
	default:
		num := ctx.Value("num").(int)
		fmt.Printf("关羽: 曹操来了%d万人\n", num)
		if num%2 == 0 {
			go Zhangfei(ctx)
		} else {
			ctx = context.WithValue(ctx, "num", 10*num)
			go Zhangfei(ctx)
		}
	}
}

func Zhangfei(ctx context.Context) {
	select {
	case <-ctx.Done():
		os.Exit(1)
	default:
		fmt.Printf("张飞: 曹操来了%d万人\n", ctx.Value("num").(int))
	}

}
