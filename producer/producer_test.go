package producer_test

import (
	"fmt"
	"gokafkaexample/producer"
	"testing"

	"github.com/Shopify/sarama/mocks"
)

func TestSendMessage(t *testing.T) {
	t.Run("Send message OK", func(t *testing.T) {
		//create producer mock
		mockedProducer := mocks.NewSyncProducer(t, nil)

		//create success expected producer or success send message
		mockedProducer.ExpectSendMessageAndSucceed()
		kafka := &producer.KafkaProducer{
			Producer: mockedProducer,
		}

		msg := "Message 1"

		err := kafka.SendMessage("test_topic", msg)
		if err != nil {
			t.Errorf("Send message should not be error but have: %v", err)
		}
	})

	t.Run("Send message NOK", func(t *testing.T) {
		mockedProducer := mocks.NewSyncProducer(t, nil)

		//create product failed to send a message
		mockedProducer.ExpectSendMessageAndFail(fmt.Errorf("Error"))
		kafka := &producer.KafkaProducer{
			Producer: mockedProducer,
		}

		msg := "Message 1"

		err := kafka.SendMessage("test_topic", msg)
		if err == nil {
			t.Error("this should be error")
		}
	})
}
