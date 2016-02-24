package pipeline

import "spider/core/data"

// 被用来处理条目的函数类型。
type ProcessItem func(item data.Item) (result data.Item, err error)
