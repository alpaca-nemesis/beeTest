package controllers

import (
	_ "beeTest/models"
	"github.com/astaxie/beego"
	"log"
	"path"
	"strings"
)

/*************  INDEX  ******************/
type FileUploadController struct {
	beego.Controller
}

var filePrefix = beego.AppConfig.String("uploaddir")

func (c *FileUploadController) Get() {
	c.TplName = "fileupload.html"
}

func (c *FileUploadController) Post() {
	file, information, err := c.GetFile("file")
	defer file.Close()
	if err != nil {
		c.Data["message"] = err
		c.TplName = "message.html"
		log.Println("client err: ", err)
		return
	} else {
		filename := information.Filename
		picture := strings.Split(filename, ".")
		layout := strings.ToLower(picture[len(picture)-1])
		if layout != "jpg" && layout != "png" && layout != "log" {
			c.Data["message"] = "请上传符合格式的图片（png、jpg、gif）"
			c.TplName = "message.html"
			log.Println("client err: ", err)
			return
		}
		savePath := path.Join(filePrefix, filename)
		//fmt.Println(savePath)
		err = c.SaveToFile("file", savePath)
		if err != nil {
			c.Ctx.WriteString("File upload failed！")
		} else {
			c.Ctx.WriteString("File upload succeed!")
		}
	}
	c.TplName = "upload.html"

}
