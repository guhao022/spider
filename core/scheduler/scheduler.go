package scheduler
import "spider/core/context"

type Scheduler interface {
	Add(req *context.Request)
	Get() *context.Request
	SetMaxDepth(depth uint)
	Remain() int
	Total() int
}