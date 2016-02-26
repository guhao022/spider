package analyzer

type Analyzer interface {
	// 根据规则解析网页
	Analyze(page *Page) ([]Item, []error)
}




