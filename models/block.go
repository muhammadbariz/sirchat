package models

import (
	"encoding/json"
	"errors"
)

type MessageBlockType string

const (
	MBTText      MessageBlockType = "text"
	MBTTable     MessageBlockType = "table"
	MBTContainer MessageBlockType = "container"
	MBTImage     MessageBlockType = "image"
)

type IBlock interface {
	BlockType() MessageBlockType
	Validate() bool
}

type Block struct {
	Type MessageBlockType `json:"type"`
}

// BlockType returns the type of the block
func (ths *Block) BlockType() MessageBlockType {
	return ths.Type
}

func Compose(block IBlock) ([]byte, error) {
	if !block.Validate() {
		return nil, errors.New("invalid block")
	}

	res, err := json.Marshal(block)
	if err != nil {
		return nil, errors.New("error when marshaling block")
	}

	return res, nil
}

func NewBlock() *Block {
	return &Block{}
}
