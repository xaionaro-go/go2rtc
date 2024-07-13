package main

import (
	"github.com/xaionaro-go/go2rtc/internalpkg/api"
	"github.com/xaionaro-go/go2rtc/internalpkg/app"
	"github.com/xaionaro-go/go2rtc/internalpkg/hass"
	"github.com/xaionaro-go/go2rtc/internalpkg/streams"
	"github.com/xaionaro-go/go2rtc/pkg/shell"
)

func main() {
	app.Init()
	streams.Init()

	api.Init()

	hass.Init()

	shell.RunUntilSignal()
}
