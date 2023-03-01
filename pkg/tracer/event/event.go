package event

import (
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/event"
)

func trace(span trace1.Span, in *npool.EventReq, index int) trace1.Span {
	return span
}

func Trace(span trace1.Span, in *npool.EventReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	return span
}

func TraceMany(span trace1.Span, infos []*npool.EventReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
