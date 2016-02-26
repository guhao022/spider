package analyzer

// 条目
type Item map[string]interface{}

func (this *Item) AddItem(key string, item interface{}) {
	this[key] = item
}

func (this *Item) GetItem(key string) (interface{}, bool) {
	i, ok := this[key]
	return i, ok
}

