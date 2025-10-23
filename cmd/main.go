package main

import (
	"context"
	"farm-service/bootstrap"
	"farm-service/infrastructure/grpc_service"
	plant_variety_service "farm-service/infrastructure/grpc_service/plant_variety"
	planting_cycle_service "farm-service/infrastructure/grpc_service/planting_cycle"

	"github.com/anhvanhoa/service-core/domain/discovery"
)

func main() {
	StartGRPCServer()
}

func StartGRPCServer() {
	app := bootstrap.App()
	env := app.Env
	log := app.Log

	discoveryConfig := &discovery.DiscoveryConfig{
		ServiceName:   env.NameService,
		ServicePort:   env.PortGrpc,
		ServiceHost:   env.HostGprc,
		IntervalCheck: env.IntervalCheck,
		TimeoutCheck:  env.TimeoutCheck,
	}

	discovery, err := discovery.NewDiscovery(discoveryConfig)
	if err != nil {
		log.Fatal("Failed to create discovery: " + err.Error())
	}
	discovery.Register()

	plantVarietyService := plant_variety_service.NewPlantVarietyService(app.Repos.PlantVarietyRepository)
	plantingCycleService := planting_cycle_service.NewPlantingCycleService(app.Repos.PlantingCycleRepository)

	grpcSrv := grpc_service.NewGRPCServer(
		env, log,
		app.Cache,
		plantVarietyService,
		plantingCycleService,
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := grpcSrv.Start(ctx); err != nil {
		log.Fatal("gRPC server error: " + err.Error())
	}
}
