package main

import (
	"StockGetService/models"
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/jasonlvhit/gocron"
	"log"
	"runtime"
	"strings"
	"sync"
	"time"
)

var cpu = runtime.NumCPU()

func main() {
	gocron.Every(1).Monday().Tuesday().Wednesday().Thursday().Friday().Do(StartCrawler)
	<-gocron.Start()
}

func StartCrawler() {
	//取得所有股票代號
	numb, _ := GetStockNumb()
	//計算總數
	stockCount := len(numb)
	//建立StockInfo Model
	models.TableCreate()
	fmt.Printf("爬取總計:%d\n", stockCount)
	//使用WaitGroup等待工作完成
	var wg sync.WaitGroup

	wg.Add(cpu)
	for i := 0; i < cpu; i++ {
		from := i * stockCount / cpu
		to := (i + 1) * stockCount / cpu
		if to == stockCount {
			to = stockCount + 1
		}
		//Goroutine 建立 cpu核心數個goroutine
		go func() {
			fmt.Printf("Goroutine:%d\n", i)
			for j := from; j <= to; j++ {
				//取得爬取資料
				stockDetail := GetStockDetails(numb[j])
				if stockDetail != nil {
					//存進Db
					models.StockInfoSave(stockDetail)
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func GetStockDetails(stockNumb string) (stockDetails []string) {
	/*
	 1、資料來源url:台灣證交所
	 2、根據html tag、class尋找所需資料
	 3、設定等待爬取時間 cpu核心 * 5 秒
	 4、字串處理
	 5、返回處理好的資料
	*/
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	var res string
	var stockName string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://mis.twse.com.tw/stock/fibest.jsp?stock=`+stockNumb),
		chromedp.Sleep(time.Duration(cpu)*5*time.Second),
		chromedp.Text(`label[class=title]`, &stockName, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.Text(`#fibestrow`, &res, chromedp.NodeVisible, chromedp.ByQuery),
	)
	if err != nil {
		log.Print(err)
	}
	stockInfo := strings.Fields(res)
	layout := `2006/01/02`
	now, _ := time.Parse(layout, time.Now().Format(`2006/01/02`))
	rnf := ""
	sName1 := strings.Replace(stockName, "[", "", -1)
	sName := strings.Split(strings.Replace(sName1, "]", "", -1), " ")
	if stockInfo[0] == "-" {
	} else {
		if strings.Contains(stockInfo[1], "▲") {
			rnf = strings.ReplaceAll(stockInfo[1], "▲", "")
		} else {
			rnf = strings.ReplaceAll(stockInfo[1], "▼", "")
		}
		fmt.Print(sName[1] + "\n")
		stockDetails = []string{sName[0], sName[1], sName[2], stockInfo[0], rnf, stockInfo[2], stockInfo[3], stockInfo[4], stockInfo[7], stockInfo[8], stockInfo[9], stockInfo[10], stockInfo[11], stockInfo[12], stockInfo[13], strings.Split(now.String(), " ")[0]}
	}
	return stockDetails
}

func GetStockNumb() (stockNumb []string, stockName []string) {
	/*
	 1、資料來源url:撿股讚
	 2、取回所有股號
	 3、回傳股號
	*/
	var nodes []*cdp.Node
	var names []*cdp.Node
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://stock.wespai.com/p/3752`),
		chromedp.Nodes(`#example > tbody > tr > td[class=sorting_1]`, &nodes, chromedp.ByQueryAll),
		chromedp.Nodes(`#example > tbody > tr > td > a`, &names, chromedp.ByQueryAll),
	)
	if err != nil {
		log.Fatal(err)
	}
	for _, n := range nodes {
		exceptDq := strings.Replace(n.Dump("", " ", false), `"`, "", -1)
		exceptTd := strings.Replace(exceptDq, `td`, "", -1)
		exceptText := strings.Replace(exceptTd, `#text`, "", -1)
		exceptClass := strings.TrimSpace(strings.Replace(exceptText, `[class=sorting_1]`, "", -1))
		stockNumb = append(stockNumb, exceptClass)
	}

	for _, n := range names {
		exceptDq := strings.Replace(n.Dump("", " ", false), `"`, "", -1)
		exceptA := strings.Replace(exceptDq, `a`, ``, -1)
		exceptText := strings.Replace(exceptA, `#text`, "", -1)
		exceptN := strings.Replace(exceptText, "\n", " ", -1)
		name := strings.Split(strings.Split(exceptN, "[")[1], "]   ")[1]
		stockName = append(stockName, name)
	}
	return stockNumb, stockName
}
