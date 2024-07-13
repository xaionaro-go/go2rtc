package main

import (
	"github.com/xaionaro-go/go2rtc/internalpkg/api"
	"github.com/xaionaro-go/go2rtc/internalpkg/api/ws"
	"github.com/xaionaro-go/go2rtc/internalpkg/app"
	"github.com/xaionaro-go/go2rtc/internalpkg/bubble"
	"github.com/xaionaro-go/go2rtc/internalpkg/debug"
	"github.com/xaionaro-go/go2rtc/internalpkg/dvrip"
	"github.com/xaionaro-go/go2rtc/internalpkg/echo"
	"github.com/xaionaro-go/go2rtc/internalpkg/exec"
	"github.com/xaionaro-go/go2rtc/internalpkg/expr"
	"github.com/xaionaro-go/go2rtc/internalpkg/ffmpeg"
	"github.com/xaionaro-go/go2rtc/internalpkg/gopro"
	"github.com/xaionaro-go/go2rtc/internalpkg/hass"
	"github.com/xaionaro-go/go2rtc/internalpkg/hls"
	"github.com/xaionaro-go/go2rtc/internalpkg/homekit"
	"github.com/xaionaro-go/go2rtc/internalpkg/http"
	"github.com/xaionaro-go/go2rtc/internalpkg/isapi"
	"github.com/xaionaro-go/go2rtc/internalpkg/ivideon"
	"github.com/xaionaro-go/go2rtc/internalpkg/mjpeg"
	"github.com/xaionaro-go/go2rtc/internalpkg/mp4"
	"github.com/xaionaro-go/go2rtc/internalpkg/mpegts"
	"github.com/xaionaro-go/go2rtc/internalpkg/nest"
	"github.com/xaionaro-go/go2rtc/internalpkg/ngrok"
	"github.com/xaionaro-go/go2rtc/internalpkg/onvif"
	"github.com/xaionaro-go/go2rtc/internalpkg/roborock"
	"github.com/xaionaro-go/go2rtc/internalpkg/rtmp"
	"github.com/xaionaro-go/go2rtc/internalpkg/rtsp"
	"github.com/xaionaro-go/go2rtc/internalpkg/srtp"
	"github.com/xaionaro-go/go2rtc/internalpkg/streams"
	"github.com/xaionaro-go/go2rtc/internalpkg/tapo"
	"github.com/xaionaro-go/go2rtc/internalpkg/webrtc"
	"github.com/xaionaro-go/go2rtc/internalpkg/webtorrent"
	"github.com/xaionaro-go/go2rtc/pkg/shell"
)

func main() {
	app.Version = "1.9.4"

	// 1. Core modules: app, api/ws, streams

	app.Init() // init config and logs

	api.Init() // init API before all others
	ws.Init()  // init WS API endpoint

	streams.Init() // streams module

	// 2. Main sources and servers

	rtsp.Init()   // rtsp source, RTSP server
	webrtc.Init() // webrtc source, WebRTC server

	// 3. Main API

	mp4.Init()   // MP4 API
	hls.Init()   // HLS API
	mjpeg.Init() // MJPEG API

	// 4. Other sources and servers

	hass.Init()       // hass source, Hass API server
	onvif.Init()      // onvif source, ONVIF API server
	webtorrent.Init() // webtorrent source, WebTorrent module

	// 5. Other sources

	rtmp.Init()     // rtmp source
	exec.Init()     // exec source
	ffmpeg.Init()   // ffmpeg source
	echo.Init()     // echo source
	ivideon.Init()  // ivideon source
	http.Init()     // http/tcp source
	dvrip.Init()    // dvrip source
	tapo.Init()     // tapo source
	isapi.Init()    // isapi source
	mpegts.Init()   // mpegts passive source
	roborock.Init() // roborock source
	homekit.Init()  // homekit source
	nest.Init()     // nest source
	bubble.Init()   // bubble source
	expr.Init()     // expr source
	gopro.Init()    // gopro source

	// 6. Helper modules

	ngrok.Init() // ngrok module
	srtp.Init()  // SRTP server
	debug.Init() // debug API

	// 7. Go

	shell.RunUntilSignal()
}
