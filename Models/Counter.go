package Models

type Counter struct {
	ID             string `json:"_id,omitempty"`
	Sequence_Value float32 `json:"sequence_value,omitempty"`
}
