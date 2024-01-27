package server

import (
	"log"

	"github.com/Sskrill/mq-log/internal/config"
	"github.com/streadway/amqp"
)

func main() {
	cfg, err := config.NewCfg()
	if err != nil {
		log.Fatal(err)
	}

	conn, err := amqp.Dial(cfg.URI)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	que, err := ch.QueueDeclare("log", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	messages, err := ch.Consume(que.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	forever := make(chan bool)

	go func() {
		for msg := range messages {
			log.Printf("Message : %s", msg.Body)
		}
	}()
	log.Print("CTRL + C to exit")
	<-forever

}
