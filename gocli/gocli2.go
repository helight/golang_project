package main

import (
	"fmt"
	"log"
	"os"
    "time"

	"github.com/urfave/cli"
)

func main() {
	tasks := []string{"cook", "clean", "laundry", "eat", "sleep", "code"}

	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name:    "complete",
			Aliases: []string{"c"},
			Usage:   "complete a task on the list",
			Action: func(c *cli.Context) error {
				fmt.Println("completed task: ", c.Args().First())
				return nil
			},
			BashComplete: func(c *cli.Context) {
				// This will complete if no args are passed
				if c.NArg() > 0 {
					return
				}
				for _, t := range tasks {
					fmt.Println(t)
				}
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	// 获取当前时间的 Unix 时间戳
	now := time.Now().Unix()

	// 将 Unix 时间戳转换为 time 类型
	t := time.Unix(now, 0)

	fmt.Println(now) // 输出当前时间的 Unix 时间戳
	fmt.Println(t)   // 输出转换后的 time 类型
}
