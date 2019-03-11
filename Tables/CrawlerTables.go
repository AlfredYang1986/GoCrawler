package Tables

import (
	"github.com/alfredyang1986/GoCrawler/List"
	"github.com/alfredyang1986/GoCrawler/Page"
	"github.com/alfredyang1986/GoCrawler/Entry"
)

type CrawlerTable struct{}

var CRAWLER_LIST_HANDLER_FACTORY = map[string]interface{}{
	"BmListPageHandler":                 List.BmListPageHandler{},
}

var CRAWLER_PAGE_HANDLER_FACTORY = map[string]interface{}{
	"BmDetailPageHandler":                 Page.BmDetailPageHandler{},
}

var CRAWLER_ENTRY_HANDLER_FACTORY = map[string]interface{}{
	"BmEntry":								Entry.BmEntry{},
}

func (t CrawlerTable) GetListHandlerByName(name string) interface{} {
	return CRAWLER_LIST_HANDLER_FACTORY[name]
}

func (t CrawlerTable) GetPageHandlerByName(name string) interface{} {
	return CRAWLER_PAGE_HANDLER_FACTORY[name]
}

func (t CrawlerTable) GetEntriesByName(name string) interface{} {
	return CRAWLER_ENTRY_HANDLER_FACTORY[name]
}
