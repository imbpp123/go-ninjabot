package indicator

import (
	"fmt"
	"time"

	"github.com/imbpp123/go-ninjabot/model"
	"github.com/imbpp123/go-ninjabot/plot"

	"github.com/markcheno/go-talib"
)

func EMA(period int, color string) plot.Indicator {
	return &ema{
		Period: period,
		Color:  color,
	}
}

type ema struct {
	Period int
	Color  string
	Values model.Series[float64]
	Time   []time.Time
}

func (e ema) Warmup() int {
	return e.Period
}

func (e ema) Name() string {
	return fmt.Sprintf("EMA(%d)", e.Period)
}

func (e ema) Overlay() bool {
	return true
}

func (e *ema) Load(dataframe *model.Dataframe) {
	if len(dataframe.Time) < e.Period {
		return
	}

	e.Values = talib.Ema(dataframe.Close, e.Period)[e.Period:]
	e.Time = dataframe.Time[e.Period:]
}

func (e ema) Metrics() []plot.IndicatorMetric {
	return []plot.IndicatorMetric{
		{
			Style:  "line",
			Color:  e.Color,
			Values: e.Values,
			Time:   e.Time,
		},
	}
}
