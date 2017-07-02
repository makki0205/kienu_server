package controller

import (
	"github.com/gin-gonic/gin"
	"os"
	"log"
	"io"
	"github.com/makki0205/tmp.fun/model"
	"github.com/cloudfoundry/bytefmt"
	"bytes"
	"strconv"
)

type File struct {
	fileRep *model.FileRepository
}

func NewFileCtr()(File){
	return File{
		fileRep: model.GetFileRepository(),
	}
}
func(self *File) UploadFile(c *gin.Context){
	file, header , err := c.Request.FormFile("file")
	filename := header.Filename
	// fileサイズ取得
	var buff bytes.Buffer
	fileSize, err := buff.ReadFrom(file)
	file.Seek(0,0) //pointerを戻す
	// DBへの保存
	uuid := self.fileRep.SaveFileData(filename, int(fileSize))
	// fileの保存
	// TODO AWS
	dirPath := "./storage/file/"+ uuid
	if _, err := os.Stat(dirPath); err != nil {
		os.MkdirAll(dirPath, 0777)
	}
	filePath := dirPath+"/"+filename
	out, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, gin.H{
		"uuid":uuid,
		"download_url": "/@" + uuid,
		"Description_url": "/api/@"+uuid,
	})
}
func (self * File)GetFile(c *gin.Context){
	uuid := c.Param("uuid")
	file := self.fileRep.GetFileFromUuid(uuid)
	c.Header("Content-Type", "application/force-download")
	c.Header("Content-Length", strconv.Itoa(file.FileSize))
	c.Header("Content-Disposition", "attachment; filename=" + file.FileName)
	c.File("./storage/file/"+uuid+"/"+file.FileName)
}
func(self *File)GetFileDescription(c *gin.Context){
	uuid := c.Param("uuid")
	file := self.fileRep.GetFileFromUuid(uuid)
	if file.ID == 0 {
		c.JSON(404, gin.H{"err":"file not Fund"})
		return
	}
	c.JSON(200, self.createGetFileDescriptionResponse(file))
}
func (self *File)createGetFileDescriptionResponse(file model.File) gin.H{
	return gin.H{
		"download_url":"/file/"+file.Uuid+"/"+file.FileName,
		"file_size": bytefmt.ByteSize(uint64(file.FileSize)),
		"create_at": file.CreatedAt,
		"exp_at": file.Exp,
	}
}



