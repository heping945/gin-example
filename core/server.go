package core

import (
	"fmt"
	"gin-example/global"
	"gin-example/initialize"
	"net/http"
	"time"
)

func RunServer() {
	if global.GVA_CONFIG.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
	}
	Router := initialize.Routers()
	//Router.Static("/form-generator", "./resource/page")
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)

	s := &http.Server{
		Addr:           address,
		Handler:        Router,
		ReadTimeout:    global.GVA_CONFIG.System.ReadTimeout,  //10s
		WriteTimeout:   global.GVA_CONFIG.System.WriteTimeout, //10s
		MaxHeaderBytes: 1 << 20,
	}
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Debug("server run success on ", address)

	fmt.Printf(`欢迎使用 Go-Example
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
`, s.Addr)
	global.GVA_LOG.Error(s.ListenAndServe())
}
