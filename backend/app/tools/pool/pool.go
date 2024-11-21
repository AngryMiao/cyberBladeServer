package pool

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
)

var (
	instance *ants.Pool
	once     sync.Once
)

func InitPool(size int) error {
	var err error
	once.Do(func() {
		instance, err = ants.NewPool(size,
			ants.WithPreAlloc(true),
			ants.WithMaxBlockingTasks(size*2),
		)
	})
	return err
}

func Submit(task func()) error {
	if instance == nil {
		return fmt.Errorf("pool not initialized")
	}
	return instance.Submit(task)
}

func Release() {
	if instance != nil {
		instance.Release()
	}
}

func GetPool() *ants.Pool {
	return instance
}
