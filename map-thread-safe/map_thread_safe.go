package map_thread_safe

import (
	"fmt"
	"sync"
	"time"
)

type Map struct {
	data       map[string]interface{}
	lock       sync.RWMutex
	addChan    chan addRequest
	getChan    chan getRequest
	deleteChan chan deleteRequest
}

type addRequest struct {
	key     string
	val     interface{}
	success chan<- bool
}

type getRequest struct {
	key      string
	response chan<- getResponse
}

type deleteRequest struct {
	key     string
	success chan<- bool
}

type getResponse struct {
	val interface{}
	ok  bool
}

func (m *Map) handleRequest() {

	for {
		select {
		case req := <-m.addChan:
			m.lock.Lock()
			m.data[req.key] = req.val
			m.lock.Unlock()
			req.success <- true

		case req := <-m.getChan:
			m.lock.RLock()
			val, ok := m.data[req.key]
			m.lock.RUnlock()
			req.response <- getResponse{val, ok}

		case req := <-m.deleteChan:
			m.lock.RLock()
			delete(m.data, req.key)
			m.lock.Unlock()
			req.success <- true
		default:
			fmt.Println("check")
			time.Sleep(time.Second)
		}

	}

}

func NewMap() *Map {

	mapObj := &Map{
		data:    make(map[string]interface{}),
		addChan: make(chan addRequest),
		getChan: make(chan getRequest),
	}

	//starting handleRequest
	go mapObj.handleRequest()

	return mapObj
}

func (m *Map) Set(key string, val interface{}) bool {
	success := make(chan bool)
	m.addChan <- addRequest{
		key,
		val,
		success,
	}
	return <-success
}

func (m *Map) Get(key string) getResponse {
	response := make(chan getResponse)
	m.getChan <- getRequest{
		key,
		response,
	}
	return <-response
}

func (m *Map) Delete(key string) bool {

	success := make(chan bool)
	m.deleteChan <- deleteRequest{
		key,
		success,
	}
	return <-success

}
