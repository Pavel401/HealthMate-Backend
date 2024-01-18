package model

type SharedProfile struct {
	BasicInfo Basic_Info `json:"basic_info" `
	Visits    []*Visit   `json:"visit" `
}
