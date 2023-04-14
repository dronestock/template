package main

import (
	"context"
)

type stepTodo struct {
	*plugin
}

// 要不要传入 plugin 可以根据自己的情况来定，如果要使用 plugin 相关方法或者字段，就可以传入
func newStepTodo(plugin *plugin) *stepTodo {
	return &stepTodo{
		plugin: plugin,
	}
}

func (st *stepTodo) Runnable() bool {
	return true
}

func (st *stepTodo) Run(_ context.Context) (err error) {
	return
}
