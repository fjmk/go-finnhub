package client_test

import (
	"reflect"
	"testing"

	"github.com/llonchj/go-finnhub/client"
	"github.com/llonchj/go-finnhub/crypto"
	"github.com/llonchj/go-finnhub/forex"
	"github.com/llonchj/go-finnhub/news"
	"github.com/llonchj/go-finnhub/stock"
)

func TestNew(t *testing.T) {
	a := client.NewAPI("token", client.Version)
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want *client.Client
	}{
		{
			name: "valid",
			args: args{key: "token"},
			want: &client.Client{
				Stock:  stock.Client{API: a},
				Forex:  forex.Client{API: a},
				Crypto: crypto.Client{API: a},
				News:   news.Client{API: a},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := client.New(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
