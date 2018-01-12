package main

import (
	"nqc.cn/XRtspServer/stream_server"
	"nqc.cn/server"
	"nqc.cn/server/gin"
	"nqc.cn/XRtspServer/rtsp"
	"net/http"
	"encoding/json"
	"nqc.cn/XRtspServer/RtspClientManager"
)

var rtspServer *rtsp.RtspServer
func main() {



	s := stream_server.NewStreamServer(":8554")

	ser := server.Default()//服务器初始化
	ser.GET("getRtspList",getRtspList)
	ser.GET("getUserList",getUserList)


	s.Wrap(ser.RunServer)
	

	var err error
	rtspServer, err = s.NewRtspServer()
	if err != nil {
		return
	}


	s.Run(rtspServer)



}

func getRtspList(c *gin.Context) {

	list := RtspClientManager.GetCurrManagers()

	body,_ := json.Marshal(list)

	c.String(http.StatusOK,string(body))


}

func getUserList(c *gin.Context) {

	//list := rtspServer.GetUserList()

	//body,_ := json.Marshal(list)

	c.String(http.StatusOK,"getUserList")

}