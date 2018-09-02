package oip042

import (
	"github.com/azer/logger"
	"github.com/bitspill/oip/datastore"
	"github.com/json-iterator/go"
	"gopkg.in/olivere/elastic.v6"
)

func on42JsonRegisterPool(any jsoniter.Any, tx datastore.TransactionData) {
	t := log.Timer()
	defer t.End("on42JsonRegisterPool", logger.Attrs{"txid": tx.Transaction.Txid})

	sig := any.Get("signature").ToString()

	var el elasticOip042Pool
	el.Pool = any.GetInterface()
	el.Meta = AMeta{
		Time:        tx.Transaction.Time,
		Txid:        tx.Transaction.Txid,
		Signature:   sig,
		BlockHash:   tx.BlockHash,
		Block:       tx.Block,
		Deactivated: false,
		Tx:          tx,
		Type:        "oip042",
	}

	bir := elastic.NewBulkIndexRequest().Index(oip042PoolIndex).Type("_doc").Id(tx.Transaction.Txid).Doc(el)
	datastore.AutoBulk.Add(bir)
}

func on42JsonEditPool(any jsoniter.Any, tx datastore.TransactionData) {
	t := log.Timer()
	defer t.End("on42JsonEditPool", logger.Attrs{"txid": tx.Transaction.Txid})

	sig := any.Get("signature").ToString()

	var el elasticOip042Edit
	el.Edit = any.GetInterface()
	el.Meta = OMeta{
		Time:      tx.Transaction.Time,
		Txid:      tx.Transaction.Txid,
		Signature: sig,
		BlockHash: tx.BlockHash,
		Block:     tx.Block,
		Completed: false,
		Tx:        tx,
		Type:      "pool",
	}

	bir := elastic.NewBulkIndexRequest().Index(oip042EditIndex).Type("_doc").Id(tx.Transaction.Txid).Doc(el)
	datastore.AutoBulk.Add(bir)
}

func on42JsonTransferPool(any jsoniter.Any, tx datastore.TransactionData) {
	t := log.Timer()
	defer t.End("on42JsonTransferPool", logger.Attrs{"txid": tx.Transaction.Txid})

	sig := any.Get("signature").ToString()

	var el elasticOip042Transfer
	el.Transfer = any.GetInterface()
	el.Meta = OMeta{
		Time:      tx.Transaction.Time,
		Txid:      tx.Transaction.Txid,
		Signature: sig,
		BlockHash: tx.BlockHash,
		Block:     tx.Block,
		Completed: false,
		Tx:        tx,
		Type:      "pool",
	}

	bir := elastic.NewBulkIndexRequest().Index(oip042TransferIndex).Type("_doc").Id(tx.Transaction.Txid).Doc(el)
	datastore.AutoBulk.Add(bir)
}

func on42JsonDeactivatePool(any jsoniter.Any, tx datastore.TransactionData) {
	t := log.Timer()
	defer t.End("on42JsonDeactivatePool", logger.Attrs{"txid": tx.Transaction.Txid})

	sig := any.Get("signature").ToString()

	var el elasticOip042DeactivateInterface
	el.Deactivate = any.GetInterface()
	el.Meta = OMeta{
		Time:      tx.Transaction.Time,
		Txid:      tx.Transaction.Txid,
		Signature: sig,
		BlockHash: tx.BlockHash,
		Block:     tx.Block,
		Completed: false,
		Tx:        tx,
		Type:      "pool",
	}

	bir := elastic.NewBulkIndexRequest().Index(oip042DeactivateIndex).Type("_doc").Id(tx.Transaction.Txid).Doc(el)
	datastore.AutoBulk.Add(bir)
}
