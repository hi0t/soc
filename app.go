package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/kardianos/service"
)

type app struct {
	p *http.Server
	l service.Logger
}

func (a *app) Start(s service.Service) error {
	var host string
	var port int

	flag.StringVar(&host, "h", "", "host")
	flag.IntVar(&port, "p", 8888, "port")
	flag.Parse()

	a.p = newProxy(fmt.Sprintf("%s:%d", host, port))

	go func() {
		if err := a.p.ListenAndServe(); err != nil {
			a.l.Error(err)
		}
	}()

	return nil
}

func (a *app) Stop(s service.Service) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return a.p.Shutdown(ctx)
}
