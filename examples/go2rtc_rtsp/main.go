package main

import (
	"github.com/xaionaro-go/go2rtc/internalpkg/app"
	"github.com/xaionaro-go/go2rtc/internalpkg/rtsp"
	"github.com/xaionaro-go/go2rtc/internalpkg/streams"
	"github.com/xaionaro-go/go2rtc/pkg/shell"
)

func main() {
	app.Init()
	streams.Init()

	rtsp.Init()

	shell.RunUntilSignal()
}
