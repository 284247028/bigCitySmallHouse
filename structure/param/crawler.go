package param

type CrawlerParam struct {
	Type string //	爬取类型，单个-列表
	Url  string // 指定房源详情页
	Page int    // 爬取列表时，第几页
}
