package model

import "errors"

// Cell represents a single cell in the Board, which can be either empty or hold a value (1-9)
type Cell struct {
	Value *int `json:"value"` // Use JSON tags to define how this field will be represented in JSON
}

// IsEmpty checks if the cell is empty
func (c *Cell) IsEmpty() bool {
	return c.Value == nil
}

// SetValue sets the value of the cell
func (c *Cell) SetValue(v int) error {
	if v < 1 || v > 9 {
		return errors.New("value must be between 1 and 9")
	}
	c.Value = &v
	return nil
}

// Clear clears the cell, making it empty
func (c *Cell) Clear() {
	c.Value = nil
}

// GetValue returns the value of the cell or an error if the cell is empty
func (c *Cell) GetValue() (int, error) {
	if c.IsEmpty() {
		return 0, errors.New("cell is empty")
	}
	return *c.Value, nil
}
