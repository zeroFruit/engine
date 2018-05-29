package repository

import (
	"log"

	"github.com/it-chain/leveldb-wrapper/key_value_db"
	"github.com/it-chain/yggdrasill"
	"github.com/it-chain/yggdrasill/impl"
	"github.com/it-chain/it-chain-Engine/blockchain"
)

type BlockRepository struct {
	yggdrasill *yggdrasill.Yggdrasill
}

func NewBlockRepository(keyValueDB key_value_db.KeyValueDB, validator impl.DefaultValidator, opts map[string]interface{}) *BlockRepository {
	ygg, err := yggdrasill.NewYggdrasill(keyValueDB, &validator, opts)

	if err != nil {
		log.Fatal(err.Error())
	}

	return &BlockRepository{
		yggdrasill: ygg,
	}
}

func (br BlockRepository) Close() {
	br.yggdrasill.Close()
}

func (br BlockRepository) AddBlock(block impl.DefaultBlock) error {
	err := br.yggdrasill.AddBlock(&block)
	if err != nil {
		return err
	}
	return nil
}

// Issue : func GetBlock~
// 현재 repository가 it-chain-engine에 속하기 때문에 retrievedBlock의 type을 it-chain-Engine/blockchain/domain/model/block.Block으로 하고 싶지만,
// 그렇게 되면 common.Block 인터페이스에 속하지 않는 것 같음. block.Block에 필요한 메소드를 다시 구현하기보다는 현재는 impl.Defaultblock 사용중(주석작성자 GitID:junk-sound).

func (br BlockRepository) GetBlockByHeight(blockHeight uint64) (*impl.DefaultBlock, error) {
	var retrievedBlock impl.DefaultBlock
	err := br.yggdrasill.GetBlockByHeight(&retrievedBlock, blockHeight)
	if err != nil {
		return nil, err
	}
	return &retrievedBlock, err
}

func (br BlockRepository) GetBlockBySeal(seal []byte) (*impl.DefaultBlock, error) {
	var retrievedBlock impl.DefaultBlock
	err := br.yggdrasill.GetBlockBySeal(&retrievedBlock, seal)
	if err != nil {
		return nil, err
	}
	return &retrievedBlock, err
}

func (br BlockRepository) GetLastBlock() (*impl.DefaultBlock, error) {
	var retrievedBlock impl.DefaultBlock
	err := br.yggdrasill.GetLastBlock(&retrievedBlock)
	if err != nil {
		return nil, err
	}
	return &retrievedBlock, err
}

func (br BlockRepository) GetTransactionByTxID(txid string) (*blockchain.Transaction, error) {
	var retrievedTx impl.DefaultTransaction
	err := br.yggdrasill.GetTransactionByTxID(&retrievedTx, txid)
	if err != nil {
		return nil, err
	}
	return &retrievedTx, err
}

func (br BlockRepository) GetBlockByTxID(txid string) (*impl.DefaultBlock, error) {
	var retrievedBlock impl.DefaultBlock
	err := br.yggdrasill.GetBlockByTxID(&retrievedBlock, txid)
	if err != nil {
		return nil, err
	}
	return &retrievedBlock, err
}