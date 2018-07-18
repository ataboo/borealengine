package services

import (
	. "github.com/ataboo/borealengine/config"
	"fmt"
	"time"
	"github.com/op/go-logging"
	"context"
	"github.com/ataboo/borealengine/services/network"
)

type Service interface {
	Start(ctx context.Context)
	Update(delta time.Duration)
}

func NewEngine() *Engine {
	engine := Engine{
		false,
		0,
		time.Time{},
		[]Service{&network.NetworkService{}},
	}

	return &engine
}


type Engine struct {
	running  bool
	delta    time.Duration
	lastTick time.Time
	services []Service
}

var Logger *logging.Logger

func (e *Engine) Start() (context.CancelFunc, error) {
	if e.running {
		return nil, fmt.Errorf("engine already running")
	}

	Logger = logging.MustGetLogger(Config.General.LogName)

	return e.startLoop()
}

func (e *Engine) startLoop() (context.CancelFunc, error) {
	ctx, stop := context.WithCancel(context.Background())

	targetDelta := time.Duration(Config.General.DeltaMicro) * time.Microsecond
	deltaTick := time.Tick(targetDelta)
	e.lastTick = time.Now().Add(-targetDelta)
	e.startServices(ctx)

	go func() {
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case <-deltaTick:
					e.delta = e.lastTick.Sub(time.Now())
					e.lastTick = time.Now()

					e.updateServices()
				default:
					//
				}
			}
		}()
	}()

	e.running = true

	return stop, nil
}

func (e *Engine) startServices(ctx context.Context) error {
	for _, s := range e.services {
		s.Start(ctx)
	}

	return nil
}

func (e *Engine) updateServices() error {
	return nil
}
