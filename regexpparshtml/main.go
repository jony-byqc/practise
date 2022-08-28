package main

import (
	"github.com/NiuStar/log/fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	//ioutil.ReadAll()
	text := "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n    <title>Document</title>\n</head>\n<body>\n    <h1>H1(heading)标签</h1>\n    <h2>H2(heading)标签</h2>\n    <h3>H3(heading)标签</h3>\n    <h4>H4(heading)标签</h4>\n    <h5>H5(heading)标签</h5>\n    <h1 align=\"center\">H(heading)标签</h1>\n    <!-- <h1> - <h6> 的 \"align\" 属性，在html4中不推荐使用，而在HTML5中，已经不被支持了。所以，我们需要设置H标签的对齐方式、颜色等细节，推荐使用CSS样式！ -->\n     \n    <p>昨日，北京、天津、河北、山西、陕西中北部、甘肃中东部、内蒙古中部及吉林东部、辽宁中西部等地出现小到中雪，北京南部、天津北部、河北廊坊、山西忻州等局地大雪(5～7毫米)。今5时，北京西北部、河北西北部、山西北部、陕西北部、甘肃中部、青海东部等地积雪深度有2～6厘米。\n    </p>\n\n     \n    <b>加粗文字</b>\n\n\n    <b>加粗文字bold 非语义标签</b> \n    <strong>加粗文字 语义化标签   强调的  希望被爬虫抓取</strong> \n    \n    <i>斜体文字 italic</i> \n    <em>斜体文字 语义化标签</em>  \n    \n    <s>中划线</s>\n    <del>中划线 语义化标签</del>\n    \n    <u>下划线</u>\n    <ins>下划线 语义化标签</ins>\n    \n    <br> \n    <br> \n    <br> \n    <!-- break  换行 -->\n    \n    <hr color=\"red\" align=\"\"  width=\"300px\">\n    <!-- align 水平对齐方式 left 左对齐  right  center居中 -->\n   \n</body>\n</html>\n————————————————\n版权声明：本文为CSDN博主「qq_45136679」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。\n原文链接：https://blog.csdn.net/qq_45136679/article/details/109920829"
	RepImages(text)
}

//goquery包在github上的地址：https://github.com/PuerkitoBio/goquery
//解析html文件里面的标签，获取标签里面的内容
func PageQuaryTest(baseUrl string) {
	var res *http.Response

	//ioutil.ReadAll()
	//RepImages(resp)
	doc, err := goquery.NewDocumentFromResponse(res)
	if err == nil {
		doc.Find("img").Each(func(i int, s *goquery.Selection) {
			//解析<div>标签
			//h,err := s.Html()
			v, t := s.Attr("src")
			fmt.Println("v--->", v, "   t--->", t)
			//fmt.Println(i, s.Text())
			s.SetAttr("src", "") //修改标签的内容
		})
	} else {
		fmt.Println("err--->", err)
	}
	doc.Find("div").Each(func(i int, s *goquery.Selection) {
		//解析<div>标签
		fmt.Println(i, s.Text())
	})
	doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("a").Text()
		title := s.Find("i").Text()
		fmt.Println("Review %d: %s - %s\n", i, band, title)
	})
}
func RepImages(htmls string) {
	var imgRE = regexp.MustCompile(`src=[^>]+\b[^>']`)
	imgs := imgRE.FindAllStringSubmatch(htmls, -1)
	out := make([]string, len(imgs))
	for i := range out {
		out[i] = imgs[i][0]

		a := imgRE.ReplaceAllString(imgs[i][0], "asd")
		fmt.Println(a)
	}
	fmt.Println(imgs)

	re := regexp.MustCompile(`<h2>(.*?)</h2>.*?<em>(.*?)</em>`)
	res := re.FindAllStringSubmatch(htmls, -1) // -1 表示匹配次数(全部匹配)
	fmt.Println(res)
	// 输出

	// [[<h2>h2内部数据h2111111</h2>asd<em>77ememe111</em> h2内部数据h2111111 77ememe111] [<h2>h2内部数据h2222222</h2>asd<em>77ememe2222</em> h2内部数据h2222222 77ememe2222] [<h2>h2内部数据h23333</h2>asd<em>77emem3333</em> h2内部数据h23333 77emem3333]] }go FindAllStringSubmatch 返回的是数组嵌套数据的数据，而且内部数组第一个是匹配的整体字符串
	if res == nil {
		return
	}
	joinRes := strings.Join(res[0], ",")
	fmt.Println(joinRes)
	fmt.Println(strings.Replace(joinRes, "内部数据", "12354324324", 0)) //n指定要在字符串中替换的字符数。如果n小于0，则替换次数数没有限制

}
