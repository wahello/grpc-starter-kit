package main

import (
	_ "github.com/xmlking/toolkit/logger/auto"

	"github.com/rs/zerolog/log"
	"github.com/xmlking/grpc-starter-kit/internal/config"
	"github.com/xmlking/grpc-starter-kit/internal/constants"
	"github.com/xmlking/grpc-starter-kit/service/emailer/registry"
	"github.com/xmlking/grpc-starter-kit/service/emailer/subscriber"
	"github.com/xmlking/toolkit/broker/cloudevents"
	"github.com/xmlking/toolkit/service"
)

func main() {
	serviceName := constants.EMAILER_SERVICE
	cfg := config.GetConfig()

	// Initialize DI Container
	ctn, err := registry.NewContainer(cfg)
	defer ctn.Clean()
	if err != nil {
		log.Fatal().Msgf("failed to build container: %v", err)
	}
	emailSubscriber := ctn.Resolve("emailer-subscriber").(*subscriber.EmailSubscriber)

	srv := service.NewService(
		service.Name(serviceName),
		service.Version(cfg.Services.Emailer.Version),
		// service.WithBrokerOptions(...),
	)
	bkr := broker.NewBroker(broker.Name("mkit.broker.emailer"))
	_, _ = bkr.NewSubscriber(cfg.Services.Emailer.Endpoint, emailSubscriber.HandleSend)
	err = bkr.Start()
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to start the Broker: %s", cfg.Services.Emailer.Endpoint)
	}

	// Start server!
	log.Info().Msg(config.GetBuildInfo())
	if err := srv.Start(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
