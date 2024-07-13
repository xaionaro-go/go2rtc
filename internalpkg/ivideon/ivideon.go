package ivideon

import (
	"github.com/xaionaro-go/go2rtc/internalpkg/streams"
	"github.com/xaionaro-go/go2rtc/pkg/core"
	"github.com/xaionaro-go/go2rtc/pkg/ivideon"
)

func Init() {
	streams.HandleFunc("ivideon", func(source string) (core.Producer, error) {
		return ivideon.Dial(source)
	})
}
