package TypeDef

type Entry interface {
	StartCrawler()
}

type ListHandler interface {
	StartCrawlerList(url string)
	NextPage(url string, page int) []string
}

type PageHandler interface {
	StartCrawlerWithDetail(urls []string)
}
