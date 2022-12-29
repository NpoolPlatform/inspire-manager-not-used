package detail

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/archivement/detail"
)

func trace(span trace1.Span, in *npool.DetailReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("UserID.%v", index), in.GetUserID()),
		attribute.String(fmt.Sprintf("DirectContributorID.%v", index), in.GetDirectContributorID()),
		attribute.String(fmt.Sprintf("GoodID.%v", index), in.GetGoodID()),
		attribute.String(fmt.Sprintf("OrderID.%v", index), in.GetOrderID()),
		attribute.Bool(fmt.Sprintf("SelfOrder.%v", index), in.GetSelfOrder()),
		attribute.String(fmt.Sprintf("PaymentID.%v", index), in.GetPaymentID()),
		attribute.String(fmt.Sprintf("CoinTypeID.%v", index), in.GetCoinTypeID()),
		attribute.String(fmt.Sprintf("PaymentCoinTypeID.%v", index), in.GetPaymentCoinTypeID()),
		attribute.String(fmt.Sprintf("PaymentCoinUSDCurrency.%v", index), in.GetPaymentCoinUSDCurrency()),
		attribute.String(fmt.Sprintf("Amount.%v", index), in.GetAmount()),
		attribute.String(fmt.Sprintf("USDAmount.%v", index), in.GetUSDAmount()),
		attribute.Int64(fmt.Sprintf("Units.%v", index), int64(in.GetUnits())),
	)
	return span
}

func Trace(span trace1.Span, in *npool.DetailReq) trace1.Span {
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
		attribute.String("DirectContributorID.Op", in.GetDirectContributorID().GetOp()),
		attribute.String("DirectContributorID.Value", in.GetDirectContributorID().GetValue()),
		attribute.String("GoodID.Op", in.GetGoodID().GetOp()),
		attribute.String("GoodID.Value", in.GetGoodID().GetValue()),
		attribute.String("OrderID.Op", in.GetOrderID().GetOp()),
		attribute.String("OrderID.Value", in.GetOrderID().GetValue()),
		attribute.String("SelfOrder.Op", in.GetSelfOrder().GetOp()),
		attribute.Bool("SelfOrder.Value", in.GetSelfOrder().GetValue()),
		attribute.String("PaymentID.Op", in.GetPaymentID().GetOp()),
		attribute.String("PaymentID.Value", in.GetPaymentID().GetValue()),
		attribute.String("CoinTypeID.Op", in.GetCoinTypeID().GetOp()),
		attribute.String("CoinTypeID.Value", in.GetCoinTypeID().GetValue()),
		attribute.String("PaymentCoinTypeID.Op", in.GetPaymentCoinTypeID().GetOp()),
		attribute.String("PaymentCoinTypeID.Value", in.GetPaymentCoinTypeID().GetValue()),
		attribute.String("PaymentCoinUSDCurrency.Op", in.GetPaymentCoinUSDCurrency().GetOp()),
		attribute.String("PaymentCoinUSDCurrency.Value", in.GetPaymentCoinUSDCurrency().GetValue()),
		attribute.String("Amount.Op", in.GetAmount().GetOp()),
		attribute.String("Amount.Value", in.GetAmount().GetValue()),
		attribute.String("USDAmount.Op", in.GetUSDAmount().GetOp()),
		attribute.String("USDAmount.Value", in.GetUSDAmount().GetValue()),
		attribute.String("Units.Op", in.GetUnits().GetOp()),
		attribute.Int64("Units.Value", int64(in.GetUnits().GetValue())),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.DetailReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
