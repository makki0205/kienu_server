package model

import (
	"github.com/jinzhu/gorm"
	"time"
	"math/rand"
)

type File struct {
	gorm.Model
	Uuid string
	FileName string
	FileSize int
	Exp time.Time
}

type FileRepository struct {}
var fileRep *FileRepository

func GetFileRepository()(*FileRepository){
	if fileRep == nil {
		fileRep = &FileRepository{}
	}
	return fileRep
}

func (self *FileRepository)Migrate(){
	db.AutoMigrate(File{})
}
func (self *FileRepository)SaveFileData(fileName string, fileSize int)string{
	var uuid string
	for {
		uuid = self.RandString(4)
		if ! self.ExistUuid(uuid){
			break
		}
	}
	file := File{
		Uuid: uuid,
		FileName: fileName,
		FileSize:fileSize,
		Exp:time.Now().AddDate(0,0,7),
	}
	db.Create(&file)
	return file.Uuid
}
func (self *FileRepository)GetFileFromUuid(uuid string)(File){
	var file File
	db.Where("uuid = ?", uuid).First(&file)
	return file
}
// DBにuuidがあればtrueを返す
func (self *FileRepository)ExistUuid(uuid string)bool{
	var file File
	db.Where("uuid = ?",uuid).First(&file)
	println(file.Uuid)
	return file.ID != 0
}


// random文字列生成
const (
	rs6Letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	IdxBits = 6
	IdxMask = 1<<IdxBits - 1
	IdxMax = 63 / IdxBits
)
var randSrc = rand.NewSource(time.Now().UnixNano())
func(self *FileRepository) RandString(n int) string {
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