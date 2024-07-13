package isapi

import (
	"github.com/xaionaro-go/go2rtc/internalpkg/streams"
	"github.com/xaionaro-go/go2rtc/pkg/core"
	"github.com/xaionaro-go/go2rtc/pkg/isapi"
)

func Init() {
	streams.HandleFunc("isapi", func(source string) (core.Producer, error) {
		return isapi.Dial(source)
	})
}
