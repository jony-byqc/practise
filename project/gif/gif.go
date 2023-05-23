package main

import (
	"github.com/CuteReimu/colortools"
	"github.com/CuteReimu/neuquant"
	"image/color"
	"image/gif"
	"image/jpeg"
	"os"
)

func main() {
	// 打开原图
	f, _ := os.Open("D:\\github\\atcive\\practise\\project\\gif\\1.jpg")
	defer func() { _ = f.Close() }()
	img, _ := jpeg.Decode(f)
	// 创建一个gif
	result := &gif.GIF{}
	for j := 0; j < 360; j += 30 {
		// 生成线性渐变的颜色数据
		c := make([]color.Color, 361)
		p := make([]float64, 361)
		for i := 0; i <= 360; i++ {
			c[i] = &colortools.HSV{H: float64(i + j), S: 1.0, V: 0.5}
			p[i] = float64(i) / 360.0
		}
		// 绘制线性渐变
		img1 := colortools.NewLineGradChgColorImage(img.Bounds(), c, p, img.Bounds())
		// 滤色一下
		img2 := colortools.Screen(img1, img)
		// 转成可以用作GIF的图
		img3 := neuquant.Paletted(img2)
		// 插入GIF的一帧
		result.Image = append(result.Image, img3)
		result.Delay = append(result.Delay, 10)
	}
	// 保存
	f2, _ := os.Create("1.gif")
	defer func() { _ = f2.Close() }()
	_ = gif.EncodeAll(f2, result)
}
