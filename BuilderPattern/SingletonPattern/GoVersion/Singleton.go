package SPGoversion

import "sync"

/**
	饿汉模式
 */
type hSingleton struct {
	name string
}
var hSt *hSingleton

func init() {
	hSt = &hSingleton{}
}

func GetHungerInstance() *hSingleton{
	return hSt
}

func (h *hSingleton) SetName(name string) {
	h.name = name
}

/**
	懒汉模式
 */
type lSingleton struct {
	name string
}
var lazyInstance *lSingleton

func GetLazyInstance() *lSingleton{
	if lazyInstance == nil {
		lazyInstance = &lSingleton{}
	}
	return lazyInstance
}

func (l *lSingleton) SetName(name string) {
	l.name = name
}
/**
懒汉+多线程加锁单例
 */
type singleton struct {
}
var lock sync.Mutex
var instance *singleton
func GetMultiThreadInstance(name string) *singleton{
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &singleton{}
		}
	}
	return instance
}
/**
通过sync.Once来确保只创建一次
内部本质上也是双重检查的方式，但在写法上会比自己写双重检查更简洁
 */
var once sync.Once

func GetOnceInstance() *singleton {
	once.Do(func() {
		instance = new(singleton)
	})
	return instance
}