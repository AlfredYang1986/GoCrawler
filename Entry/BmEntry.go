package Entry

import (
	"fmt"
	"github.com/alfredyang1986/GoCrawler/TypeDef"
)

type BmEntry struct {
	Url string
	Handler TypeDef.ListHandler
}

func (h BmEntry) NewBmEntryHandler(u string, hd interface{}) *BmEntry{
	return &BmEntry{ Url: u, Handler: hd.(TypeDef.ListHandler) }
}

func (h *BmEntry) StartCrawler() {
	fmt.Println("starts")
	h.Handler.StartCrawlerList(h.Url)
}