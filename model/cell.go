package model

import (
	"encoding/json"
	"errors"
)

// Cell represents a single cell in the Board, which can be either empty or hold a value (0-9)
type Cell struct {
	Value *int
}

// UnmarshalJSON customizes the unmarshalling of a Cell from JSON
func (c *Cell) UnmarshalJSON(data []byte) error {
	// If the data is "null", set the cell as empty
	if string(data) == "null" {
		c.Value = nil
		return nil
	}

	// Otherwise, unmarshal the number into the cell's value
	var v int
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	// Set the cell's value
	c.Value = &v
	return nil
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
