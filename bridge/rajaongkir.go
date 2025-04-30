package bridge

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"healthcare-allopurinol/entity"
	"io"
	"net/http"
	netUrl "net/url"
	"strconv"
	"strings"
)

type ROBridgeItf interface {
	GetDataWithPostal(ctx context.Context, postalCode string) (*entity.RajaOngkir, error)
	GetShippingCost(ctx context.Context, from, to int64, weight decimal.Decimal, courier string) (*entity.RajaOngkirCost, error)
}

type ROBridgeImpl struct {
	baseUrl string
	apiKey  string
}

func NewROBridge(baseUrl, apiKey string) ROBridgeItf {
	return &ROBridgeImpl{
		baseUrl: baseUrl,
		apiKey:  apiKey,
	}
}

func (b *ROBridgeImpl) GetDataWithPostal(ctx context.Context, postalCode string) (*entity.RajaOngkir, error) {
	url := fmt.Sprintf("%s/destination/domestic-destination?search=%s&limit=1", b.baseUrl, postalCode)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("key", b.apiKey)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var ro *entity.RajaOngkir
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&ro)
	if err != nil {
		return nil, err
	}

	if ro.Meta.Code != 200 {
		return nil, errors.New(fmt.Sprintf("RajaOngkir: %s", ro.Meta.Message))
	}
	if len(ro.Data) == 0 {
		return nil, errors.New("RajaOngkir: No data")
	}

	return ro, nil
}

func (b *ROBridgeImpl) GetShippingCost(ctx context.Context, from, to int64, weight decimal.Decimal, courier string) (*entity.RajaOngkirCost, error) {
	url := fmt.Sprintf("%s/calculate/domestic-cost", b.baseUrl)

	form := netUrl.Values{}
	form.Set("origin", strconv.FormatInt(from, 10))
	form.Set("destination", strconv.FormatInt(to, 10))
	form.Set("weight", weight.String())
	form.Set("courier", courier)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("key", b.apiKey)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var ro *entity.RajaOngkirCost
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&ro)
	if err != nil {
		return nil, err
	}

	if ro.Meta.Code != 200 {
		return nil, errors.New(fmt.Sprintf("RajaOngkir: %s", ro.Meta.Message))
	}
	if len(ro.Data) == 0 {
		return nil, errors.New("RajaOngkir: No data")
	}

	return ro, nil
}
