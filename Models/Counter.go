package Models

type Counter struct {
	ID             string  `json:"_id,omitempty"`
	Sequence_Value string `json:"sequence_value,omitempty"`
}
