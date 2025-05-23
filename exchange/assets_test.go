package exchange

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/imbpp123/go-ninjabot/model"
)

func TestAssetInfoSetGet(t *testing.T) {
	a := NewAssetInfo()
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
	a := NewAssetInfo()

	_, err := a.Get("INVALID")
	assert.Error(t, err)
}

func TestValidateQuantity(t *testing.T) {
	a := NewAssetInfo()
	a.Set("BTCUSDT", model.AssetInfo{
		MinQuantity: 0.01,
		MaxQuantity: 5.0,
	})

	assert.NoError(t, a.ValidateQuantity("BTCUSDT", 1.0))
	assert.Error(t, a.ValidateQuantity("BTCUSDT", 0.001))
	assert.Error(t, a.ValidateQuantity("BTCUSDT", 10.0))
}

func TestValidateLeverage(t *testing.T) {
	a := NewAssetInfo()
	a.Set("BTCUSDT", model.AssetInfo{
		MinLeverage: 2,
		MaxLeverage: 5,
	})

	assert.NoError(t, a.ValidateLeverage("BTCUSDT", 3))
	assert.Error(t, a.ValidateLeverage("BTCUSDT", 1))
	assert.Error(t, a.ValidateLeverage("BTCUSDT", 10.0))
}

func TestCalculateLotQuantity(t *testing.T) {
	a := NewAssetInfo()
	a.Set("BTCUSDT", model.AssetInfo{
		StepSize:           0.001,
		BaseAssetPrecision: 6,
	})

	qty, err := a.CalculateLotQuantity("BTCUSDT", 1.234567)
	assert.NoError(t, err)
	assert.Equal(t, 1.234, qty)
}

func TestCalculateLotQuantityError(t *testing.T) {
	a := NewAssetInfo()

	qty, err := a.CalculateLotQuantity("BTCUSDT", 1.234567)
	assert.ErrorIs(t, err, ErrInvalidAsset)
	assert.Equal(t, 1.234567, qty)
}

func TestCalculateLotPrice(t *testing.T) {
	a := NewAssetInfo()
	a.Set("BTCUSDT", model.AssetInfo{
		TickSize:       0.01,
		QuotePrecision: 2,
	})

	price, err := a.CalculateLotPrice("BTCUSDT", 1234.5678)
	assert.NoError(t, err)
	assert.Equal(t, 1234.56, price)
}

func TestCalculateLotPriceError(t *testing.T) {
	a := NewAssetInfo()

	price, err := a.CalculateLotPrice("BTCUSDT", 1234.5678)
	assert.ErrorIs(t, err, ErrInvalidAsset)
	assert.Equal(t, 1234.5678, price)
}
