package talking

import (
	"context"
	"log/slog"
	"os"
	"github.com/segmentio/kafka-go"
)

func GetMessage(topic string, ch chan string, chfe chan error){
	loger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   topic,
})

	defer func() {
		if err := reader.Close(); err != nil {
			loger.Error("failed to close reader: " + err.Error())
			chfe <-err
		}
	}()

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			loger.Error("failed to read message: " + err.Error())
			chfe <- err
			continue
		}
		ch <- string(msg.Value)
		loger.Info("Получено сообщение:", string(msg.Value), nil)
	}
}
