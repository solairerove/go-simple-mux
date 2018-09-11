package entity

import "github.com/solairerove/go-simple-mux/pkg/entity"

// Person ... tbd
type Person struct {
	ID        string          `json:"id,omitempty"`
	Firstname string          `json:"firstname,omitempty"`
	Lastname  string          `json:"lastname,omitempty"`
	Address   *entity.Address `json:"address,omitempty"`
}
