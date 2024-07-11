package server

import (
	"kratos-test/internal/es"
	netHttp "net/http"

	"github.com/go-kratos/kratos/v2/transport/http"
)

func errorEncoder(w netHttp.ResponseWriter, r *netHttp.Request, err error) {

	var se = es.FromError(err)
	codec, _ := http.CodecForRequest(r, "Accept")
	body, err := codec.Marshal(se)
	if err != nil {
		w.WriteHeader(netHttp.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(se.Code)
	_, _ = w.Write(body)
}
