package controllers

import (
	_ "beeTest/models"
	"github.com/astaxie/beego"
	"path"
	"strings"
)

/*************  INDEX  ******************/
type FileUploadController struct {
	beego.Controller
}

func (c *FileUploadController) Get() {
	c.TplName = "fileupload.html"
}

func (c *FileUploadController) Post() {
	file, information, err := c.GetFile("file")
	if err != nil {
		c.Ctx.WriteString("File retrieval failure")
		return
	} else {
		filename := information.Filename
		picture := strings.Split(filename, ".")
		layout := strings.ToLower(picture[len(picture)-1])
		if layout != "jpg" && layout != "png" && layout != "log" {
			c.Ctx.WriteString("请上传符合格式的图片（png、jpg、gif）")
			return
		}
		savePath := path.Join("/home/crowix/go/src/beeTest/static/upload", filename)
		//fmt.Println(savePath)
		err = c.SaveToFile("file", savePath)
		if err != nil {
			c.Ctx.WriteString("File upload failed！")
		} else {
			c.Ctx.WriteString("File upload succeed!")
		}
	}
	defer file.Close()
	c.TplName = "upload.html"

}
