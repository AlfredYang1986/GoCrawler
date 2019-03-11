package Config

import (
	"io/ioutil"
	"fmt"
	"github.com/alfredyang1986/GoCrawler/TypeDef"
	"github.com/go-yaml/yaml"
	"github.com/alfredyang1986/BmServiceDef/BmSingleton"
	"github.com/alfredyang1986/GoCrawler/Tables"
	"github.com/alfredyang1986/BmServiceDef/BmPanic"
)

type CrawlerPod struct {
	Name string
	Res  map[string]interface{}
	conf CGConfig

	Factory Tables.CrawlerTable

	Entries map[string]interface{}
	ListHandlers map[string]interface{}
	DetailHandlers map[string]interface{}
}

func (p *CrawlerPod) RegisterSerFromYAML(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("error")
	}
	//check(err)
	p.conf = CGConfig{}
	err = yaml.Unmarshal(data, &p.conf)
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
		panic(BmPanic.ALFRED_TEST_ERROR)
	}

	p.CreateDetailHandlers()
	p.CreateListHandlers()
	p.CreateEntries()
}

func (p *CrawlerPod) CreateDetailHandlers() {
	if p.DetailHandlers== nil {
		p.DetailHandlers= make(map[string]interface{})
	}

	for _, d := range p.conf.Details {
		any := p.Factory.GetPageHandlerByName(d.Name)
		name := d.Method
		args := d.Args
		inc, _ := BmSingleton.GetFactoryInstance().ReflectFunctionCall(any, name, args)
		p.DetailHandlers[d.Name] = inc.Interface()
	}
}

func (p *CrawlerPod) CreateListHandlers() {
	if p.ListHandlers == nil {
		p.ListHandlers = make(map[string]interface{})
	}

	for _, d := range p.conf.Lists {
		any := p.Factory.GetListHandlerByName(d.Name)
		name := d.Method
		args := d.Args

		var tmp TypeDef.PageHandler
		dt := args["detail"]
		if len(dt) > 0 {
			tmp = p.DetailHandlers[dt].(TypeDef.PageHandler)
		}

		inc, _ := BmSingleton.GetFactoryInstance().ReflectFunctionCall(any, name, args, tmp)
		p.ListHandlers[d.Name] = inc.Interface()
	}
}

func (p *CrawlerPod) CreateEntries() {
	if p.Entries == nil {
		p.Entries = make(map[string]interface{})
	}

	for _, d := range p.conf.Entry {
		any := p.Factory.GetEntriesByName(d.Name)
		name := d.Method
		url := d.Url
		dt := p.ListHandlers[d.Handler]

		inc, _ := BmSingleton.GetFactoryInstance().ReflectFunctionCall(any, name, url, dt)
		p.Entries[d.Title] = inc.Interface()
	}
}
