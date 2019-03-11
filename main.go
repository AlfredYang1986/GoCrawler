package main

import (
	"fmt"
	"github.com/alfredyang1986/GoCrawler/Tables"
	"github.com/alfredyang1986/GoCrawler/Config"
	"github.com/alfredyang1986/GoCrawler/Entry"
)

func main() {
	fmt.Println("alfred")
	fmt.Println("pod archi begins")

	fac := Tables.CrawlerTable{}
	var pod = Config.CrawlerPod{ Name: "alfred test", Factory:fac }
	pod.RegisterSerFromYAML("Resources/GoCrawlerConfig.yaml")

	for _, item := range pod.Entries {
		fmt.Println(item)
		item.(*Entry.BmEntry).StartCrawler()
	}

	fmt.Println("pod archi ends")
}