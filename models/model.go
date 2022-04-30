package models

type TextInput struct {
	Text string `json:"text"`
}

type WordFrequency struct {
	Key   string `json:"word"`
	Count int    `json:"count"`
}
