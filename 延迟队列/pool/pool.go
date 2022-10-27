package pool

import (
	"context"
	"delay_demo/pack"
	"errors"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

var (
	// return if pool size
	ErrInvalidPoolCap = errors.New("invalid pool cap")
	// put task but pool already closed
	ErrPoolAlreadyClosed = errors.New("pool already closed")
)

const (
	RUNNING = 1
	STOPED  = 0
)

type Task struct {
	Member string
}

//Pool task pool
type Pool struct {
	Capacity       uint64 //容量
	RunningWorkers uint64 //正在工作的
	Status         int64  //状态
	ChTask         chan *Task
	PanicHandler   func(interface{})
	sync.Mutex
}

// NewPool init pool
func NewPool(capacity uint64) (*Pool, error) {
	if capacity <= 0 {
		return nil, ErrInvalidPoolCap
	}
	p := &Pool{
		Capacity: capacity,
		Status:   RUNNING,
		ChTask:   make(chan *Task, capacity),
	}
	return p, nil
}

func (p *Pool) checkWorker() {
	p.Lock()
	defer p.Unlock()
	if p.RunningWorkers == 0 && len(p.ChTask) > 0 {
		p.run()
	}
}
func (p *Pool) GetCap() uint64 {
	return p.Capacity
}

//GetRunningWorkers get running workers
func (p *Pool) GetRunningWorkers() uint64 {
	return atomic.LoadUint64(&p.RunningWorkers)
}

func (p *Pool) incRunning() {
	atomic.AddUint64(&p.RunningWorkers, 1)
}

func (p *Pool) decRunning() {
	atomic.AddUint64(&p.RunningWorkers, ^uint64(0))
}

//Put put a task to pool
func (p *Pool) Put(task *Task) error {
	p.Lock()
	defer p.Unlock()
	if p.Status == STOPED {
		return ErrPoolAlreadyClosed
	}
	//run workers
	if p.GetRunningWorkers() < p.GetCap() {
		p.run()
	}
	//send task
	if p.Status == RUNNING {
		p.ChTask <- task
	}
	return nil
}

func (p *Pool) run() {
	p.incRunning()
	redisCoon := pack.GetRedisDb()
	ctx := context.Background()
	go func() {
		defer func() {
			p.decRunning()
			if r := recover(); r != nil {
				if p.PanicHandler != nil {
					p.PanicHandler(r)
				} else {
					log.Printf("Worker panic: %s\n", r)
				}
			}
			p.checkWorker() //check worker avoid no worker running
		}()
		for {
			select {
			case _, ok := <-p.ChTask:
				if !ok {
					fmt.Println("OK")
					return
				}
				err := ConsumeQueue(ctx, redisCoon)
				if err != nil {
					fmt.Println("消费队列发生错误：", err)
					return
				}

			}
		}
	}()

}

func (p *Pool) setStatus(status int64) bool {
	p.Lock()
	defer p.Unlock()
	if p.Status == status {
		return false
	}
	p.Status = status
	return true
}

//Close close pool graceful
func (p *Pool) Close() {
	if !p.setStatus(STOPED) {
		//stop put task
		return
	}
	for len(p.ChTask) > 0 {
		//wait all task be consumed
		time.Sleep(1e6) //reduce cpu load
	}
	close(p.ChTask)
}
