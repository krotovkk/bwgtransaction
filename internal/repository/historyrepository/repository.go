package historyrepository

import (
	"strconv"

	"github.com/Shopify/sarama"
)

type HistoryRepository struct {
	producer sarama.SyncProducer
}

func NewHistoryRepository(producer sarama.SyncProducer) *HistoryRepository {
	return &HistoryRepository{
		producer: producer,
	}
}

func (r *HistoryRepository) Save(clientId int, diff float64) error {
	_, _, err := r.producer.SendMessage(&sarama.ProducerMessage{
		Topic: "ChangeBalance",
		Key:   sarama.StringEncoder(clientId),
		Value: sarama.StringEncoder(strconv.FormatFloat(diff, 'f', -1, 64)),
	})

	return err
}
