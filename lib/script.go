package lib

import (
	"fmt"
	"github.com/robertkrimen/otto"
	"io/ioutil"
	"path"
	"strings"
	"sync"
)

var Scripts = sync.Map{}

type Script struct {
	Script   *otto.Script
	Vm       *otto.Otto
	Keywords []string
}

func (script *Script) Run(line string) {

	script.Vm.Set("line", line)
	script.Vm.Run(script.Script)
}

func LoadScripts(scripts_path string) {
	// 声明接口，由脚本实现接口

	Scripts = sync.Map{}

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

		vm := otto.New()
		script, err := vm.Compile(file_path, nil)

		vm.Set("stacks", nil)
		vm.Run(script) // init vm env

		if err != nil {
			fmt.Println("laod script failed", scripts_path)
			continue
		}

		v, err := vm.Get("keywords")
		if err != nil {
			fmt.Println("script keywords error", err)
		}
		str_v, err := v.ToString()
		if (err != nil) || str_v == "undefined" {
			fmt.Println("script keywords undefined", err)
		}

		Scripts.Store(file.Name(), &Script{
			Script:   script,
			Vm:       vm,
			Keywords: strings.Split(str_v, ","),
		})
	}
}
