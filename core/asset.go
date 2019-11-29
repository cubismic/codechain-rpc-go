package core

import (
	"encoding/json"

	"github.com/CodeChain-io/codechain-rpc-go/primitives"
)

type AssetJSON struct {
	AssetType              string   `json:"assetType"`
	LockScriptHash         string   `json:"lockScriptHash"`
	Parameters             []string `json:"parameters"`
	Quantity               string   `json:"quantity"`
	OrderHash              string   `json:"orderHash,omitempty"`
	ShardID                int      `json:"shardId"`
	Tracker                string   `json:"tracker"`
	TransactionOutputIndex int      `json:"transactionOutputIndex"`
}

type AssetData struct {
	AssetType              primitives.H160
	ShardID                int
	LockScriptHash         primitives.H160
	Parameters             []string
	Quantity               primitives.U64
	OrderHash              primitives.H256
	Tracker                primitives.H256
	TransactionOutputIndex int
}

type Asset struct {
	AssetType      primitives.H160
	ShardID        int
	LockScriptHash primitives.H160
	Parameters     []string
	Quantity       primitives.U64
	OrderHash      primitives.H256
	OutPoint       interface{} // AssetOutPoint
}

func NewAsset(data AssetData) (Asset, error) {
	return Asset{
			data.AssetType,
			data.ShardID,
			data.LockScriptHash,
			data.Parameters,
			data.Quantity,
			data.OrderHash,
			nil}, // TODO AssetOutPoint
		nil
}

func AssetFromJSON(data AssetJSON) (Asset, error) {
	assetType, _ := primitives.StringToH160(data.AssetType)
	lockScriptHash, _ := primitives.StringToH160(data.LockScriptHash)
	quantity := primitives.NewU64(data.Quantity)
	orderHash, _ := primitives.StringToH256(data.OrderHash)
	tracker, _ := primitives.StringToH256(data.Tracker)
	return NewAsset(AssetData{
		assetType,
		data.ShardID,
		lockScriptHash,
		data.Parameters,
		quantity,
		orderHash,
		tracker,
		data.TransactionOutputIndex})
}

func (a Asset) ToJSON() (string, error) {
	b, err := json.Marshal(a)
	return string(b), err
}

// TODO func (a Asset) createTransferInput()
// TODO func (a Asset) createTransferTransaction()
