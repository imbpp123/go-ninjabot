package exchange

type OrderType = string

const (
	OrderTypeLimit  OrderType = "Limit"
	OrderTypeMarket OrderType = "Market"
)

type OrderSide = string

const (
	OrderSideBuy  OrderSide = "Buy"
	OrderSideSell OrderSide = "Sell"
)

type MarginType = string

const (
	MarginTypeIsolated MarginType = "ISOLATED"
	MarginTypeCross    MarginType = "CROSSED"
)

type KlineInterval = string

const (
	KlineInterval1Min   KlineInterval = "1"
	KlineInterval3Min   KlineInterval = "3"
	KlineInterval5Min   KlineInterval = "5"
	KlineInterval15Min  KlineInterval = "15"
	KlineInterval30Min  KlineInterval = "30"
	KlineInterval1Hour  KlineInterval = "60"
	KlineInterval2Hour  KlineInterval = "120"
	KlineInterval4Hour  KlineInterval = "240"
	KlineInterval6Hour  KlineInterval = "360"
	KlineInterval12Hour KlineInterval = "720"
	KlineInterval1Day   KlineInterval = "D"
	KlineInterval1Week  KlineInterval = "W"
	KlineInterval1Month KlineInterval = "M"
)

type PairOption struct {
	Pair       string
	Leverage   int
	MarginType MarginType
}
