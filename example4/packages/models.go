package mypkg

const (
	Query1Helper = "query1"
	Query2Helper = "query2"
)

type ReceivedData1 struct {
	TypeHelper string `json:"type_helper,omitempty"`
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
}

type ReceivedData2 struct {
	TypeHelper string `json:"type_helper,omitempty"`
	ID         int    `json:"id,omitempty"`
	City       string `json:"city,omitempty"`
	State      string `json:"state,omitempty"`
}

type TransformedData struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	City   string `json:"city"`
}
