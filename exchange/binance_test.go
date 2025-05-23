package exchange

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/imbpp123/go-ninjabot/model"
)

func TestFormatQuantity(t *testing.T) {
	assets := NewAssetInfo()
	assets.Set("BTCUSDT", model.AssetInfo{
		StepSize:           0.00001000,
		TickSize:           0.00001000,
		BaseAssetPrecision: 5,
		QuotePrecision:     5,
	})
	assets.Set("BATUSDT", model.AssetInfo{
		StepSize:           0.01,
		TickSize:           0.01,
		BaseAssetPrecision: 2,
		QuotePrecision:     2,
	})
	binance := Binance{assetsInfo: assets}

	tt := []struct {
		pair     string
		quantity float64
		expected string
	}{
		{"BTCUSDT", 1.1, "1.1"},
		{"BTCUSDT", 11, "11"},
		{"BTCUSDT", 11, "11"},
		{"BTCUSDT", 1.1111111111, "1.11111"},
		{"BTCUSDT", 1.9999999999999, "1.99999"},
		{"BTCUSDT", 1111111.1111111111, "1111111.11111"},
		{"BATUSDT", 111.111, "111.11"},
		{"BATUSDT", 9.9999999999, "9.99"},
		{"BATUSDT", 9.9999999999, "9.99"},
		{"BATUSDT", 10, "10"},
		{"BATUSDT", 10.11111, "10.11"},
		{"BATUSDT", 0.01, "0.01"},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("given %f %s", tc.quantity, tc.pair), func(t *testing.T) {
			require.Equal(t, tc.expected, binance.formatQuantity(tc.pair, tc.quantity))
			require.Equal(t, tc.expected, binance.formatPrice(tc.pair, tc.quantity))
		})
	}
}
