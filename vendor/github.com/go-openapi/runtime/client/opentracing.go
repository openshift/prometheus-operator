package client

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/go-openapi/strfmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"

	"github.com/go-openapi/runtime"
)

type tracingTransport struct {
	transport runtime.ClientTransport
	host      string
	opts      []opentracing.StartSpanOption
}

func newOpenTracingTransport(transport runtime.ClientTransport, host string, opts []opentracing.StartSpanOption,
) runtime.ClientTransport {
	return &tracingTransport{
		transport: transport,
		host:      host,
		opts:      opts,
	}
}

func (t *tracingTransport) Submit(op *runtime.ClientOperation) (interface{}, error) {
	authInfo := op.AuthInfo
	reader := op.Reader

	var span opentracing.Span
	defer func() {
		if span != nil {
			span.Finish()
		}
	}()

	op.AuthInfo = runtime.ClientAuthInfoWriterFunc(func(req runtime.ClientRequest, reg strfmt.Registry) error {
		span = createClientSpan(op, req.GetHeaderParams(), t.host, t.opts)
		if authInfo == nil {
			return nil
		}
		return authInfo.AuthenticateRequest(req, reg)
	})

	op.Reader = runtime.ClientResponseReaderFunc(func(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
		if span != nil {
			code := response.Code()
			ext.HTTPStatusCode.Set(span, uint16(code))
			if code >= 400 {
				ext.Error.Set(span, true)
			}
		}
		return reader.ReadResponse(response, consumer)
	})

	submit, err := t.transport.Submit(op)
	if err != nil && span != nil {
		ext.Error.Set(span, true)
		span.LogFields(log.Error(err))
	}
	return submit, err
}

func createClientSpan(op *runtime.ClientOperation, header http.Header, host string,
	opts []opentracing.StartSpanOption) opentracing.Span {
	ctx := op.Context
	span := opentracing.SpanFromContext(ctx)

	if span != nil {
		opts = append(opts, ext.SpanKindRPCClient)
		span, _ = opentracing.StartSpanFromContextWithTracer(
			ctx, span.Tracer(), toSnakeCase(op.ID), opts...)

		ext.Component.Set(span, "go-openapi")
		ext.PeerHostname.Set(span, host)
		span.SetTag("http.path", op.PathPattern)
		ext.HTTPMethod.Set(span, op.Method)

		_ = span.Tracer().Inject(
			span.Context(),
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(header))

		return span
	}
	return nil
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}