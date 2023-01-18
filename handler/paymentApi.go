package handler

import (
	"context"
	"errors"
	common "github.com/bufengmobuganhuo/micro-service-common"
	go_micro_service_payment "github.com/bufengmobuganhuo/micro-service-payment/proto/payment"
	paymentApi "github.com/bufengmobuganhuo/micro-service-paymentApi/proto/paymentApi"
	"net/http"
	"strconv"
)

type PaymentApi struct {
	PaymentService go_micro_service_payment.PaymentService
}

var (
	ClientId string = ""
)

// AliPayRefund 通过API向外暴露为/paymentApi/aliPayRefund, 接收http请求
func (p PaymentApi) AliPayRefund(ctx context.Context, request *paymentApi.Request,
	response *paymentApi.Response) error {
	// 验证入参
	if err := isOk("payment_id", request); err != nil {
		response.StatusCode = http.StatusBadRequest
		return err
	}
	if err := isOk("refund_id", request); err != nil {
		response.StatusCode = http.StatusBadRequest
		return err
	}
	if err := isOk("money", request); err != nil {
		response.StatusCode = http.StatusBadRequest
		return err
	}
	paymentId, err := strconv.ParseInt(request.Get["payment_id"].Values[0], 10, 64)
	if err != nil {
		common.Error(err)
		return err
	}
	paymentInfo, err := p.PaymentService.FindPaymentByID(ctx, &go_micro_service_payment.PaymentID{PaymentId: paymentId})
	if err != nil {
		common.Error(err)
		return err
	}
	common.Info(paymentInfo)
	response.StatusCode = http.StatusOK
	response.Body = request.Get["refund_id"].Values[0] + "退款成功"
	return nil
}

func isOk(key string, request *paymentApi.Request) error {
	if _, ok := request.Get[key]; !ok {
		err := errors.New(key + " 参数异常")
		common.Error(err)
		return err
	}
	return nil
}
