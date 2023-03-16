package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {

			select {
			case <-ctx.Done():

				fmt.Println("监控退出，停止了...")
				return

			default:

				fmt.Println("goroutine监控中...")

				time.Sleep(2 * time.Second)

			}

		}

	}(ctx)

	time.Sleep(10 * time.Second)

	fmt.Println("可以了，通知监控停止")

	cancel() //为了检测监控过是否停止，如果没有监控输出，就表示停止了

	time.Sleep(5 * time.Second)

}

func main1() {

	ctx, cancel := context.WithCancel(context.Background())

	go watch(ctx, "【监控1】")

	go watch(ctx, "【监控2】")

	go watch(ctx, "【监控3】")

	time.Sleep(10 * time.Second)

	fmt.Println("可以了，通知监控停止")

	cancel() //为了检测监控过是否停止，如果没有监控输出，就表示停止了

	time.Sleep(5 * time.Second)

}

func watch(ctx context.Context, name string) {
	for {

		select {
		case <-ctx.Done():

			fmt.Println(name, "监控退出，停止了...")
			return

		default:

			fmt.Println(name, "goroutine监控中...")

			time.Sleep(2 * time.Second)

		}

	}

}

type favContextKey string

func main2() {

	wg := &sync.WaitGroup{}

	values := []string{"https://www.baidu.com/", "https://www.zhihu.com/"}

	ctx, cancel := context.WithCancel(context.Background())
	for _, url := range values {

		wg.Add(1)

		subCtx := context.WithValue(ctx, favContextKey("url"), url)
		go reqURL(subCtx, wg)
	}

	go func() {

		time.Sleep(time.Second * 3)

		cancel()

	}()

	wg.Wait()

	fmt.Println("exit main goroutine")

}
func reqURL(ctx context.Context, wg *sync.WaitGroup) {

	defer wg.Done()

	url, _ := ctx.Value(favContextKey("url")).(string)
	for {

		select {
		case <-ctx.Done():

			fmt.Printf("stop getting url:%sn", url)
			return

		default:

			r, err := http.Get(url)
			if r.StatusCode == http.StatusOK && err == nil {

				body, _ := ioutil.ReadAll(r.Body)

				subCtx := context.WithValue(ctx, favContextKey("resp"), fmt.Sprintf("%s%x", url, md5.Sum(body)))

				wg.Add(1)
				go showResp(subCtx, wg)
			}

			r.Body.Close()

			//启动子goroutine是为了不阻塞当前goroutine，这里在实际场景中可以去执行其他逻辑，这里为了方便直接sleep一秒

			// doSometing()

			time.Sleep(time.Second * 1)

		}

	}

}

func showResp(ctx context.Context, wg *sync.WaitGroup) {

	defer wg.Done()
	for {

		select {
		case <-ctx.Done():

			fmt.Println("stop showing resp")
			return

		default: //子goroutine里一般会处理一些IO任务，如读写数据库或者rpc调用，这里为了方便直接把数据打印

			fmt.Println("printing ", ctx.Value(favContextKey("resp")))

			time.Sleep(time.Second * 1)

		}

	}

}
