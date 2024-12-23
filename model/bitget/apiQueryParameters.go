package bitget

const SPOT_ROOT_ROUTE = "/api/v2/spot"
const SPOT_GET_FILLS_SUB_ROUTE = "/trade/fills"

type ApiQueryParameters struct {
	Route string
}
