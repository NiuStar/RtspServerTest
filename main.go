package main

import (
	"github.com/NiuStar/XRtspServer/stream_server"
	"github.com/NiuStar/server"
	"github.com/NiuStar/server/gin"
	"github.com/NiuStar/XRtspServer/rtsp"
	"net/http"
	"encoding/json"
	"github.com/NiuStar/XRtspServer/RtspClientManager"
	"fmt"
)

var rtspServer *rtsp.RtspServer
func main() {


	
	s := stream_server.NewStreamServer(":8554")

	ser := server.Default()//服务器初始化
	ser.GET("getRtspList",getRtspList)

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

	fmt.Println("list:",list)

	body,_ := json.Marshal(list)

	c.String(http.StatusOK,string(body))


}