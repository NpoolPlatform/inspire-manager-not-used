package specialoffer

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/specialoffer"
)

func trace(span trace1.Span, in *npool.SpecialOfferReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("UserID.%v", index), in.GetUserID()),
		attribute.String(fmt.Sprintf("Amount.%v", index), in.GetAmount()),
		attribute.String(fmt.Sprintf("ReleaseByUserID.%v", index), in.GetReleaseByUserID()),
		attribute.Int(fmt.Sprintf("StartAt.%v", index), int(in.GetStartAt())),
		attribute.Int(fmt.Sprintf("DurationDays.%v", index), int(in.GetDurationDays())),
		attribute.String(fmt.Sprintf("Message.%v", index), in.GetMessage()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.SpecialOfferReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Value", in.GetID().GetValue()),
		attribute.String("AppID.Op", in.GetAppID().GetOp()),
		attribute.String("AppID.Value", in.GetAppID().GetValue()),
		attribute.String("StartAt.Op", in.GetStartAt().GetOp()),
		attribute.Int("StartAt.Value", int(in.GetStartAt().GetValue())),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.SpecialOfferReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
