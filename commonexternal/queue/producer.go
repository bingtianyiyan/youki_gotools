/*
Author:ydy
Date:
Desc:
*/
package queue

type (
	Producer interface {
		Produce()(string,bool)
		AddListener(listener ProducerListener)
	}

	ProducerListener interface {
		OnProducerPause()
		OnProducerResume()
	}

	ProducerFactory func()(Producer,error)
)
