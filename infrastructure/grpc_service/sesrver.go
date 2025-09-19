package grpc_service

import (
	"farm-service/bootstrap"

	grpc_server "github.com/anhvanhoa/service-core/bootstrap/grpc"
	"github.com/anhvanhoa/service-core/domain/log"
	proto_plant_variety "github.com/anhvanhoa/sf-proto/gen/plant_variety/v1"
	proto_planting_cycle "github.com/anhvanhoa/sf-proto/gen/planting_cycle/v1"
	"google.golang.org/grpc"
)

func NewGRPCServer(
	env *bootstrap.Env,
	log *log.LogGRPCImpl,
	plantVarietyService proto_plant_variety.PlantVarietyServiceServer,
	plantingCycleService proto_planting_cycle.PlantingCycleServiceServer,
) *grpc_server.GRPCServer {
	config := &grpc_server.GRPCServerConfig{
		IsProduction: env.IsProduction(),
		PortGRPC:     env.PortGrpc,
		NameService:  env.NameService,
	}
	return grpc_server.NewGRPCServer(
		config,
		log,
		func(server *grpc.Server) {
			proto_plant_variety.RegisterPlantVarietyServiceServer(server, plantVarietyService)
			proto_planting_cycle.RegisterPlantingCycleServiceServer(server, plantingCycleService)
		},
	)
}
