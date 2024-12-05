package kafkaimpl

type Config struct {
	Address   string `json:"address"`
	Partition int    `json:"partition"`
}
