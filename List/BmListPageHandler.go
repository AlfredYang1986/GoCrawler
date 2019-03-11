package List

import (
	"github.com/alfredyang1986/GoCrawler/Page"
	"strconv"
	"net/http"
	"github.com/antchfx/antch"
	"github.com/antchfx/htmlquery"
)

type BmListFinishPipeline struct{
	ch chan <- struct{}
	c chan <- []string
}

func (p *BmListFinishPipeline) ServePipeline(v antch.Item) {
	close(p.ch)
	p.c <- v.([]string)
}

func newOutputPipeline(end chan<-struct{}, ret chan<-[]string) antch.Pipeline {
	return func(next antch.PipelineHandler) antch.PipelineHandler {
		return &BmListFinishPipeline{ch : end, c : ret}
	}
}

type BmListPageHandler struct {
	Args map[string]string
	DetailHandler *Page.BmDetailPageHandler
}

func (h BmListPageHandler) NewListPageHandler(a map[string]string, dt *Page.BmDetailPageHandler) *BmListPageHandler {
	return &BmListPageHandler{ Args: a, DetailHandler: dt }
}

func (h *BmListPageHandler) StartCrawlerList(url string) {
	plb, _ := strconv.Atoi(h.Args["plb"])
	pub, _ := strconv.Atoi(h.Args["pub"])

	var details []string
	for page := plb; page <= pub; page++ {
		details = append(details, h.NextPage(url, page)...)
	}

	// TODO: 等全部测试好了之后在用50个chan buffer，如果被封了在想办法
	//pool := make(chan []string, 50)
	pool := make(chan []string, 1)
	defer close(pool)
	content := details[0]
	//for _, content := range details {
		go h.CrawlerListContent(pool, content)
	//go h.CrawlerListContent(pool, content)
	//}

	var result []string
	// TODO: 等全部测试好了之后在用50个chan buffer，如果被封了在想办法
	//for page := plb; page <= pub; page++ {
	for page := plb; page <= 1; page++ {
		result = append(result, <-pool...)
	}

	// TODO: 将result全部弄出来，一次十条一次利用chan一次性处理
	h.DetailHandler.StartCrawlerWithDetail(result)
}

func (h *BmListPageHandler) NextPage(url string, page int) []string {
	uri := url + "p" + strconv.Itoa(page)
	return []string{ uri }
}

func (h *BmListPageHandler) CrawlerListContent(c chan <- []string, url string) {

	ch := make(chan struct{})

	crawler := &antch.Crawler{ Exit: ch }
	crawler.UseCompression()

	crawler.UserAgent = h.Args["useragent"]
	crawler.Handle(h.Args["handlerpattern"], h)
	crawler.UsePipeline(newOutputPipeline(ch, c))

	go func() {
		crawler.StartURLs([]string{ url })
	}()
	<-crawler.Exit
}

func (h *BmListPageHandler) ServeSpider(c chan<- antch.Item, res *http.Response) {
	doc, err := antch.ParseHTML(res)
	if err != nil {
		panic(err)
	}

	var ret []string
	for _, node := range htmlquery.Find(doc, h.Args["linkqueryparam"]) {
		link := htmlquery.SelectAttr(node, "href")
		ret = append(ret, link[2:])
	}
	c <- ret
}

