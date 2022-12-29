package general

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/archivement/general"
)

func trace(span trace1.Span, in *npool.GeneralReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("UserID.%v", index), in.GetUserID()),
		attribute.String(fmt.Sprintf("GoodID.%v", index), in.GetGoodID()),
		attribute.String(fmt.Sprintf("CoinTypeID.%v", index), in.GetCoinTypeID()),
		attribute.String(fmt.Sprintf("TotalAmount.%v", index), in.GetTotalAmount()),
		attribute.String(fmt.Sprintf("SelfAmount.%v", index), in.GetSelfAmount()),
		attribute.Int64(fmt.Sprintf("TotalUnits.%v", index), int64(in.GetTotalUnits())),
		attribute.Int64(fmt.Sprintf("SelfUnits.%v", index), int64(in.GetSelfUnits())),
		attribute.String(fmt.Sprintf("TotalCommission.%v", index), in.GetTotalCommission()),
		attribute.String(fmt.Sprintf("SelfCommission.%v", index), in.GetSelfCommission()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.GeneralReq) trace1.Span {
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
		attribute.String("GoodID.Op", in.GetGoodID().GetOp()),
		attribute.String("GoodID.Value", in.GetGoodID().GetValue()),
		attribute.String("CoinTypeID.Op", in.GetCoinTypeID().GetOp()),
		attribute.String("CoinTypeID.Value", in.GetCoinTypeID().GetValue()),
		attribute.String("TotalAmount.Op", in.GetTotalAmount().GetOp()),
		attribute.String("TotalAmount.Value", in.GetTotalAmount().GetValue()),
		attribute.String("SelfAmount.Op", in.GetSelfAmount().GetOp()),
		attribute.String("SelfAmount.Value", in.GetSelfAmount().GetValue()),
		attribute.String("TotalUnits.Op", in.GetTotalUnits().GetOp()),
		attribute.Int64("TotalUnits.Value", int64(in.GetTotalUnits().GetValue())),
		attribute.String("SelfUnits.Op", in.GetSelfUnits().GetOp()),
		attribute.Int64("SelfUnits.Value", int64(in.GetSelfUnits().GetValue())),
		attribute.String("TotalCommission.Op", in.GetTotalCommission().GetOp()),
		attribute.String("TotalCommission.Value", in.GetTotalCommission().GetValue()),
		attribute.String("SelfCommission.Op", in.GetSelfCommission().GetOp()),
		attribute.String("SelfCommission.Value", in.GetSelfCommission().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.GeneralReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
