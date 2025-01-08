package bitget

const SPOT_ROOT_ROUTE = "/api/v2/spot"
const SPOT_GET_FILLS_SUB_ROUTE = "/trade/fills"

const FUTURE_ROOT_ROUTE = "/api/v2/mix"
const FUTURE_GET_TRANSACTION_SUB_ROUTE = "/order/fill-history"

const TAX_ROOT_ROUTE = "/api/v2/tax"
const TAX_GET_FUTURE_TRANSACTION_SUB_ROUTE = "/future-record"

type ApiQueryParameters struct {
	Route string
}
