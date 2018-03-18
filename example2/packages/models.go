package mypkg

type ReceivedData struct {
	ID int `json:"id,omitempty"`
}

type TransformedData struct {
	UserID       int    `json:"user_id,omitempty"`
	SomeNewField string `json:"some_new_field"`
}
