package server

import (
	"io"
	"os"

	"github.com/opentracing/opentracing-go"

	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
)

func newJaegerConf() jaegercfg.Configuration {
	return jaegercfg.Configuration{
		ServiceName: os.Getenv("JAEGER_SERVICE_NAME"),
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}
}

func newTracer() (io.Closer, error) {
	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	cfg := newJaegerConf()
	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
	if err != nil {
		return nil, err
	}
	opentracing.SetGlobalTracer(tracer)
	return closer, nil
}
