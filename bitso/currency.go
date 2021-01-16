package bitso

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// Currency represents currencies
type Currency uint8

// Currencies
const (
	CurrencyNone Currency = iota

	ARS
	BAT
	BCH
	BTC
	DAI
	ETH
	GNT
	LTC
	MANA
	MXN
	TUSD
	USD
	XRP
	BRL
)

var currencyNames = map[Currency]string{
	ARS:  "ars",
	BAT:  "bat",
	BCH:  "bch",
	BTC:  "btc",
	DAI:  "dai",
	ETH:  "eth",
	GNT:  "gnt",
	LTC:  "ltc",
	MANA: "mana",
	MXN:  "mxn",
	TUSD: "tusd",
	USD:  "usd",
	XRP:  "xrp",
	BRL:  "brl",
}

func getCurrencyByName(name string) (*Currency, error) {
	for c, n := range currencyNames {
		if n == name {
			return &c, nil
		}
	}
	return nil, fmt.Errorf("no such currency: %q", name)
}

// MarshalJSON implements json.Marshaler
func (c Currency) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

func (c *Currency) fromString(z string) error {
	for k, v := range currencyNames {
		if v == z {
			*c = k
			return nil
		}
	}
	return fmt.Errorf("unsupported currency: %v", z)
}

// UnmarshalJSON implements json.Unmarshaler
func (c *Currency) UnmarshalJSON(in []byte) error {
	var z string
	if err := json.Unmarshal(in, &z); err != nil {
		return err
	}
	return c.fromString(z)
}

func (c Currency) String() string {
	if z, ok := currencyNames[c]; ok {
		return z
	}
	panic(fmt.Sprintf("unsupported currency: %q", string(c)))
}

func (c Currency) Value() (driver.Value, error) {
	return c.String(), nil
}

func (c *Currency) Scan(value interface{}) error {
	return c.fromString(value.(string))
}
