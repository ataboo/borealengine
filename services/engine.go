package services

import (
	"fmt"
	"time"
	"context"
	"github.com/ataboo/borealengine/services/network"
	"github.com/ataboo/borealengine/config"
	"github.com/ataboo/borealengine/services/entity"
)

func NewEngine() *Engine {
	engine := Engine{
		false,
		0,
		time.Time{},
		network.NewService(),
		entity.Service{},
	}

	return &engine
}


type Engine struct {
	running  bool
	delta    time.Duration
	lastTick time.Time
	network *network.Service
	entity entity.Service
}

func (e *Engine) Start() (context.CancelFunc, error) {
	if e.running {
		return nil, fmt.Errorf("engine already running")
	}

	return e.startLoop()
}

func (e *Engine) startLoop() (context.CancelFunc, error) {
	ctx, stop := context.WithCancel(context.Background())

	targetDelta := time.Duration(config.Config.General.DeltaMicro) * time.Microsecond
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
				case userConn := <-e.network.OnConnect:
					e.entity.PlayerConnected(userConn.User, userConn.Conn)
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
	e.network.Start(ctx)
	e.entity.Start(ctx)

	return nil
}

func (e *Engine) updateServices() error {
	e.entity.Update(e.delta)

	return nil
}
