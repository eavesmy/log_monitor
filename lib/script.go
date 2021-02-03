package lib

import (
	"fmt"
	"github.com/robertkrimen/otto"
	"io/ioutil"
	"os"
	"path"
	"plugin"
	"strings"
	"sync"
)

const (
	ACTION_ALTER = 1
)

var Scripts = sync.Map{}
var Hook plugin.Symbol

type Script struct {
	Script   *otto.Script
	Vm       *otto.Otto
	Keywords []string
}

func (script *Script) Run(line string) {

	script.Vm.Set("line", line)
	script.Vm.Run(script.Script)
	v, _ := script.Vm.Get("result")

	r, _ := v.ToString() // 给个 http 钩子?

	go HookFunc(r)
}

func init() {
	hookProgPath := os.Getenv("MONITOR_HOOK")

	if hookProgPath == "" {
		return
	}

	hook, err := plugin.Open(hookProgPath)
	if err != nil {
		fmt.Println(err)
	}

	f, err := hook.Lookup("Run")

	if err != nil {
		fmt.Println("Run method required!")
		return
	}

	Hook = f

	// 放给一个全局变量

	// hook.Lookup()

	// 加载插件
	// 判断文件是否存在
	// 存在则调用
	// 不存在则略过

}

func LoadScripts(scripts_path string) {
	// 声明接口，由脚本实现接口

	dir, err := ioutil.ReadDir(scripts_path)

	if err != nil {
		panic("dir not found " + scripts_path)
	}

	for _, file := range dir {

		file_path := path.Join(scripts_path, file.Name())
		ext := path.Ext(file_path)

		if ext != ".js" {
			continue
		}

		ins, exists := Scripts.Load(file_path)

		if !exists {
			ins = &Script{}
			vm := otto.New()
			ins.(*Script).Vm = vm
			vm.Set("stacks", nil)
		}

		script, err := ins.(*Script).Vm.Compile(file_path, nil)

		if err != nil {
			fmt.Println("laod script failed", scripts_path)
			continue
		}

		ins.(*Script).Vm.Run(script) // init vm env

		v, err := ins.(*Script).Vm.Get("keywords")
		if err != nil {
			fmt.Println("script keywords error", err)
		}
		str_v, err := v.ToString()
		if (err != nil) || str_v == "undefined" {
			fmt.Println("script keywords undefined", err)
		}

		ins.(*Script).Keywords = strings.Split(str_v, ",")
		ins.(*Script).Script = script

		Scripts.Store(file.Name(), ins)
	}
}

func HookFunc(result string) {
	Hook.(func(string))(result)
}
