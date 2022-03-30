/*
Author:ydy
Date:
Desc:
*/
package queue

import (
	"fmt"
	"github.com/bingtianyiyan/youki_gotools/commonexternal/rescue"
	"github.com/bingtianyiyan/youki_gotools/commonexternal/threading"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

const queueName = "queue"

type (
	// A Queue is a message queue.
	Queue struct {
		name                 string
		//metrics              *stat.Metrics
		producerFactory      ProducerFactory
		producerRoutineGroup *threading.RoutineGroup
		consumerFactory      ConsumerFactory
		consumerRoutineGroup *threading.RoutineGroup
		producerCount        int
		consumerCount        int
		active               int32
		channel              chan string
		quit                 chan struct{}
		listeners            []Listener
		eventLock            sync.Mutex
		eventChannels        []chan interface{}
	}

	// A Listener interface represents a listener that can be notified with queue events.
	Listener interface {
		OnPause()
		OnResume()
	}

	// A Poller interface wraps the method Poll.
	Poller interface {
		Name() string
		Poll() string
	}

	// A Pusher interface wraps the method Push.
	Pusher interface {
		Name() string
		Push(string) error
	}
)

func NewQueue(producerFactory ProducerFactory, consumerFactory ConsumerFactory) *Queue{
	return &Queue{
		name: queueName,
		producerFactory: producerFactory,
		producerRoutineGroup: threading.NewRoutineGroup(),
		consumerFactory: consumerFactory,
		consumerRoutineGroup: threading.NewRoutineGroup(),
        producerCount: runtime.NumCPU(),
        consumerCount: runtime.NumCPU() << 1,
        channel: make(chan string),
        quit: make(chan struct{}),
	}
}

func (m *Queue) SetName(name string){
	m.name = name
}

// SetNumConsumer sets the number of consumers.
func (m *Queue) SetNumConsumer(count int) {
	m.consumerCount = count
}

// SetNumProducer sets the number of producers.
func (m *Queue) SetNumProducer(count int) {
	m.producerCount = count
}


func (m *Queue) AddListener(listener Listener){
	m.listeners = append(m.listeners,listener)
}

func (m *Queue) BroadEventCast(message interface{}){
   go func() {
   	   m.eventLock.Lock()
   	   defer m.eventLock.Unlock()
   	   for _,eventchannel := range m.eventChannels{
   	   	  eventchannel <- message
	   }
   }()
}

func (m *Queue) Start(){
    //producer
	 m.startProducers(m.producerCount)
	//consumer
	 m.startConsumers(m.consumerCount)

    m.producerRoutineGroup.Wait()
	close(m.channel)//
	m.consumerRoutineGroup.Wait()
}

func (m *Queue) startProducers(pdcount int){
  for i:=0;i<pdcount;i++{
  	m.producerRoutineGroup.Run(
		func() {
			m.produce()
		},
	)
  }
}

func (m *Queue) produce(){
	var producer Producer

	//create
	for{
		var err error
		if producer,err = m.producerFactory();err != nil {
			//fmt.Println("Error on creating producer")
			time.Sleep(time.Second)
		}else {
			break
		}
	}

	atomic.AddInt32(&m.active,1)
	//listener
	producer.AddListener(routineListener{
		queue: m,
	})

	//producter data
	for {
		select {
		case <-m.quit:
			//
			return
		default:
			if v, ok := m.produceOne(producer); ok {
				fmt.Println(v)
				m.channel <- v
			}
		}
	}
}

func (m *Queue) produceOne(producer Producer) (string, bool) {
	// avoid panic quit the producer, just log it and continue
	defer rescue.Recover()

	return producer.Produce()
}

func (m *Queue) startConsumers(number int) {
	for i := 0; i < number; i++ {
		eventChan := make(chan interface{})
		m.eventLock.Lock()
		m.eventChannels = append(m.eventChannels, eventChan)
		m.eventLock.Unlock()
		m.consumerRoutineGroup.Run(func() {
			m.consume(eventChan)
		})
	}
}


func (m *Queue) consume(eventChan chan interface{}) {
	var consumer Consumer

	for {
		var err error
		if consumer, err = m.consumerFactory(); err != nil {
			fmt.Sprintf("Error on creating consumer: %v", err)
			time.Sleep(time.Second)
		} else {
			break
		}
	}

	for {
		select {
		case message, ok := <-m.channel:
			if ok {
				m.consumerOne(consumer, message)
			} else {
				fmt.Println("Task channel was closed, quitting consumer...")
				return
			}
		case event := <-eventChan:
			consumer.OnEvent(event)
		}
	}
}

func (m *Queue) consumerOne(consumer Consumer,msg string){
	threading.RunSafe(func() {
		if err := consumer.Consume(msg); err != nil {
			fmt.Println("Error occurred while consuming")
		}
	})
}

func (m *Queue) Stop(){
  close(m.quit)
}

func (m *Queue) pause(){
   for _,listener := range m.listeners{
   	   listener.OnPause()
   }
}

func (m *Queue) resume(){
  for _,listener := range m.listeners{
  	   listener.OnResume()
  }
}


type routineListener struct {
	queue *Queue
}

func (rl routineListener) OnProducerPause() {
	if atomic.AddInt32(&rl.queue.active, -1) <= 0 {
		rl.queue.pause()
	}
}

func (rl routineListener) OnProducerResume() {
	if atomic.AddInt32(&rl.queue.active, 1) == 1 {
		rl.queue.resume()
	}
}


