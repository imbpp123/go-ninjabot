![Ninjabot Fork](https://user-images.githubusercontent.com/7620947/161434011-adc89d1a-dccb-45a7-8a07-2bb55e62d2d9.png)

> ⚠️ This is a **community fork** of the [original Ninjabot](https://github.com/rodrigo-brito/ninjabot) project.

The original author is not actively maintaining the project anymore (as of 2025), and since I'm actively working with automated trading systems and deeply interested in this domain, I decided to continue development based on my own needs.

This fork will evolve in a more opinionated direction, adapted for my strategies, workflow, and integrations — including possible support for other exchanges like Bybit.

If you're looking for the original version of the project, please refer to [rodrigo-brito/ninjabot](https://github.com/rodrigo-brito/ninjabot).

| DISCLAIMER                                                                                                                                                                                                           |
|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| This software is for educational purposes only. Do not risk money which you are afraid to lose.  USE THE SOFTWARE AT YOUR OWN RISK. THE AUTHORS AND ALL AFFILIATES ASSUME NO RESPONSIBILITY FOR YOUR TRADING RESULTS |

## Installation

`go get -u github.com/imbpp123/go-ninjabot/...`

## Examples of Usage

Check [examples](examples) directory:

- Paper Wallet (Live Simulation)
- Backtesting (Simulation with historical data)
- Real Account (Binance)

### CLI

To download historical data you can download ninjabot CLI from:

- Pre-build binaries in [release page](https://github.com/imbpp123/go-ninjabot/releases)
- Or with `go install github.com/imbpp123/go-ninjabot/cmd/ninjabot@latest`

**Example of usage**
```bash
# Download candles of BTCUSDT to btc.csv file (Last 30 days, timeframe 1D)
ninjabot download --pair BTCUSDT --timeframe 1d --days 30 --output ./btc.csv
```

### Backtesting Example

- Backtesting a custom strategy from [examples](examples) directory:
```
go run examples/backtesting/main.go
```

Output:

```
INFO[2023-03-25 13:54] [SETUP] Using paper wallet                   
INFO[2023-03-25 13:54] [SETUP] Initial Portfolio = 10000.000000 USDT 
---------+--------+-----+------+--------+--------+-----+----------+-----------+
|  PAIR   | TRADES | WIN | LOSS | % WIN  | PAYOFF | SQN |  PROFIT  |  VOLUME   |
+---------+--------+-----+------+--------+--------+-----+----------+-----------+
| ETHUSDT |      9 |   6 |    3 | 66.7 % |  3.407 | 1.3 | 21748.41 | 407769.64 |
| BTCUSDT |     14 |   6 |    8 | 42.9 % |  5.929 | 1.5 | 13511.66 | 448030.05 |
+---------+--------+-----+------+--------+--------+-----+----------+-----------+
|   TOTAL |     23 |  12 |   11 | 52.2 % |  4.942 | 1.4 | 35260.07 | 855799.68 |
+---------+--------+-----+------+--------+--------+-----+----------+-----------+

-- FINAL WALLET --
0.0000 BTC = 0.0000 USDT
0.0000 ETH = 0.0000 USDT
45260.0735 USDT

----- RETURNS -----
START PORTFOLIO     = 10000.00 USDT
FINAL PORTFOLIO     = 45260.07 USDT
GROSS PROFIT        =  35260.073493 USDT (352.60%)
MARKET CHANGE (B&H) =  407.09%

------ RISK -------
MAX DRAWDOWN = -11.76 %

------ VOLUME -----
BTCUSDT         = 448030.05 USDT
ETHUSDT         = 407769.64 USDT
TOTAL           = 855799.68 USDT
-------------------
Chart available at http://localhost:8080

```

### Plot result

<img width="100%"  src="https://user-images.githubusercontent.com/7620947/139601478-7b1d826c-f0f3-4766-951e-b11b1e1c9aa5.png" />

### Features

|                    	| Binance Spot 	| Binance Futures 	 |
|--------------------	|--------------	|-------------------|
| Order Market       	|       :ok:      	| :ok:              |
| Order Market Quote 	|       :ok:      	| 	                 |
| Order Limit        	|       :ok:      	| :ok:              |
| Order Stop         	|       :ok:      	| :ok:              |
| Order OCO          	|       :ok:     	| 	                 |
| Backtesting        	|       :ok:     	| :ok:         	    |

- [x] Backtesting
  - [x] Paper Wallet (Live Trading with fake wallet)
  - [x] Load Feed from CSV
  - [x] Order Limit, Market, Stop Limit, OCO

- [x] Bot Utilities
  - [x] CLI to download historical data
  - [x] Plot (Candles + Sell / Buy orders, Indicators)
  - [x] Telegram Controller (Status, Buy, Sell, and Notification)
  - [x] Heikin Ashi candle type support
  - [x] Trailing stop tool
  - [x] In app order scheduler

# Roadmap
  - [ ] Include Web UI Controller
  - [ ] Include more chart indicators - [Details](https://github.com/imbpp123/go-ninjabot/issues/110)

### Exchanges

Currently, we only support [Binance](https://www.binance.com/en?ref=35723227) exchange. If you want to include support for other exchanges, you need to implement a new `struct` that implements the interface `Exchange`. You can check some examples in [exchange](./pkg/exchange) directory.
