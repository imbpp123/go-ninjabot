package exchange_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/imbpp123/go-ninjabot/exchange"
	"github.com/imbpp123/go-ninjabot/model"
)

func TestAssetInfoSetGet(t *testing.T) {
	a := exchange.NewAssetInfo()
	info := model.AssetInfo{
		BaseAsset:          "BTC",
		QuoteAsset:         "USDT",
		BaseAssetPrecision: 6,
		QuotePrecision:     2,
		MinQuantity:        0.001,
		MaxQuantity:        100.0,
		StepSize:           0.001,
		MinPrice:           10.0,
		MaxPrice:           100000.0,
		TickSize:           0.01,
	}

	a.Set("BTCUSDT", info)

	ret, err := a.Get("BTCUSDT")
	assert.NoError(t, err)
	assert.Equal(t, info, ret)
}

func TestAssetInfoGetInvalid(t *testing.T) {
	a := exchange.NewAssetInfo()

	_, err := a.Get("INVALID")
	assert.Error(t, err)
}

func TestValidateQuantity(t *testing.T) {
	a := exchange.NewAssetInfo()
	a.Set("BTCUSDT", model.AssetInfo{
		MinQuantity: 0.01,
		MaxQuantity: 5.0,
	})

	assert.NoError(t, a.ValidateQuantity("BTCUSDT", 1.0))
	assert.Error(t, a.ValidateQuantity("BTCUSDT", 0.001))
	assert.Error(t, a.ValidateQuantity("BTCUSDT", 10.0))
}

func TestCalculateLotQuantity(t *testing.T) {
	a := exchange.NewAssetInfo()
	a.Set("BTCUSDT", model.AssetInfo{
		StepSize:           0.001,
		BaseAssetPrecision: 6,
	})

	qty, err := a.CalculateLotQuantity("BTCUSDT", 1.234567)
	assert.NoError(t, err)
	assert.Equal(t, 1.234, qty)
}

func TestCalculateLotPrice(t *testing.T) {
	a := exchange.NewAssetInfo()
	a.Set("BTCUSDT", model.AssetInfo{
		TickSize:       0.01,
		QuotePrecision: 2,
	})

	price, err := a.CalculateLotPrice("BTCUSDT", 1234.5678)
	assert.NoError(t, err)
	assert.Equal(t, 1234.56, price)
}
