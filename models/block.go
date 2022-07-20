package models

type MessageBlockType string

const (
	MBTText      MessageBlockType = "text"
	MBTTable     MessageBlockType = "table"
	MBTContainer MessageBlockType = "container"
	MBTImage     MessageBlockType = "image"
	MBTInput     MessageBlockType = "input"
	MBTButton    MessageBlockType = "button"
)

type IBlock interface {
	BlockType() MessageBlockType
	Validate() error
}

type Block struct {
	Type MessageBlockType `json:"type"`
}

// BlockType returns the type of the block
func (ths *Block) BlockType() MessageBlockType {
	return ths.Type
}

func validateBlock(block IBlock) error {
	if err := block.Validate(); err != nil {
		return err
	}

	return nil
}

func NewBlocks(blocks ...IBlock) []IBlock {
	return blocks
}
