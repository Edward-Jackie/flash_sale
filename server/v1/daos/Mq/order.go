package daos

import (
	"Flash_Sale/driver"
	"github.com/streadway/amqp"
)

func OrderQueue(n string) (*amqp.Channel, *amqp.Queue, error) {
	if driver.MQ.IsClosed() {
		driver.InitMQ()
	}

	ch, err := driver.MQ.Channel()
	//defer ch.Close()
	if err != nil {
		return nil, nil, err
	}

	order, err := ch.QueueDeclare(
		n,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, nil, err
	}
	return ch, &order, nil
}
