package main

import (
	"github.com/labstack/gommon/log"
	"go-test/pdf/dbops"
	"fmt"
)

func main() {
	downloadData, err := dbops.GetMemberInfo() // 获取需要导出数据
	if err != nil {
		log.Printf("%s", err)
		return
	}
	// 导出数据
	err = dbops.WritePdf("./pdf/data/斗罗大陆人员名单.pdf", downloadData)
	if err != nil {
		log.Printf("%s", err)
		return
	}
	fmt.Println("执行完毕")
}
