package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

func main() {
	//图片，网上随便找了一张
	img_file, err := os.Open("bg.png")
	if err != nil {
		fmt.Println("打开图片出错")
		fmt.Println(err)
		os.Exit(-1)
	}
	defer img_file.Close()
	img, err := jpeg.Decode(img_file)
	if err != nil {
		fmt.Println(err, "bg.png")
		return
	}

	//水印,用的是我自己支付宝的二维码
	wmb_file, err := os.Open("fg.png")
	if err != nil {
		fmt.Println("打开水印图片出错")
		fmt.Println(err)
		os.Exit(-1)
	}
	defer wmb_file.Close()
	wmb_img, err := jpeg.Decode(wmb_file)
	if err != nil {
		fmt.Println(err, "fg.png")
		return
	}

	//把水印写在右下角，并向0坐标偏移10个像素
	offset := image.Pt(img.Bounds().Dx()-wmb_img.Bounds().Dx()-10, img.Bounds().Dy()-wmb_img.Bounds().Dy()-10)
	b := img.Bounds()
	//根据b画布的大小新建一个新图像
	m := image.NewRGBA(b)

	//image.ZP代表Point结构体，目标的源点，即(0,0)
	//draw.Src源图像透过遮罩后，替换掉目标图像
	//draw.Over源图像透过遮罩后，覆盖在目标图像上（类似图层）
	draw.Draw(m, b, img, image.ZP, draw.Src)
	draw.Draw(m, wmb_img.Bounds().Add(offset), wmb_img, image.ZP, draw.Over)

	//生成新图片new.jpg,并设置图片质量
	imgw, err := os.Create("new.png")
	png.Encode(imgw, m)
	defer imgw.Close()

	fmt.Println("添加水印图片结束请查看")
}
