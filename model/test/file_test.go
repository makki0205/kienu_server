package test

import (
	"testing"
	"github.com/HALDevelopersTeam/crow_server/model"
)


func TestFileModelMigrate(t *testing.T) {
	urep := model.GetFileRepository()
	urep.Migrate()
}

func Testファイルモデル(t *testing.T) {
	urep := model.GetFileRepository()
	urep.SaveFileData("hoge.txt", 1256765)
}
