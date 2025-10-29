package grpc_service

import (
	"farm-service/bootstrap"

	grpc_service "github.com/anhvanhoa/service-core/bootstrap/grpc"
	"github.com/anhvanhoa/service-core/domain/cache"
	"github.com/anhvanhoa/service-core/domain/log"
	"github.com/anhvanhoa/service-core/domain/token"
	"github.com/anhvanhoa/service-core/domain/user_context"
	proto_plant_variety "github.com/anhvanhoa/sf-proto/gen/plant_variety/v1"
	proto_planting_cycle "github.com/anhvanhoa/sf-proto/gen/planting_cycle/v1"
	"google.golang.org/grpc"
)

func NewGRPCServer(
	env *bootstrap.Env,
	log *log.LogGRPCImpl,
	cacher cache.CacheI,
	plantVarietyService proto_plant_variety.PlantVarietyServiceServer,
	plantingCycleService proto_planting_cycle.PlantingCycleServiceServer,
) *grpc_service.GRPCServer {
	config := &grpc_service.GRPCServerConfig{
		IsProduction: env.IsProduction(),
		PortGRPC:     env.PortGrpc,
		NameService:  env.NameService,
	}
	middleware := grpc_service.NewMiddleware(
		token.NewToken(env.AccessSecret),
	)
	return grpc_service.NewGRPCServer(
		config,
		log,
		func(server *grpc.Server) {
			proto_plant_variety.RegisterPlantVarietyServiceServer(server, plantVarietyService)
			proto_planting_cycle.RegisterPlantingCycleServiceServer(server, plantingCycleService)
		},
		middleware.AuthorizationInterceptor(
			env.SecretService,
			func(action string, resource string) bool {
				hasPermission, err := cacher.Get(resource + "." + action)
				if err != nil {
					return false
				}
				return hasPermission != nil && string(hasPermission) == "true"
			},
			func(id string) *user_context.UserContext {
				userData, err := cacher.Get(id)
				if err != nil || userData == nil {
					return nil
				}
				uCtx := user_context.NewUserContext()
				uCtx.FromBytes(userData)
				return uCtx
			},
		),
	)
}
