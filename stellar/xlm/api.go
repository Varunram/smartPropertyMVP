package xlm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	clients "github.com/stellar/go/clients/horizon"
	protocols "github.com/stellar/go/protocols/horizon"
)

// xlm is a set of functions that interface with stellar-core without needing the
// horizon API that stellar provides, which is incomplete

var TestNetClient = &clients.Client{
	URL:  "https://horizon-testnet.stellar.org",
	HTTP: http.DefaultClient,
}

func GetStateHash(a string) (string, error) {
	var res string
	var err error

	_, err = GetBlockHash("1000000")
	return res, err
}

/*
type Ledger struct {
    Links struct {
        Self         hal.Link `json:"self"`
        Transactions hal.Link `json:"transactions"`
        Operations   hal.Link `json:"operations"`
        Payments     hal.Link `json:"payments"`
        Effects      hal.Link `json:"effects"`
    }   `json:"_links"`
    ID               string    `json:"id"`
    PT               string    `json:"paging_token"`
    Hash             string    `json:"hash"`
    PrevHash         string    `json:"prev_hash,omitempty"`
    Sequence         int32     `json:"sequence"`
    TransactionCount int32     `json:"transaction_count"`
    OperationCount   int32     `json:"operation_count"`
    ClosedAt         time.Time `json:"closed_at"`
    TotalCoins       string    `json:"total_coins"`
    FeePool          string    `json:"fee_pool"`
    BaseFee          int32     `json:"base_fee_in_stroops"`
    BaseReserve      int32     `json:"base_reserve_in_stroops"`
    MaxTxSetSize     int32     `json:"max_tx_set_size"`
    ProtocolVersion  int32     `json:"protocol_version"`
    HeaderXDR        string    `json:"header_xdr"`
}

*/
// we are needing to have to call this API because the go sdk for stellar is
// incomplete. Once we run a local node / node server, we can call our server's
// endpoint here
func GetLedgerData(blockNumber string) ([]byte, error) {
	var err error
	var data []byte
	resp, err := http.Get(TestNetClient.URL + "/ledgers/" + blockNumber)
	if err != nil {
		return data, err
	}
	if resp.Status != "200 OK" {
		return data, fmt.Errorf("API Request did not succeed")
	}
	defer resp.Body.Close()
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return data, err
	}
	return data, err
}

func GetBlockHash(blockNumber string) (string, error) {
	var err error
	var hash string
	b, err := GetLedgerData(blockNumber)
	if err != nil {
		return hash, err
	}
	var x protocols.Ledger
	err = json.Unmarshal(b, &x)
	if err != nil {
		return hash, err
	}
	hash = x.Hash
	log.Printf("The block hash for block %d is: %s and the prev hash is %s", x.Sequence, hash, x.PrevHash)
	return hash, err
}
