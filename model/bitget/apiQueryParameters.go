package bitget

const SPOT_ROOT_ROUTE = "/api/v2/spot"
const SPOT_GET_FILLS_SUB_ROUTE = "/trade/fills"

const FUTURE_ROOT_ROUTE = "/api/v2/mix"
const FUTURE_GET_TRANSACTION_SUB_ROUTE = "/order/fill-history"

type ApiQueryParameters struct {
	Route string
}
