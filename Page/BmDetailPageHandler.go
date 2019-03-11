package Page

import "fmt"

type BmDetailPageHandler struct {
	args map[string]string
}

func (h BmDetailPageHandler) NewDetailPageHandler(a map[string]string) *BmDetailPageHandler {
	return &BmDetailPageHandler{ args: a }
}

func (h *BmDetailPageHandler) StartCrawlerWithDetail(urls []string) {
	fmt.Println("BmDetailPageHandler")
	fmt.Println(urls)
}