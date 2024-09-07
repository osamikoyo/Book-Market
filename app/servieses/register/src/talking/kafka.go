package talking

import (
	"context"
	"log/slog"
	"os"

	"github.com/segmentio/kafka-go"
)


func consumeMessages(topic string) {
	loger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   topic,
	})

	defer reader.Close()

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			loger.Error(err.Error())
		}
		loger.Info("Получено сообщение: %s\n", string(msg.Value), nil)
	}
}