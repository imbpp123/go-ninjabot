package indicator

import (
	"fmt"
	"time"

	"github.com/imbpp123/go-ninjabot/model"
	"github.com/imbpp123/go-ninjabot/plot"

	"github.com/markcheno/go-talib"
)

func RSI(period int, color string) plot.Indicator {
	return &rsi{
		Period: period,
		Color:  color,
	}
}

type rsi struct {
	Period int
	Color  string
	Values model.Series[float64]
	Time   []time.Time
}

func (e rsi) Warmup() int {
	return e.Period
}

func (e rsi) Name() string {
	return fmt.Sprintf("RSI(%d)", e.Period)
}

func (e rsi) Overlay() bool {
	return false
}

func (e *rsi) Load(dataframe *model.Dataframe) {
	if len(dataframe.Time) < e.Period {
		return
	}

	e.Values = talib.Rsi(dataframe.Close, e.Period)[e.Period:]
	e.Time = dataframe.Time[e.Period:]
}

func (e rsi) Metrics() []plot.IndicatorMetric {
	return []plot.IndicatorMetric{
		{
			Color:  e.Color,
			Style:  "line",
			Values: e.Values,
			Time:   e.Time,
		},
	}
}
