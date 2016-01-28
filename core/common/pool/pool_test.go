package pool

import (
	"testing"
	"math"
	"fmt"

)

func Test_Pool(t *testing.T) {
	p := NewPool(new(en), 10)

	p.Take()

	fmt.Printf("实体池总容量: %d\n",p.Cap())

	fmt.Printf("实体池使用量: %d\n", p.Used())

	p.Take()
	p.Take()
	p.Take()
	p.Take()
	p.Take()
	p.Take()
	p.Take()
	p.Take()
	e := p.Take()

	go p.Free(e)

	fmt.Printf("实体池总容量: %d\n",p.Cap())
	fmt.Printf("实体池使用量: %d\n", p.Used())
	p.Take()
	fmt.Printf("实体池总容量: %d\n",p.Cap())
	fmt.Printf("实体池使用量: %d\n", p.Used())
}

type en struct {
	Id int
}

func (e *en) New() Entity {
	id := 0
	if int32(id) < math.MaxInt32 {
		id++
	}
	return &en{
		Id: id,
	}
}

func (e *en) Clean(){}
func (e *en) Close(){}

func (e *en) Usable() bool {
	return true
}
