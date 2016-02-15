package pool

import (
	"sync"
	"time"
)

// 实体的接口类型。
type Entity interface {
	New()	Entity
	Close()
	Clean()
	Usable() bool
}


// 实体池接口
type Pool interface {
	// 获取实体
	Take() Entity
	// 释放实体
	Free(Entity)
	// 实体池容量
	Cap() int
	// 实体池使用量
	Used() int
}

// 创建实体池
func NewPool(entity Entity, size ...int) Pool {
	if len(size) <= 0 {
		size = append(size, 1024)
	}

	return &pool{
		cap: size[0] - 1,
		entity: entity,
		container: make(map[Entity]bool),
	}
}

// 实体池类型
type pool struct {
	// 池的总容量。
	cap		int
	// 实体
	entity	Entity
	// 实体容器
	container map[Entity]bool
	// 互斥锁
	mu	sync.Mutex
}

// 取出实体
func (p *pool) Take() Entity {
	p.mu.Lock()
	defer p.mu.Unlock()

	for {
		for k, v := range p.container {
			if v {
				continue
			}
			if !k.Usable() {
				p.Remove(k)
				continue
			}
			p.container[k] = true
			return k
		}

		if len(p.container) <= p.cap {
			p.increment()
		} else {
			time.Sleep(time.Second)
		}
	}

}

// 释放实体
func (p *pool) Free(entity Entity) {
	defer func() {
		recover()
	}()
	entity.Clean()

	p.container[entity] = false

}

func (p *pool) Remove(entity Entity) {
	defer func() {
		recover()
	}()
	entity.Close()
	delete(p.container, entity)
}

// 重置资源池
func (p *pool) Reset() {
	for k, _ := range p.container {
		k.Close()
		delete(p.container, k)
	}
}

// 实体池总容量
func (p *pool) Cap() int {
	return p.cap + 1
}

// 实体池使用量
func (p *pool) Used() int {
	used := 0
	for _, v := range p.container {
		if v {
			used++
		}
	}
	return used
}

// 动态增加资源
func (p *pool) increment() {
	entity := p.entity.New()
	if entity == nil {
		return
	}

	p.container[entity] = false
}



