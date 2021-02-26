package workpool

import (
	"fmt"
	"time"
)

//协程池结构
type pool struct {
	//接受任务的管道
	taskChan chan *task
	//最大协程数
	maxChan int
	//任务管道序列
	jobChan chan *task
}

//创建协程池
func NewPool(capacity int) *pool {
	return &pool{make(chan *task), capacity, make(chan *task)}
}

//任务结构
type task struct {
	f func() error
}

//创建任务
func NewTask(f func() error) *task {
	return &task{f}
}

//执行任务
func (t *task) execute() {
	err := t.f()
	if err != nil {
		fmt.Println("Execute err")
	}
}

//开启协程从管道中读取并执行任务
func (p *pool) start(taskId int) {
	//从jobchan管道中读取任务并执行
	for t := range p.jobChan {
		t.execute()
		fmt.Println("taskId ID ", taskId, " 执行完毕任务")
	}
}

//让协程池Pool开始工作
func (p *pool) Run() {
	//开启最大数量的协程
	for i := 0; i < p.maxChan; i++ {
		go p.start(i)
	}

	// 将taskChan中的值传入jobChan里
	//range p.taskChan产生的迭代值为Channel中发送的值，它会一直迭代直到channel被关闭
	for task := range p.taskChan {
		p.jobChan <- task
	}

	//执行完毕关闭管道
	close(p.jobChan)
	close(p.taskChan)
}

func PoolTest() {
	//创建一个Task
	//这里使用传入一个打印时间的匿名函数
	t := NewTask(func() error {
		fmt.Println(time.Now().Clock())
		return nil
	})

	//创建一个协程池,最大开启3个协程worker
	p := NewPool(3)

	//开一个协程不断的向 Pool 输送打印一条时间的task任务
	go func() {
		for {
			p.taskChan <- t
			time.Sleep(1 * time.Second)
		}
	}()
	//启动协程池p
	p.Run()
}
