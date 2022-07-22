package main

import (
	"encoding/json"
	"fmt"
	"gin/Model"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/golang/glog"

	"context"
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

var (
	topics = "heima"
)

type Colly struct {
	Id           string `json:"id" gorm:"primaryKey"`
	Uri          string
	Download_uri string
	Web_id       string
	Time         string
	Name         string
	Artist       string
}

func (Colly) TableName() string {
	return "ipfs_colly"
}

// kafka生产者
func ProducerFunc(start int, end int) {
	fmt.Println("ProducerFunc start")
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	producer, err := sarama.NewSyncProducer([]string{"192.168.71.130:9092"}, config)

	if err != nil {
		panic(err)
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic:     topics,
		Partition: int32(-1),
		Key:       sarama.StringEncoder("key"),
	}

	var value string
	for i := start; i < end; i++ {
		uri := "https://www.kuwo.cn/mvplay/" + strconv.Itoa(i)
		resp, err := http.Get(uri)
		if err != nil {
			fmt.Println("http.Get err=", err)
			continue
		}
		// defer resp.Body.Close() // 函数结束时关闭Body
		resp.Body.Close()
		if resp.StatusCode != 200 {
			fmt.Printf("id: %d http.Get StatusCode=%d \n", i, resp.StatusCode)
			continue
		}

		example, err7 := GetHttpHtmlContent(uri, ".mv_out", "document.querySelector('body')")
		if err7 != nil {
			fmt.Println("GetHttpHtmlContent err:", err7)
			continue
		}
		dom, err4 := goquery.NewDocumentFromReader(strings.NewReader(example))
		if err4 != nil {
			fmt.Println("goquery.NewDocumentFromReader err:", err4)
			continue
		}

		src, bool := dom.Find("#vjs_video_403_html5_api").Eq(0).Attr("src")
		if !bool || src == "res not found" {
			fmt.Printf("dom.Find -> id:%d empty \n", i)
			continue
		}
		name := dom.Find(".mv_name").Text()
		artist := dom.Find(".artist").Text()

		// 判断是否存在，存在就不插入数据
		ret := &Colly{}
		_, count := Model.GetList(ret, "web_id = "+strconv.Itoa(i), "", "1")
		if count > 0 {
			fmt.Printf("id:%s is exist \n", msg.Value)
			continue
		}

		ret.Web_id = strconv.Itoa(i)
		ret.Uri = uri
		ret.Download_uri = src
		ret.Name = name
		ret.Artist = artist
		ret.Time = time.Now().Format("2006/01/02 15:04:05")
		json, _ := json.Marshal(ret)
		value = strings.Replace(string(json), "\n", "", -1)
		msg.Value = sarama.ByteEncoder(value)
		paritition, offset, err := producer.SendMessage(msg)

		if err != nil {
			fmt.Println("Send Message Fail")
		}

		fmt.Printf("Partion = %d, offset = %d\n", paritition, offset)
	}

}

// kafka消费者
func ConsumerFunc() {
	fmt.Println("ConsumerFunc start")
	groupID := "test-consumersync-group"
	config := cluster.NewConfig()
	config.Group.Return.Notifications = true
	// config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.CommitInterval = 3 * time.Second
	config.Consumer.Offsets.Initial = sarama.OffsetNewest //初始从最新的offset开始

	cc, err9 := cluster.NewConsumer(strings.Split("192.168.71.130:9092", ","), groupID, strings.Split(topics, ","), config)
	if err9 != nil {
		glog.Errorf("Failed open consumersync: %v", err9)
		return
	}
	defer cc.Close()
	go func(cc *cluster.Consumer) {
		errors := cc.Errors()
		noti := cc.Notifications()
		for {
			select {
			case err8 := <-errors:
				glog.Errorln(err8)
			case <-noti:
			}
		}
	}(cc)

	for msg := range cc.Messages() {
		// fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s", msg.Topic, msg.Partition, msg.Offset, msg.Value)
		ret := &Colly{}
		json.Unmarshal(msg.Value, ret)

		fmt.Printf("ret: %+v", ret)
		err2 := Model.Add(ret) // 此处只是用于写入数据库，可以做自己需要的操作
		if err2 != nil {
			fmt.Println(err2)
		}

		cc.MarkOffset(msg, "") // 消费者处理完毕一条消息后，提交告知kafka
	}
}

//得到具体的数据
func GetSpecialData(htmlContent string, selector string) (string, error) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		return "", err
	}

	var str string
	dom.Find(selector).Each(func(i int, selection *goquery.Selection) {
		str = selection.Text()
	})
	return str, nil
}

//获取网站上爬取的数据
func GetHttpHtmlContent(url string, selector string, sel interface{}) (string, error) {
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true), // debug使用
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}
	//初始化参数，先传一个空的数据
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	c, _ := chromedp.NewExecAllocator(context.Background(), options...)

	// create context
	chromeCtx, cancels := chromedp.NewContext(c, chromedp.WithLogf(log.Printf))
	defer cancels() // 不可少，否则一直开后台chrome进程，狂吃内存
	// 执行一个空task, 用提前创建Chrome实例
	chromedp.Run(chromeCtx, make([]chromedp.Action, 0, 1)...)

	//创建一个上下文，超时时间为40s
	timeoutCtx, cancel := context.WithTimeout(chromeCtx, 40*time.Second)
	defer cancel() // 不可少，否则一直开后台chrome进程，狂吃内存

	var htmlContent string
	err := chromedp.Run(timeoutCtx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(selector),
		chromedp.OuterHTML(sel, &htmlContent, chromedp.ByJSPath),
	)
	if err != nil {
		return "", err
	}
	//log.Println(htmlContent)

	return htmlContent, nil
}
