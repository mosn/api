package api

import (
	"context"
)

const RequestTranscodeFail ResponseFlag = 0x2000

// Transcoder provide ability to transcoding request/response
type Transcoder interface {
	// Accept
	Accept(ctx context.Context, headers HeaderMap, buf IoBuffer, trailers HeaderMap) bool
	// TranscodingRequest
	TranscodingRequest(ctx context.Context, headers HeaderMap, buf IoBuffer, trailers HeaderMap) (HeaderMap, IoBuffer, HeaderMap, error)
	// TranscodingResponse
	TranscodingResponse(ctx context.Context, headers HeaderMap, buf IoBuffer, trailers HeaderMap) (HeaderMap, IoBuffer, HeaderMap, error)
}
