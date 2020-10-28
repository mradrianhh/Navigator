package models

// Screen ...
type Screen interface {
	Show(state *State) error
}

// Empty ...
type Empty struct {
}

// Show ...
func (empty Empty) Show(state *State) error {
	return nil
}
