package runtime

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang/protobuf/jsonpb"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/herryg91/cdd/grst/errors"
	"google.golang.org/protobuf/proto"
)

var jsonpbmarshal = &jsonpb.Marshaler{EmitDefaults: true}

func ForwardResponseMessage(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, req *http.Request, resp proto.Message, opts ...func(context.Context, http.ResponseWriter, proto.Message) error) {
	// var buf bytes.Buffer
	buf, err := marshaler.Marshal(resp)
	if err != nil {
		errors.HTTPError(ctx, mux, marshaler, w, req, err)
		return
	}
	var data interface{}
	// log if something goes wrong with unmarshalling response
	if err := json.Unmarshal(buf, &data); err != nil {
		errors.HTTPError(ctx, mux, marshaler, w, req, err)
		return
	}

	// Parse Start Time
	latency := "0ms"
	starttime, errParse := time.Parse(time.RFC3339Nano, req.Header.Get("grst.starttime"))
	if errParse == nil {
		latency = time.Since(starttime).String()
	}
	// template key value for REST response
	formattedResponse := &ResponseSuccess{
		Data:        resp,
		ProcessTime: latency,
		HTTPStatus:  http.StatusOK,
	}

	runtime.ForwardResponseMessage(ctx, mux, marshaler, w, req, formattedResponse, opts...)
}
