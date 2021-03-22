package siacentral

import (
	"fmt"
	"net/http"
)

type (
	siaAPI struct {
	}
)

// FindAddressBalance gets the balance and transactions of the addresses
func (s *siaAPI) FindAddressBalance(limit int, page int, addresses []string) (WalletBalance, error) {
	var resp transactionsResp

	u := fmt.Sprintf("https://api.siacentral.com/v2/wallet/addresses?limit=%d&page=%d", limit, page)
	code, err := makeAPIRequest(http.MethodPost, u, map[string]interface{}{
		"addresses": addresses,
	}, &resp)

	if err != nil {
		return WalletBalance{}, err
	}

	if code < 200 || code >= 300 || resp.Type != "success" {
		return WalletBalance{}, fmt.Errorf(resp.Message)
	}

	return resp.WalletBalance, nil
}

// FindUsedAddresses returns an array of all used addresses
func (s *siaAPI) FindUsedAddresses(addresses []string) ([]AddressUsage, error) {
	var resp addressesResp

	code, err := makeAPIRequest(http.MethodPost, "https://api.siacentral.com/v2/wallet/addresses/used", map[string]interface{}{
		"addresses": addresses,
	}, &resp)

	if err != nil {
		return nil, err
	}

	if code < 200 || code >= 300 || resp.Type != "success" {
		return nil, fmt.Errorf(resp.Message)
	}

	return resp.Addresses, nil
}

func (s *siaAPI) GetBlockHeight() (uint64, error) {
	var resp blockResp

	code, err := makeAPIRequest(http.MethodGet, "https://api.siacentral.com/v2/explorer/blocks", nil, &resp)
	if err != nil {
		return 0, err
	}

	if code < 200 || code >= 300 || resp.Type != "success" {
		return 0, fmt.Errorf(resp.Message)
	}

	return resp.Block.Height, nil
}

// NewSiaAPI creates a new Sia client to access the Sia api
func NewSiaAPI() API {
	return new(siaAPI)
}
