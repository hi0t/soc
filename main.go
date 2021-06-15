package main

import (
	"log"
	"os"

	"github.com/kardianos/service"
)

func isAction(f string) bool {
	for _, c := range service.ControlAction {
		if c == f {
			return true
		}
	}
	return false
}

func main() {
	svcConfig := &service.Config{
		Name: "soc",
	}

	a := &app{}
	s, err := service.New(a, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) > 1 && isAction(os.Args[1]) {
		svcConfig.Arguments = os.Args[2:]
		err = service.Control(s, os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	a.l, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Run(); err != nil {
		a.l.Error(err)
	}
}
