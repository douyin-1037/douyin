package service

import (
	"github.com/opentracing/opentracing-go"
	"io"
)

var Tracer opentracing.Tracer
var Closer io.Closer
