/*
Copyright © 2023 lxb

*/
package main

import (
	"certificate/cmd"
	"certificate/logger"
	"certificate/setting"
	"fmt"
)

func main() {
	if err := setting.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	if err := logger.InitLogger(); err != nil {
		fmt.Printf("init logger faield：%v\n", err)
		return
	}
	cmd.Execute()

}
