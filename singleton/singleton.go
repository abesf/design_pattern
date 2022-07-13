package singleton

import (
	"sync"
	"sync/atomic"
)

type Singleton struct{}
var instance *Singleton
var mu sync.Mutex
var initialized uint32


//Non concurrent security
func NewSingleton1() *Singleton {
	if instance!=nil{
		return instance
	}
	return &Singleton{}
}
//Concurrency security but no good enough ,Every call will try to lock
//If the instance exists, it is unnecessary to lock it
func NewSingleton2() *Singleton {
	// Can optimize
	mu.Lock()
	defer mu.Unlock()
	if instance!=nil{
		return instance
	}
	return &Singleton{}
}
// Complete singleton mode
func NewSingleton3() *Singleton {
	// Can optimize
	if atomic.LoadUint32(&initialized) == 1 {  // 原子操作
		return instance
	}
	mu.Lock()
	defer mu.Unlock()
	if initialized == 0 {
		instance = &Singleton{}
		atomic.StoreUint32(&initialized, 1)
	}

	return &Singleton{}
}