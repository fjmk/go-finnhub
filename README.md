# go-finnhub

> This repository was forked from github.com/llonchj/go-finnhub. Thanks @m1!
> At this point, the API is unstable and might change.

[![GoDoc](https://godoc.org/github.com/llonchj/go-finnhub?status.svg)](https://godoc.org/github.com/llonchj/go-finnhub)
[![Build Status](https://travis-ci.org/llonchj/go-finnhub.svg?branch=master)](https://travis-ci.org/llonchj/go-finnhub)
[![Go Report Card](https://goreportcard.com/badge/github.com/llonchj/go-finnhub)](https://goreportcard.com/report/github.com/llonchj/go-finnhub)
[![Release](https://img.shields.io/github/release/llonchj/go-finnhub.svg)](https://github.com/llonchj/go-finnhub/releases/latest)
[![codecov](https://codecov.io/gh/llonchj/go-finnhub/branch/master/graph/badge.svg)](https://codecov.io/gh/llonchj/go-finnhub)

__Simple and easy to use client for stock, forex and crpyto data from [finnhub.io](https://finnhub.io/) written in Go. Access real-time market data from 60+ stock exchanges, 10 forex brokers, and 15+ crypto exchanges__

## Installation

`go get github.com/llonchj/go-finnhub`

## Usage

First sign up for your api token here [finnhub.io](https://finnhub.io/)

Follow this basic example, for more in-depth documentation see the [docs](https://godoc.org/github.com/llonchj/go-finnhub):

```go
c := client.New("your_token_here")

// Stocks
company, err := c.Stock.GetProfile("AAPL")
ceo, err := c.Stock.GetCEO("AAPL")
recommendation, err := c.Stock.GetRecommendations("AAPL")
target, err := c.Stock.GetPriceTarget("AAPL")
options, err := c.Stock.GetOptionChain("DBD")
peers, err := c.Stock.GetPeers("AAPL")
earnings, err := c.Stock.GetEarnings("AAPL")
candle, err := c.Stock.GetCandle("AAPL", finnhub.CandleResolutionDay, nil)
exchanges, err := c.Stock.GetExchanges()
symbols, err := c.Stock.GetSymbols("US")
gradings, err := c.Stock.GetGradings(&finnhub.GradingParams{Symbol: "AAPL"})

// Crypto
exchanges, err := c.Crypto.GetExchanges()
symbols, err := c.Crypto.GetSymbols("Binance")
candle, err := c.Crypto.GetCandle("BINANCE:BEAMUSDT", finnhub.CandleResolutionMonth, nil)

// Forex
exchanges, err := c.Forex.GetExchanges()
symbols, err := c.Forex.GetSymbols("oanda")
candle, err := c.Forex.GetCandle("OANDA:XAU_GBP", finnhub.CandleResolutionMonth, nil)

// News
news, err := c.News.Get(nil)
news, err = c.News.Get(&finnhub.NewsParams{Category: finnhub.NewsCategoryCrypto})
news, err = c.News.GetCompany("AAPL")
sentiment, err := c.News.GetSentiment("AAPL")
```
