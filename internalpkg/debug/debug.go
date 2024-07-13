package debug

import (
	"github.com/xaionaro-go/go2rtc/internalpkg/api"
)

func Init() {
	api.HandleFunc("api/stack", stackHandler)
}
