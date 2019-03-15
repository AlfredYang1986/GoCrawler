package Page

import (
	"encoding/csv"
	"fmt"
	"github.com/antchfx/antch"
	"github.com/antchfx/htmlquery"
	"net/http"
	"os"
	"strconv"
	"time"
)
var num int
var w * csv.Writer
type BmDetailFinishPipeline struct{
	ch chan <- struct{}
	c chan <- []string
}

func (p *BmDetailFinishPipeline) ServePipeline(v antch.Item) {
	close(p.ch)
	p.c <- v.([]string)
}

func newOutputPipeline(end chan<-struct{}, ret chan<-[]string) antch.Pipeline {
	return func(next antch.PipelineHandler) antch.PipelineHandler {
		return &BmDetailFinishPipeline{ch : end, c : ret}
	}
}

type BmDetailPageHandler struct {
	args map[string]string
}

func (h BmDetailPageHandler) NewDetailPageHandler(a map[string]string) *BmDetailPageHandler {
	return &BmDetailPageHandler{ args: a }
}

func (h *BmDetailPageHandler) StartCrawlerWithDetail(urls []string) {
	fmt.Println("BmDetailPageHandler")
	num=-1
	f, err := os.Create("DianpingShop.csv") //创建文件
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	w = csv.NewWriter(f) //创建一个新的写入文件流
	w.Write([]string{"商户名称", "分类", "区域","人均消费","点评数量","创立时间","商户介绍","联系电话","商户地址","分店数量"})  
	var result []string
	i:=0
	num=i
	link:= urls[i]
	pool := make(chan []string, 1)
	defer close(pool)
	for {
		link= urls[i]
		if i%492==0||link==urls[len(urls)-1]{
			//time.Sleep(time.Duration(5)*time.Second)
			w.Flush()
			f, err = os.Create(strconv.Itoa(i)+"DianpingShop.csv") //创建文件
			if err != nil {
				panic(err)
			}
			defer f.Close()
			f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
			w = csv.NewWriter(f) //创建一个新的写入文件流
			w.Write([]string{"商户名称", "分类", "区域","人均消费","点评数量","创立时间","商户介绍","联系电话","商户地址","分店数量"})
		}
		i++
		if i%30!=0{		
			go h.CrawlerListContent(pool,link)
		}else {
			time.Sleep(time.Duration(10)*time.Second)
			continue
		}	
		//for page := 1; page <= 30; page++ {
			result = append(result, <-pool...)
		//}
	}
}
func (h *BmDetailPageHandler) CrawlerListContent(c chan <- []string, url string) {

	ch := make(chan struct{})

	crawler := &antch.Crawler{ Exit: ch }
	crawler.UseCompression()

	crawler.UserAgent = h.args["useragent"]
	crawler.Handle("dianping.com", h)
	crawler.UsePipeline(newOutputPipeline(ch, c))

	//go func() {
		crawler.StartURLs([]string{ url })
	//}()
	<-crawler.Exit
}

func (h *BmDetailPageHandler) ServeSpider(c chan<- antch.Item, res *http.Response) {
	doc, err := antch.ParseHTML(res)
	num++
	var AverageSpend string
	var Introduction string
	var CreateTime string
	var ShopName string
	var Type string
	var Region string
	var ContentCount string
	var Phone string
	var ShopAddress string
	var BranchNum string
	if err != nil {
		panic(err)
	}
	node := htmlquery.FindOne(doc, h.args["shopnameparam"])
	ShopName = htmlquery.InnerText(node)

	node = htmlquery.FindOne(doc, h.args["typeparam"])
	if node!=nil{
		Type = htmlquery.InnerText(node)
	}else {
		Type = "-"
	}

	node =htmlquery.FindOne(doc, h.args["regionparam"])
	if node!=nil{
		Region = htmlquery.InnerText(node)
		Region=Region[0:len(Region)-2]
	}else {
		Region = "-"
	}
	
	node = htmlquery.FindOne(doc, h.args["priceparam"])
	if node!=nil{
		AverageSpend = htmlquery.InnerText(node)
	}else {
		AverageSpend = "-"
	}
	node =htmlquery.FindOne(doc, h.args["contentcountparam"])
	if node!=nil{
		ContentCount = htmlquery.InnerText(node)
	}else {
		ContentCount = "-"
	}

	node=htmlquery.FindOne(doc, h.args["createtimeparam"])
	if node!=nil{
		CreateTime = htmlquery.InnerText(node)
	}else {
		CreateTime = "-"
	}

	node=htmlquery.FindOne(doc, h.args["introductionparam"])
	if node != nil{
		Introduction = htmlquery.InnerText(node)
	}else {
		Introduction = "-"
	}
	//fmt.Println(Introduction)
	node=htmlquery.FindOne(doc, h.args["phoneparam"])
	if node != nil{
		Phone = htmlquery.InnerText(node)
	}else {
		Phone = "-"
	}

	node=htmlquery.FindOne(doc, h.args["addressparam"])
	if node != nil{
		ShopAddress = htmlquery.InnerText(node)
	}else {
		ShopAddress = "-"
	}

	node=htmlquery.FindOne(doc, h.args["branchcount"])
	if node!=nil{
		BranchNum = htmlquery.InnerText(node)
	}else {
		BranchNum = "-"
	}

	w.Write([]string{ShopName,Type,Region,AverageSpend,ContentCount,CreateTime,Introduction,Phone,ShopAddress,BranchNum})
	fmt.Println(num)
	var ret []string
	c <- ret
}