package main

import (
	"bufio"
	"fmt"
	"github.com/eavesmy/log_monitor/lib"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

var tasks = sync.Map{}

var scripts_path = ""

func main() {

	if len(os.Args) >= 2 {
		scripts_path = os.Args[1]
	}

	lib.LoadScripts(scripts_path)

	cmd := exec.Command("sh", "-c", "tail -f /root/source/server/game-server/logs/*.log")

	stdout, err := cmd.StdoutPipe()

	if err != nil { //获取输出对象，可以从该对象中读取输出结果
		fmt.Println("stdout error", err)
		panic(err)
	}

	go notify()

	go func() {
		reader := bufio.NewReader(stdout)
		queue := &lib.Queue{}

		for {

			b, _, err := reader.ReadLine()

			// buffer 满了则从前面取出

			if err != nil || err == io.EOF {
				// fmt.Println(string(b), err)
				continue
			}

			if queue.Len() > 10000 {
				queue.Drop()
			}

			line := string(b)
			queue.Push(line)

			// 检查有无关键字
			// 有关键字则创建任务并执行

			// 遍历所有脚本
			lib.Scripts.Range(func(_, script interface{}) bool {

				ins := script.(*lib.Script)

				for _, key := range ins.Keywords {
					if strings.Index(line, key) > -1 {
						ins.Run(line)
					}
				}
				return true
			})

			/*
				tasks.Range(func(k, v interface{}) bool {
					if err = v.(*lib.Task).Accept(b); err != nil {
						fmt.Println("run script error", err)
					}
					return true
				})
			*/
		}
	}()

	if err = cmd.Run(); err != nil {
		fmt.Println("cmd error", err)
	}

}

func notify() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGUSR1)

	for {
		sig := <-sigs
		fmt.Println("Receive sig", sig)

		lib.LoadScripts(scripts_path)
	}

}
