package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Service struct {
	sync.Mutex
	c1    chan string
	state string
}

func (s *Service) Start() error {
	if s.state == "stopped" {
		s.state = "start"
		for {
			time.Sleep(time.Second * 1)
			fmt.Println("Yo!")

			select {
			case <-s.c1:
				fmt.Print("Program stopped")
				return nil
			default:
				continue
			}
		}
	}

	return errors.New("process already started")
}

func (s *Service) Stop() error {
	s.Lock()
	if s.state == "stopped" {
		s.Unlock()
		return errors.New("process already stoped")
	}

	s.state = "stopped"
	s.Unlock()

	s.c1 <- "done"
	return nil
}

func (s *Service) Close() {
	close(s.c1)
}

func NewService() *Service {
	return &Service{
		c1:    make(chan string),
		state: "stopped",
	}
}

func main() {
	service := NewService()

	go func() {
		time.Sleep(time.Second * 10)

		_ = service.Stop()
	}()

	_ = service.Start()
}
