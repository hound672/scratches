package main

import (
	"log"
	"sync"
)

type TopicName string

type Subscriber interface {
	Notify(topic TopicName)
}

type Publisher struct {
	subscribers map[TopicName][]Subscriber
	mutex       sync.RWMutex
}

func NewPublisher() *Publisher {
	return &Publisher{
		subscribers: make(map[TopicName][]Subscriber),
	}
}

func (p *Publisher) Subscribe(topic TopicName, subscriber Subscriber) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.subscribers[topic] = append(p.subscribers[topic], subscriber)
}

func (p *Publisher) Unsubscribe(topic TopicName, subscriber Subscriber) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	subscribers, ok := p.subscribers[topic]
	if !ok {
		log.Printf("there is no any subscribers")
		return
	}

	for i, s := range subscribers {
		if s == subscriber {
			log.Printf("MATCH")
			subscribers = append(subscribers[:i], subscribers[i+1:]...)
		}
	}
}

func (p *Publisher) Notify(topic TopicName) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	subscribers, ok := p.subscribers[topic]
	if !ok {
		log.Printf("there is no any subscribers")
		return
	}

	for _, subscriber := range subscribers {
		subscriber.Notify(topic)
	}
}

type OwnSubscriber struct {
	
}

func (s *OwnSubscriber) Notify(topic TopicName) {
	log.Printf("got notification: %s", topic)
}

func main() {
	log.Printf("PubSub start")
	topic := TopicName("test_topic")

	ownSubs := &OwnSubscriber{}
	ownSubs2 := &OwnSubscriber{}
	_ = ownSubs
	publisher := NewPublisher()
	
	publisher.Subscribe(topic, ownSubs)
	publisher.Subscribe(topic, ownSubs2)
	
	publisher.Notify(topic)

	log.Printf("---------------")
	publisher.Unsubscribe(topic, ownSubs2)

	publisher.Notify(topic)
}
