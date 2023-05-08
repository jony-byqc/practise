package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"gitee.com/baixudong/gospider/re"
	"gitee.com/baixudong/gospider/requests"
	"gitee.com/baixudong/gospider/tools"
	"gitee.com/baixudong/gospider/websocket"
	"github.com/google/uuid"
)

func createMsg(path, requestId, contentType string, con string) []byte {
	txt := fmt.Sprintf("Path: %s\r\nX-RequestId: %s\r\nX-Timestamp: %s\r\nContent-Type: %s\r\n\r\n%s",
		path,
		requestId,
		time.Now().Format("2006-01-02T15:04:05.271Z"),
		contentType,
		con,
	)
	return tools.StringToBytes(txt)
}

var quid = strings.ToUpper(re.Sub("-", "", uuid.New().String()))

func main() {
	socketUrl := "wss://eastus.api.speech.microsoft.com/cognitiveservices/websocket/v1?TrafficType=AzureDemo&Authorization=bearer%20undefined&X-ConnectionId=" + strings.ToUpper(re.Sub("-", "", uuid.New().String()))
	reqCli, err := requests.NewClient(nil)
	if err != nil {
		log.Panic(err)
	}
	resp, err := reqCli.Request(nil, "get", socketUrl, requests.RequestOption{
		Headers: map[string]string{
			"Accept-Encoding": "gzip, deflate, br",
			"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
			"Cache-Control":   "no-cache",
			"Host":            "eastus.api.speech.microsoft.com",
			"Origin":          "https://azure.microsoft.com",
			"Pragma":          "no-cache",
			"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36 Edg/107.0.1418.26",
		},
	})
	if err != nil {
		log.Panic(err)
	}
	wsCli := resp.WebSocket()
	if err = wsCli.Send(nil,
		websocket.MessageText,
		createMsg("speech.config", quid, "application/json",
			tools.Any2json(map[string]any{
				"context": map[string]any{
					"system": map[string]string{"name": "SpeechSDK", "version": "1.19.0", "buildtags": "JavaScript", "lang": "JavaScript"},
					"os": map[string]string{"platform": "Browser/Win32", "name": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36 Edg/107.0.1418.26",
						"version": "5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36 Edg/107.0.1418.26"},
				},
			}).Raw),
	); err != nil {
		log.Panic(err)
	}
	if err = wsCli.Send(nil, websocket.MessageText, createMsg("synthesis.context", quid, "application/json",
		tools.Any2json(map[string]any{
			"synthesis": map[string]any{
				"audio": map[string]any{
					"metadataOptions": map[string]bool{"bookmarkEnabled": false, "sentenceBoundaryEnabled": false, "visemeEnabled": false, "wordBoundaryEnabled": false},
					"outputFormat":    "audio-24khz-96kbitrate-mono-mp3",
				},
				"language": map[string]bool{"autoDetection": false},
			},
		}).Raw,
	)); err != nil {
		log.Panic(err)
	}
	if err = wsCli.Send(nil, websocket.MessageText, createMsg("ssml", quid, "application/ssml+xml",
		`<speak xmlns="http://www.w3.org/2001/10/synthesis" xmlns:mstts="http://www.w3.org/2001/mstts" xmlns:emo="http://www.w3.org/2009/10/emotionml" version="1.0" xml:lang="en-US"><voice name="zh-CN-XiaoxiaoNeural"><prosody rate="0%" pitch="0%">     
        招标125456请尽情使用招标125456
        </prosody></voice></speak>`,
	)); err != nil {
		log.Panic(err)
	}
	for {
		msgType, msgCon, err := wsCli.Recv(nil)
		if err != nil {
			log.Panic(err)
		}
		switch msgType {
		case websocket.MessageText:
			log.Print(string(msgCon))
			if strings.Contains(string(msgCon), "Path:turn.end") {
				log.Print("转换结束")
				break
			}
		case websocket.MessageBinary: //音频流
			log.Print("读取音频流中")
		}
	}
}
