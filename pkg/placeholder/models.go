package placeholder

import (
	"encoding/json"
	"io"
)

type TODO struct {
	UserID    int    `json:"userId,omitempty"`
	ID        int    `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Completed bool   `json:"completed"`
}

func ToDoFromJSON(data io.Reader) (*TODO, error) {
	var td TODO
	if err := json.NewDecoder(data).Decode(&td); err != nil {
		return nil, err
	}
	return &td, nil
}

func (td *TODO) ToJSON() string {
	b, _ := json.Marshal(td)
	return string(b)
}

func (td *TODO) Sanitise() {
	td.UserID = 0
}

func (td *TODO) IsValid() bool {
	return td.Title != "" && td.UserID != 0
}

func TODOSFromJSON(data io.Reader) ([]*TODO, error) {
	var td []*TODO
	if err := json.NewDecoder(data).Decode(&td); err != nil {
		return nil, err
	}
	return td, nil
}
