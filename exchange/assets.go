package exchange

import (
	"errors"
	"fmt"
	"math"
	"sync"

	"github.com/imbpp123/go-ninjabot/model"
)

type AssetInfo struct {
	data     map[string]model.AssetInfo
	dataLock sync.RWMutex
}

var (
	ErrInvalidAsset = errors.New("invalid asset")
)

func NewAssetInfo() *AssetInfo {
	return &AssetInfo{
		data: make(map[string]model.AssetInfo),
	}
}

func (a *AssetInfo) Set(pair string, info model.AssetInfo) {
	a.dataLock.Lock()
	defer a.dataLock.Unlock()

	a.data[pair] = info
}

func (a *AssetInfo) Get(pair string) (model.AssetInfo, error) {
	a.dataLock.RLock()
	defer a.dataLock.RUnlock()

	asset, ok := a.data[pair]
	if !ok {
		return model.AssetInfo{}, fmt.Errorf("%w: %s", ErrInvalidAsset, pair)
	}

	return asset, nil
}

func (a *AssetInfo) ValidateQuantity(pair string, quantity float64) error {
	info, err := a.Get(pair)
	if err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	if quantity > info.MaxQuantity || quantity < info.MinQuantity {
		return fmt.Errorf("%w: min: %f max: %f", ErrInvalidQuantity, info.MinQuantity, info.MaxQuantity)
	}

	return nil
}

func (a *AssetInfo) CalculateLotQuantity(pair string, quantity float64) (float64, error) {
	info, err := a.Get(pair)
	if err != nil {
		return quantity, fmt.Errorf("calculate lot quantity convert error: %w", err)
	}

	return math.Trunc(math.Floor(quantity/info.StepSize)*info.StepSize*math.Pow10(info.BaseAssetPrecision)) / math.Pow10(info.BaseAssetPrecision), nil
}

func (a *AssetInfo) CalculateLotPrice(pair string, price float64) (float64, error) {
	info, err := a.Get(pair)
	if err != nil {
		return price, fmt.Errorf("calculate lot price convert error: %w", err)
	}

	return math.Trunc(math.Floor(price/info.TickSize)*info.TickSize*math.Pow10(info.QuotePrecision)) / math.Pow10(info.QuotePrecision), nil
}
