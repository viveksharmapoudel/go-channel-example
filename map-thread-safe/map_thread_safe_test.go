package map_thread_safe

import (
	"log"
	"reflect"
	"testing"
)

func TestNewMap(t *testing.T) {
	log.Println("********************Initialize TestNewMap******************** ")
	m := NewMap()
	if reflect.TypeOf(m) != reflect.TypeOf(&Map{}) {
		t.Errorf("type mis match")
	}

	if m.addChan == nil {
		t.Errorf("chan add is nil")
	}

	if m.getChan == nil {
		t.Errorf("get chan is nil")
	}
	log.Println("********************END TestNewMap******************** ")

}

func TestMapSet(t *testing.T) {

	log.Println("********************Initialize TestMapSet******************** ")

	m := NewMap()

	if ok := m.Set("a", 20); !ok {
		t.Errorf("Value not set properly")
	}
	response := m.Get("a")
	if !response.ok && response.val != 20 {
		t.Errorf("value is not set properly ")
	}
	log.Println("********************End TestMapSet******************** ")

}
