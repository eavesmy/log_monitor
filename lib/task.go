package lib

// 每个任务对应一个处理脚本。
// 提供单独的堆栈
// 任务声明接口，由脚本来实现接口
// 任务提供基本的功能可以由脚本来调用。
// 每个任务可能会产生多个子任务，每产生一个子任务由父任务进行管理
// 需要编排任务的生命周期:
//      1. 接收流
//      2. 调用脚本判断存入堆栈还是直接处理 <- 通过脚本返回的状态来执行对应步骤操作
//      3.

// eg.
//	1. 获取到 结算 行,脚本返回 {id: 2, str: []string{"UPLOAD SUCCESS",312243}}
//  2. 任务调用总线的栈查找符合条件的行,并将结果传入脚本
//  3. 脚本返回结果，并调用对应方法进行处理

import (
	"fmt"
	"github.com/eavesmy/golang-lib/crypto"
	"github.com/robertkrimen/otto"
	"sync"
)

const (
	TASK_SAVE   = 1
	TASK_SEARCH = 2
)

// task 是每一个监控任务
type Task struct {
	Lines    sync.Map
	Active   bool
	Script   *otto.Script
	FilePath string
	vm       *otto.Otto
}

func NewTask(file_path string) *Task {

	vm := otto.New()
	task := &Task{Lines: sync.Map{}, vm: vm, FilePath: file_path}

	task.Reload()
	return task
}

func (t *Task) Accept(line []byte) error {
	if t.Script == nil || t.vm == nil {
		return nil
	}

	t.vm.Set("line", string(line))
	value, err := t.vm.Run(t.Script)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println("return ::", value)

	// switch value {
	// value.Object()
	// }

	return nil
}

func (t *Task) Save(line []byte) {
	md5 := crypto.Md5_16(string(line))
	t.Lines.Store(md5, line)
}

func (t *Task) Search(str string) {
	// queue 中查找
}

func (t *Task) Empty() {

}

func (t *Task) Reload() *otto.Script { // reload

	script, err := t.vm.Compile(t.FilePath, nil)

	if err != nil {
		fmt.Println("compile script failed", t.FilePath, err)
	}

	t.Script = script

	return script
}
