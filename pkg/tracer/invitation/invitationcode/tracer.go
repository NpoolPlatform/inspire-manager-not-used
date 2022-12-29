package invitationcode

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/invitationcode"
)

func trace(span trace1.Span, in *npool.InvitationCodeReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("UserID.%v", index), in.GetUserID()),
		attribute.String(fmt.Sprintf("InvitationCode.%v", index), in.GetInvitationCode()),
		attribute.Bool(fmt.Sprintf("Confirmed.%v", index), in.GetConfirmed()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.InvitationCodeReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Value", in.GetID().GetValue()),
		attribute.String("AppID.Op", in.GetAppID().GetOp()),
		attribute.String("AppID.Value", in.GetAppID().GetValue()),
		attribute.String("UserID.Op", in.GetUserID().GetOp()),
		attribute.String("UserID.Value", in.GetUserID().GetValue()),
		attribute.String("InvitationCode.Op", in.GetInvitationCode().GetOp()),
		attribute.String("InvitationCode.Value", in.GetInvitationCode().GetValue()),
		attribute.String("Confirmed.Op", in.GetConfirmed().GetOp()),
		attribute.Bool("Confirmed.Value", in.GetConfirmed().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.InvitationCodeReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
