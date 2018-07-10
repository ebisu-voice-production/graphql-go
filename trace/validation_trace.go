package trace

import (
	"github.com/ebisu-voice-production/graphql-go/errors"
)

type ValidationTracer interface {
	TraceValidation() TraceQueryFinishFunc
}

type NoopValidationTracer struct{}

func (NoopValidationTracer) TraceValidation() TraceQueryFinishFunc {
	return func(errs []*errors.QueryError) {}
}
