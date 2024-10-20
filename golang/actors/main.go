package main

import (
	"log"
	"time"
)

const (
	NetworkActor = "network"
	ProcessorActor = "processor"
	SenderActor = "sender"
)

type MessageProcessor struct {}

func (p *MessageProcessor) Execute(income Message) {
	log.Printf("received message from [%s:%s]: %s", income.From, income.To, income.Body)

	outcome := Message{
		From: income.To,
		To: SenderActor,
		Body: "processed_message",
	}

	if err := GetActorManager().SendMessage(outcome); err != nil {
		log.Printf("failed to send message: %s", err.Error())
	}
}

type MessageSender struct {}

func (s *MessageSender) Execute(income Message) {
	log.Printf("received message [%s:%s]: %s", income.From, income.To, income.Body)
	log.Printf("message successfullly sent")
}

func main() {
	manager := GetActorManager()
	if err := manager.CreateActor(ProcessorActor, &MessageProcessor{}); err != nil {
		panic("failed to create actor")
	}

	if err := manager.CreateActor(SenderActor, &MessageSender{}); err != nil {
		panic("failed to create actor")
	}

	message := Message{
		From: NetworkActor,
		To: ProcessorActor,
		Body: "received_message",
	}

	if err := manager.SendMessage(message); err != nil {
		log.Printf("failed to send message: %s", err.Error())
	}

	time.Sleep(time.Second)
}
