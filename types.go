package krakenapi

import (
	"encoding/json"
	"math/big"
	"reflect"
	"strconv"
)

const (
	BUY    = "b"
	SELL   = "s"
	MARKET = "m"
	LIMIT  = "l"
)

// KrakenResponse wraps the Kraken API JSON response
type KrakenResponse struct {
	Error []string `json:"error"`
}

// TimeResponse represents the server's time
type TimeResponse struct {
	KrakenResponse
	Result struct {
		// Unix timestamp
		Unixtime int `json:"unixtime"`
		// RFC 1123 time format
		Rfc1123 string `json:"rfc1123"`
	} `json:"result"`
}

// AssetPairsResponse includes asset pair informations
type AssetPairsResponse struct {
	KrakenResponse
	Result map[string]AssetPairInfo `json:"result"`
}

// AssetPairInfo represents asset pair information
type AssetPairInfo struct {
	// Alternate pair name
	Altname string `json:"altname"`
	// Asset class of base component
	AssetClassBase string `json:"aclass_base"`
	// Asset id of base component
	Base string `json:"base"`
	// Asset class of quote component
	AssetClassQuote string `json:"aclass_quote"`
	// Asset id of quote component
	Quote string `json:"quote"`
	// Volume lot size
	Lot string `json:"lot"`
	// Scaling decimal places for pair
	PairDecimals int `json:"pair_decimals"`
	// Scaling decimal places for volume
	LotDecimals int `json:"lot_decimals"`
	// Amount to multiply lot volume by to get currency volume
	LotMultiplier int `json:"lot_multiplier"`
	// Array of leverage amounts available when buying
	LeverageBuy []float64 `json:"leverage_buy"`
	// Array of leverage amounts available when selling
	LeverageSell []float64 `json:"leverage_sell"`
	// Fee schedule array in [volume, percent fee] tuples
	Fees [][]float64 `json:"fees"`
	// // Maker fee schedule array in [volume, percent fee] tuples (if on maker/taker)
	FeesMaker [][]float64 `json:"fees_maker"`
	// // Volume discount currency
	FeeVolumeCurrency string `json:"fee_volume_currency"`
	// Margin call level
	MarginCall int `json:"margin_call"`
	// Stop-out/Liquidation margin level
	MarginStop int `json:"margin_stop"`
}

// AssetsResponse includes asset informations
type AssetsResponse struct {
	Assets map[string]AssetInfo
}

// AssetInfo represents an asset information
type AssetInfo struct {
	// Alternate name
	Altname string
	// Asset class
	AssetClass string `json:"aclass"`
	// Scaling decimal places for record keeping
	Decimals int
	// Scaling decimal places for output display
	DisplayDecimals int `json:"display_decimals"`
}

type BalanceResponse struct {
	Balances map[string]float64
}

type TradeBalanceResponse struct {
	EquivalentBalance         float64 `json:"eb,string"`
	TradeBalance              float64 `json:"tb,string"`
	MarginOP                  float64 `json:"m,string"`
	UnrealizedNetProfitLossOP float64 `json:"n,string"`
	CostBasisOP               float64 `json:"c,string"`
	CurrentValuationOP        float64 `json:"v,string"`
	Equity                    float64 `json:"e,string"`
	FreeMargin                float64 `json:"mf,string"`
	MarginLevel               float64 `json:"ml,string"`
}

// TickerResponse includes the requested ticker pairs
type TickerResponse struct {
	KrakenResponse
	Result map[string]PairTickerInfo `json:"result"`
}

// DepositAddressesResponse is the response type of a DepositAddresses query to the Kraken API.
type DepositAddressesResponse []struct {
	Address  string `json:"address"`
	Expiretm string `json:"expiretm"`
	New      bool   `json:"new,omitempty"`
}

// WithdrawResponse is the response type of a Withdraw query to the Kraken API.
type WithdrawResponse struct {
	RefID int `json:"refid"`
}

// WithdrawInfoResponse is the response type showing withdrawal information for a selected withdrawal method.
type WithdrawInfoResponse struct {
	Method string    `json:"method"`
	Limit  big.Float `json:"limit"`
	Amount big.Float `json:"amount"`
	Fee    big.Float `json:"fee"`
}

// GetPairTickerInfo is a helper method that returns given `pair`'s `PairTickerInfo`
func (v *TickerResponse) GetPairTickerInfo(pair string) PairTickerInfo {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(pair)

	return f.Interface().(PairTickerInfo)
}

// PairTickerInfo represents ticker information for a pair
type PairTickerInfo struct {
	// Ask array(<price>, <whole lot volume>, <lot volume>)
	Ask []string `json:"a"`
	// Bid array(<price>, <whole lot volume>, <lot volume>)
	Bid []string `json:"b"`
	// Last trade closed array(<price>, <lot volume>)
	Close []string `json:"c"`
	// Volume array(<today>, <last 24 hours>)
	Volume []string `json:"v"`
	// Volume weighted average price array(<today>, <last 24 hours>)
	VolumeAveragePrice []string `json:"p"`
	// Number of trades array(<today>, <last 24 hours>)
	Trades []int `json:"t"`
	// Low array(<today>, <last 24 hours>)
	Low []string `json:"l"`
	// High array(<today>, <last 24 hours>)
	High []string `json:"h"`
	// Today's opening price
	OpeningPrice float64 `json:"o,string"`
}

// TradesResponse represents a list of the last trades
type TradesResponse struct {
	KrakenResponse
	Result struct {
		Last   int `json:"last,string"`
		Trades map[string]TradeInfo
	} `json:"result"`
}

// TradesHistoryResponse represents a list executed trade
type TradesHistoryResponse struct {
	Trades map[string]TradeHistoryInfo `json:"trades"`
	Count  int                         `json:"count"`
}

type TradeHistoryInfo struct {
	TransactionId string  `json:"ordertxid"`
	PostxId       string  `json:"postxid"`
	AssetPair     string  `json:"pair"`
	Time          float64 `json:"time"`
	Type          string  `json:"type"`
	OrderType     string  `json:"ordertype"`
	Price         float64 `json:"price,string"`
	Cost          float64 `json:"cost,string"`
	Fee           float64 `json:"fee,string"`
	Volume        float64 `json:"vol,string"`
	Margin        float64 `json:"margin,string"`
	Misc          string  `json:"misc"`
}

// TradeInfo represents a trades information
type TradeInfo struct {
	Price         float64 `json:"0",string`
	Volume        float64 `json:"1",string`
	Time          float64 `json:"2"`
	BuySell       string  `json:"3"`
	MarketLimit   string  `json:"4"`
	Buy           bool
	Sell          bool
	Market        bool
	Limit         bool
	Miscellaneous string `json:"5"`
}

// LedgersResponse represents an associative array of ledgers infos
type LedgersResponse struct {
	Ledger map[string]LedgerInfo `json:"ledger"`
}

type LedgerInfo struct {
	RefID   string    `json:"refid"`
	Time    float64   `json:"time"`
	Type    string    `json:"type"`
	Aclass  string    `json:"aclass"`
	Asset   string    `json:"asset"`
	Amount  big.Float `json:"amount"`
	Fee     big.Float `json:"fee"`
	Balance big.Float `json:"balance"`
}

// OrderTypes for AddOrder
const (
	OTMarket              = "market"
	OTLimit               = "limit"                  // (price = limit price)
	OTStopLoss            = "stop-loss"              // (price = stop loss price)
	OTTakeProfi           = "take-profit"            // (price = take profit price)
	OTStopLossProfit      = "stop-loss-profit"       // (price = stop loss price, price2 = take profit price)
	OTStopLossProfitLimit = "stop-loss-profit-limit" // (price = stop loss price, price2 = take profit price)
	OTStopLossLimit       = "stop-loss-limit"        // (price = stop loss trigger price, price2 = triggered limit price)
	OTTakeProfitLimit     = "take-profit-limit"      // (price = take profit trigger price, price2 = triggered limit price)
	OTTrailingStop        = "trailing-stop"          // (price = trailing stop offset)
	OTTrailingStopLimit   = "trailing-stop-limit"    // (price = trailing stop offset, price2 = triggered limit offset)
	OTStopLossAndLimit    = "stop-loss-and-limit"    // (price = stop loss price, price2 = limit price)
	OTSettlePosition      = "settle-position"
)

// OrderDescription represents an orders description
type OrderDescription struct {
	AssetPair      string `json:"pair"`
	Close          string `json:"close"`
	Leverage       string `json:"leverage"`
	Order          string `json:"order"`
	OrderType      string `json:"ordertype"`
	PrimaryPrice   string `json:"price"`
	SecondaryPrice string `json:"price2"`
	Type           string `json:"type"`
}

// Order represents a single order
type Order struct {
	TransactionID  string           `json:"-"`
	ReferenceID    string           `json:"refid"`
	UserRef        int              `json:"userref"`
	Status         string           `json:"status"`
	OpenTime       float64          `json:"opentm"`
	StartTime      float64          `json:"starttm"`
	ExpireTime     float64          `json:"expiretm"`
	Description    OrderDescription `json:"descr"`
	Volume         string           `json:"vol"`
	VolumeExecuted float64          `json:"vol_exec,string"`
	Cost           float64          `json:"cost,string"`
	Fee            float64          `json:"fee,string"`
	Price          float64          `json:"price,string"`
	StopPrice      float64          `json:"stopprice.string"`
	LimitPrice     float64          `json:"limitprice,string"`
	Misc           string           `json:"misc"`
	OrderFlags     string           `json:"oflags"`
	CloseTime      float64          `json:"closetm"`
	Reason         string           `json:"reason"`
}

// ClosedOrdersResponse represents a list of closed orders, indexed by id
type ClosedOrdersResponse struct {
	Closed map[string]Order `json:"closed"`
	Count  int              `json:"count"`
}

// OrderBookItem is a piece of information about an order.
type OrderBookItem struct {
	Price  float64
	Amount float64
	Ts     int64
}

// UnmarshalJSON takes a json array from kraken and converts it into an OrderBookItem.
func (o *OrderBookItem) UnmarshalJSON(data []byte) error {
	tmp_struct := struct {
		price  string
		amount string
		ts     int64
	}{}
	tmp_arr := []interface{}{&tmp_struct.price, &tmp_struct.amount, &tmp_struct.ts}
	err := json.Unmarshal(data, &tmp_arr)
	if err != nil {
		return err
	}

	o.Price, err = strconv.ParseFloat(tmp_struct.price, 64)
	if err != nil {
		return err
	}
	o.Amount, err = strconv.ParseFloat(tmp_struct.amount, 64)
	if err != nil {
		return err
	}
	o.Ts = tmp_struct.ts
	return nil
}

// DepthResponse is a response from kraken to Depth request.
type DepthResponse struct {
	KrakenResponse
	Result map[string]OrderBook `json:"result"`
}

// OrderBook contains top asks and bids.
type OrderBook struct {
	Asks []OrderBookItem `json:"asks"`
	Bids []OrderBookItem `json:"bids"`
}

type OpenOrdersResponse struct {
	Open  map[string]Order `json:"open"`
	Count int              `json:"count"`
}

type AddOrderResponse struct {
	Description    OrderDescription `json:"descr"`
	TransactionIds []string         `json:"txid"`
}

type CancelOrderResponse struct {
	Count   int  `json:"count"`
	Pending bool `json:"pending"`
}

type QueryOrdersResponse map[string]Order
