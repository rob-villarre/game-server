package udp

type Message struct {
	Type      string      `json:"type"`
	Player    interface{} `json:"player,omitempty"`
	Timestamp int64       `json:"timestamp,omitempty"`
}
