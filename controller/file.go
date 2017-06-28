package controller

import (
	"github.com/gin-gonic/gin"
	"os"
	"log"
	"io"
	"time"
	"math/rand"
)

type File struct {}

func NewFileCtr()(File){
	return File{}
}
func(self *File) UploadFile(c *gin.Context){
	// TODO DBで永続化
	file, header , err := c.Request.FormFile("file")
	filename := header.Filename
	uuid := self.RandString(4)

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
		"download_url": "/file/"+uuid+"/"+filename,
		"Description_url": "/api/@"+uuid,
	})
}
func(self *File)GetFileDescription(c *gin.Context){
	uuid := c.Param("uuid")
	c.JSON(200, uuid)
}

var randSrc = rand.NewSource(time.Now().UnixNano())

const (
	rs6Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	IdxBits = 6
	IdxMask = 1<<IdxBits - 1
	IdxMax = 63 / IdxBits
)

func(self *File) RandString(n int) string {
	b := make([]byte, n)
	cache, remain := randSrc.Int63(), IdxMax
	for i := n-1; i >= 0; {
		if remain == 0 {
			cache, remain = randSrc.Int63(), IdxMax
		}
		idx := int(cache & IdxMask)
		if idx < len(rs6Letters) {
			b[i] = rs6Letters[idx]
			i--
		}
		cache >>= IdxBits
		remain--
	}
	return string(b)
}