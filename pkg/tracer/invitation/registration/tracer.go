package registration

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/inspire/invitation/registration"
)

func trace(span trace1.Span, in *npool.RegistrationReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("InviterID.%v", index), in.GetInviterID()),
		attribute.String(fmt.Sprintf("InviteeID.%v", index), in.GetInviteeID()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.RegistrationReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Value", in.GetID().GetValue()),
		attribute.String("AppID.Op", in.GetAppID().GetOp()),
		attribute.String("AppID.Value", in.GetAppID().GetValue()),
		attribute.String("InviterID.Op", in.GetInviterID().GetOp()),
		attribute.String("InviterID.Value", in.GetInviterID().GetValue()),
		attribute.String("InviteeID.Op", in.GetInviteeID().GetOp()),
		attribute.String("InviteeID.Value", in.GetInviteeID().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.RegistrationReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
