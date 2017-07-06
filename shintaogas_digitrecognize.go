// Copyright 2011 <chaishushan@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	// "fmt"
	// "os"
	"path"
	"runtime"
	"github.com/lazywei/go-opencv/opencv"
	//"../opencv" // can be used in forks, comment in real application
)

func main() {
	_, currentfile, _, _ := runtime.Caller(0)
	filename := path.Join(path.Dir(currentfile), "/images/sample.bmp")

	image := opencv.LoadImage(filename)
	if image == nil {
		panic("LoadImage fail")
	}
	defer image.Release()
	imageGray := opencv.CreateImage(160, 120, opencv.IPL_DEPTH_8U, 1)
	w := 17
	opencv.CvtColor(image, imageGray, opencv.CV_RGB2GRAY)

	cropAll := opencv.Crop(imageGray, 15, 50, 70, 19 )
	opencv.Threshold(cropAll, cropAll, 165, 255, opencv.CV_THRESH_BINARY)
	num1 := opencv.Crop(imageGray, 15, 50, w, 19 )
	opencv.Threshold(num1, num1, 165, 255, opencv.CV_THRESH_BINARY)
	num2 := opencv.Crop(imageGray, 15 + w * 1, 50, w, 19 )
	opencv.Threshold(num2, num2, 165, 255, opencv.CV_THRESH_BINARY)
	num3 := opencv.Crop(imageGray, 15 + w * 2, 50, w, 19 )
	opencv.Threshold(num3, num3, 165, 255, opencv.CV_THRESH_BINARY)
	num4 := opencv.Crop(imageGray, 15 + w * 3, 50, w, 19 )
	opencv.Threshold(num4, num4, 165, 255, opencv.CV_THRESH_BINARY)

	wImage := opencv.NewWindow("Shintaogas digit recognize")
	defer wImage.Destroy()
	wImage.Move(0, 0)
	wImage.ShowImage(image)

	wGrat := opencv.NewWindow("gray")
	defer wGrat.Destroy()
	wGrat.Move(160, 0)
	wGrat.ShowImage(imageGray)

	wCropAll := opencv.NewWindow("cropAll")
	defer wCropAll.Destroy()
	wCropAll.Move(320+50, 0)
	wCropAll.ShowImage(cropAll)

	wNum1 := opencv.NewWindow("num1")
	defer wNum1.Destroy()
	wNum1.Move(370 + w, 120)
	wNum1.ShowImage(num1)

	wNum2 := opencv.NewWindow("num2")
	defer wNum2.Destroy()
	wNum2.Move(370 + w * 2, 240)
	wNum2.ShowImage(num2)

	wNum3 := opencv.NewWindow("num3")
	defer wNum3.Destroy()
	wNum3.Move(370 + w * 3, 360)
	wNum3.ShowImage(num3)

	wNum4 := opencv.NewWindow("num4")
	defer wNum4.Destroy()
	wNum4.Move(370 + w * 4, 480)
	wNum4.ShowImage(num4)

	opencv.WaitKey(0)
}
