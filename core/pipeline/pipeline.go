package pipeline

import (
	"fmt"
	"sync/atomic"
	"spider/core/data"
)

// 条目处理管道
type ItemPipeline interface {
	Send(item data.Item) []error
	Count() []uint64
	ProcessingNumber() uint64
	Summary() string
}

func NewItemPipeline(itemProcessors []ProcessItem) ItemPipeline {
	if itemProcessors == nil {
		panic(fmt.Errorf("Invalid item processor list!"))
	}
	innerItemProcessors := make([]ProcessItem, 0)
	for i, ip := range itemProcessors {
		if ip == nil {
			panic(fmt.Errorf("Invalid item processor[%d]!\n", i))
		}
		innerItemProcessors = append(innerItemProcessors, ip)
	}
	return &myItemPipeline{itemProcessors: innerItemProcessors}
}

type myItemPipeline struct {
	itemProcessors   []ProcessItem // 条目处理器的列表。
	sent             uint64        // 已被发送的条目的数量。
	accepted         uint64        // 已被接受的条目的数量。
	processed        uint64        // 已被处理的条目的数量。
	processingNumber uint64        // 正在被处理的条目的数量。
}

func (this *myItemPipeline) Send(item data.Item) []error {
	return nil
}

func (this *myItemPipeline) Count() []uint64 {
	counts := make([]uint64, 3)
	counts[0] = atomic.LoadUint64(&this.sent)
	counts[1] = atomic.LoadUint64(&this.accepted)
	counts[2] = atomic.LoadUint64(&this.processed)
	return counts
}

func (this *myItemPipeline) ProcessingNumber() uint64 {
	return atomic.LoadUint64(&this.processingNumber)
}

func (this *myItemPipeline) Summary() string {
	counts := this.Count()
	summary := fmt.Sprintf("processorNumber: %d, sent: %d, accepted: %d, processed: %d, processing: %d", len(this.itemProcessors), counts[0], counts[1], counts[2], this.ProcessingNumber())

	return summary
}




