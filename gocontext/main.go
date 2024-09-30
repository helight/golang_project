package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	luckyNum := 6
	go guessNum(ctx, luckyNum)

	time.Sleep(5 * time.Second)

	cancel()

	time.Sleep(1 * time.Second)
}

func guessNum(ctx context.Context, luckyNum int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Do not guess the right num so we return")
			return
		default:
			randnum := rand.Intn(10)
			if randnum == luckyNum {
				fmt.Println("guess the right num :", randnum, " good job!")
				return
			} else {
				fmt.Println("guess the wrong num :", randnum, " try again 1 second later")
				time.Sleep(1 * time.Second)
			}
		}
	}
}
