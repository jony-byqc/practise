package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"unicode"
)

func main() {
	// base url
	baseUrl, err := url.Parse("http://www.mywebsite.com")
	if err != nil {
		fmt.Println("Malformed URL: ", err.Error())
		return
	}

	// 添加query string参数
	baseUrl.Path += "path with?reserved characters"

	// 参数准备
	params := url.Values{}
	params.Add("q", "Hello World")
	params.Add("u", "@rajeev")

	// 附加query string到url上
	baseUrl.RawQuery = params.Encode()

	fmt.Printf("Encoded URL is %q\n", baseUrl.String())

	uncodeStr := StrToUnicode("十大户￥@！#%……&……*（）——+《》、，。、；‘、配【】")
	fmt.Println(uncodeStr)
	main5()
}

func main1() {
	path := "path with?reserved+characters"
	fmt.Println(url.PathEscape(path))
}

func main2() {
	params := url.Values{}
	params.Add("name", "@Rajeev")
	params.Add("phone", "+919999999999")

	fmt.Println(params.Encode())
}

func main3() {
	query := "Hellö Wörld@Golang"
	fmt.Println(url.QueryEscape(query))
}
func StrToUnicode(str string) string {
	DD := []rune(str) //需要分割的字符串内容，将它转为字符，然后取长度。
	finallStr := ""
	for i := 0; i < len(DD); i++ {
		if unicode.Is(unicode.Scripts["Han"], DD[i]) {
			textQuoted := strconv.QuoteToASCII(string(DD[i]))
			finallStr += textQuoted[1 : len(textQuoted)-1]
		} else {
			h := fmt.Sprintf("%x", DD[i])
			finallStr += "\\u" + isFullFour(h)
		}
	}
	return finallStr
}

func isFullFour(str string) string {
	if len(str) == 1 {
		str = "000" + str
	} else if len(str) == 2 {
		str = "00" + str
	} else if len(str) == 3 {
		str = "0" + str
	}
	return str
}

func main5() {
	rs := []rune("golang中文unicode编码")
	json := ""
	html := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			json += string(r)
			html += string(r)
		} else {
			json += "\\u" + strconv.FormatInt(int64(rint), 16) // json
			html += "&#" + strconv.Itoa(int(r)) + ";"          // 网页
		}
	}
	fmt.Printf("JSON: %s\n", json)
	fmt.Printf("HTML: %s\n", html)
}

func main6() {
	rs := []rune("golang中文unicode编码")
	json := ""
	html := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			json += string(r)
			html += string(r)
		} else {
			json += "\\u" + strconv.FormatInt(int64(rint), 16) // json
			html += "&#" + strconv.Itoa(int(r)) + ";"          // 网页
		}
	}
	fmt.Printf("JSON: %s\n", json)
	fmt.Printf("HTML: %s\n", html)
}

func main7() {
	input := []byte("hello golang base64 快乐编程http://www.01happy.com +~")

	// 演示base64编码
	encodeString := base64.StdEncoding.EncodeToString(input)
	fmt.Println(encodeString)

	// 对上面的编码结果进行base64解码
	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(decodeBytes))

	fmt.Println()

	// 如果要用在url中，需要使用URLEncoding
	uEnc := base64.URLEncoding.EncodeToString([]byte(input))
	fmt.Println(uEnc)

	uDec, err := base64.URLEncoding.DecodeString(uEnc)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(uDec))
}
