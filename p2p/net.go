package p2p

import (
	"context"
	"time"

	application "github.com/pravahio/go-mesh/application"
	bootservice "github.com/pravahio/go-mesh/service/bootstrap"
)

var BootstrapList = []string{
	"/ip4/13.234.78.241/udp/4000/quic/p2p/QmVbcMycaK8ni5CeiM7JRjBRAdmwky6dQ6KcoxLesZDPk9",
}

type P2P struct {
	app       *application.Application
	BCService *BCService
}

func NewP2P(ctx context.Context) (*P2P, error) {
	/* h, err := libp2p.New(ctx)
	if err != nil {
		return nil, err
	}
	return &P2P{
		host: h,
	}, nil */

	app, err := application.NewApplication(
		ctx,
		nil,
		nil,
		"0.0.0.0",
		"0",
		false,
	)
	if err != nil {
		return nil, err
	}

	bservice := bootservice.NewBootstrapService(false, "rpoint", BootstrapList, 5*time.Second)
	bcService := NewBCService()

	app.InjectService(bservice)
	app.InjectService(bcService)

	app.Start()

	return &P2P{
		app:       app,
		BCService: bcService,
	}, nil
}

func (p *P2P) Wait() {
	p.app.Wait()
}
