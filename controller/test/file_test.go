package test

import (
	"testing"
	"github.com/HALDevelopersTeam/crow_server/controller"
)

func TestRandString(t *testing.T){
	fileRep := controller.NewFileCtr()
	hoge := fileRep.RandString(4)
	 println(hoge)
}
