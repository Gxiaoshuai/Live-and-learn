package util

import (
	"image"
	"image/draw"
	"image/jpeg"
	"os"
)
type ImageUtil struct{
	Error string
}

func (this *ImageUtil)RepeatImage(filePath string,savePath string,repeatTimesX int,repeatTimesY int)(bool){
	var path = new (Path)
	if(!path.IsExist(filePath) || !path.IsFile(filePath) ){
		this.Error = "图片文件不存在"
		return false
	}
	if(path.IsExist(savePath)){
		this.Error = "目标文件已经存在"
		return false
	}
	file, err := os.Open(filePath)
	if err != nil {
		this.Error = err.Error()
		return false
	}
	defer file.Close()
	newFile,err := os.Create(savePath)
	if err != nil{
		this.Error = err.Error()
		return false
	}
	defer newFile.Close()
	sourceImage,_ := jpeg.Decode(file)

	imageRe := sourceImage.Bounds()
	var x = imageRe.Max.X
	var y = imageRe.Max.Y
	if repeatTimesX >0{
		x *= repeatTimesX
	}
	if repeatTimesY > 0{
		y *= repeatTimesY
	}
	var newImage = image.NewNRGBA(image.Rect(0,0,x,y))
	var i = 0;
	for i < repeatTimesX{
		var j = 0
		for j < repeatTimesY{
			draw.Draw(newImage, imageRe.Add(image.Point{imageRe.Max.X * i,imageRe.Max.Y  *  j}), sourceImage, image.Point{0,0}, draw.Src)
			j++
		}
		i++
	}
	draw.Draw(newImage, imageRe, sourceImage, image.Point{0,0}, draw.Src)
	err = jpeg.Encode(newFile,newImage,nil)
	if err ==nil{
		return true
	}
	return false
}


