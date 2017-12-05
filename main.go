package main

func main() {
	//gin.SetMode(gin.ReleaseMode) //全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
	r := initRouter()
	r.Run(":7070") // listen and serve on 0.0.0.0:8080
}
