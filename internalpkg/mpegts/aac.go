package mpegts

import (
	"net/http"

	"github.com/xaionaro-go/go2rtc/internalpkg/api"
	"github.com/xaionaro-go/go2rtc/internalpkg/streams"
	"github.com/xaionaro-go/go2rtc/pkg/aac"
)

func apiStreamAAC(w http.ResponseWriter, r *http.Request) {
	src := r.URL.Query().Get("src")
	stream := streams.Get(src)
	if stream == nil {
		http.Error(w, api.StreamNotFound, http.StatusNotFound)
		return
	}

	cons := aac.NewConsumer()
	cons.WithRequest(r)

	if err := stream.AddConsumer(cons); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "audio/aac")

	_, _ = cons.WriteTo(w)

	stream.RemoveConsumer(cons)
}
