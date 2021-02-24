package main

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

const (
	fontFile = "my.ttf" // 需要使用的字体文件
	fontDPI  = 72       // 屏幕每英寸的分辨率
)

var fontSize float64 = 20 // 字体尺寸

var font *truetype.Font

func main() {
	if !loadFont() {
		fmt.Println("字体加载失败:my.ttf")
		return
	} // 加载字体
	/////////////////////
	if len(os.Args) == 1 {
		fmt.Println("使用方式:")
		fmt.Println("./image.exe 存放图片的文件夹")
		fmt.Println("例如 ./image.exe ./images/")
		return
	}
	whichdir := os.Args[1]
	os.MkdirAll("./imagefinish", 0755)
	fmt.Println("正在打开文件夹", whichdir)
	files, err := ioutil.ReadDir(whichdir)
	if err != nil {
		fmt.Println("没有这个文件夹", whichdir)
		return
	}
	var idx int
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		idx++
		fmt.Println(idx, "处理：", f.Name())
		imageopen := openimg(path.Join(whichdir, f.Name()))
		if imageopen == nil {
			fmt.Println(idx, "非图片：", f.Name())
			continue
		}
		haode := writeText(imageopen, f.Name())
		distpath := path.Join("./imagefinish", f.Name())
		os.Remove(distpath)
		imgw, _ := os.Create(distpath)
		png.Encode(imgw, haode)
		imgw.Close()
		fmt.Println(idx, "完成", f.Name())
	}
	/*
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
	*/
	//生成新图片new.jpg,并设置图片质量
	/*
		imgw, err := os.Create("new.png")
		png.Encode(imgw, m)
		defer imgw.Close()

		fmt.Println("添加水印图片结束请查看")
	*/
}

func openimg(name string) image.Image {
	//图片，网上随便找了一张
	databyte, _ := ioutil.ReadFile(name)
	if len(databyte) == 0 {
		fmt.Println(name, "文件不存在")
		return nil
	}
	jpegdata := make([]byte, len(databyte))
	pngdata := make([]byte, len(databyte))
	for idx := range databyte {
		jpegdata[idx] = databyte[idx]
		pngdata[idx] = databyte[idx]
	}
	jpegreader := bytes.NewReader(jpegdata)
	pngreader := bytes.NewReader(pngdata)
	name = strings.ToLower(name)
	if strings.HasSuffix(name, ".png") || strings.HasSuffix(name, ".jpg") || strings.HasSuffix(name, ".jpeg") {
		img, err := jpeg.Decode(jpegreader)
		if err != nil {
			img, err := png.Decode(pngreader)
			if err != nil {
				fmt.Println(err, name)
				return nil
			}
			return img
		}
		return img
	}
	return nil
}

func loadFont() bool {
	// 读字体数据
	fontBytes, err := ioutil.ReadFile(fontFile)
	if err != nil {
		fmt.Println(err)
		return false
	}
	_font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		fmt.Println(err)
		return false
	}
	font = _font
	return true
}

func writeText(img image.Image, text string) *image.RGBA {
	segs := strings.Split(text, ".")
	if len(segs) == 2 {
		text = segs[0]
	}
	bounds := img.Bounds()
	//根据b画布的大小新建一个新图像
	drawimage := image.NewRGBA(bounds)
	draw.Draw(drawimage, bounds, img, image.ZP, draw.Src)
	width := bounds.Dx()
	height := bounds.Dy()

	c := freetype.NewContext()
	c.SetDPI(fontDPI)
	c.SetFont(font)
	fontSize = float64(height) / 11
	c.SetFontSize(fontSize)
	c.SetClip(drawimage.Bounds())
	c.SetDst(drawimage)
	c.SetSrc(image.White)
	pt := freetype.Pt(width/10, height/10) // 字出现的位置

	_, err := c.DrawString(text, pt)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	{
		// 日期
		c := freetype.NewContext()
		c.SetDPI(fontDPI)
		c.SetFont(font)
		fontSize = float64(height) / 11
		c.SetFontSize(fontSize)
		c.SetClip(drawimage.Bounds())
		c.SetDst(drawimage)
		c.SetSrc(image.White)
		pt := freetype.Pt(width/10, height/5) // 字出现的位置
		c.DrawString(os.Args[2], pt)
	}
	return drawimage
}
