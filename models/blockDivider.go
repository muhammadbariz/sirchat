package models

import "errors"

// DividerBlock is a subtype of block. It represents a divider block. It
// will render as a divider line between blocks.
type dividerBlock struct {
	block
}

// Validate performs validation to the DividerBlock.
func (ths *dividerBlock) Validate() (bool, []error) {
	// DividerBlock validation implementation
	var errs []error
	if ths.Type != MBTDivider {
		errs = append(errs, errors.New("invalid container block type"))
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

// NewDividerBlock returns a new instance of a divider block to be rendered
func NewDividerBlock() *dividerBlock {
	block := dividerBlock{}
	block.Type = MBTDivider

	return &block
}
