package entity

import "context"

type GameLoop struct {

}

func (g *GameLoop) Start(ctx context.Context) {
	go func() {
		for {
			select {
			case <- ctx.Done():
				return
			default:
				//
			}
		}
	}()
}
