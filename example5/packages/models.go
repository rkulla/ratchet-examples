package mypkg

type ReceivedData struct {
	ID   int    `json:"id,omitempty"`
	City string `json:"city"`
}

type TransformedData struct {
	UserID int    `json:"user_id,omitempty"`
	City   string `json:"city"`
}

type UserID struct {
	UserID int `json:"id,omitempty"`
}
