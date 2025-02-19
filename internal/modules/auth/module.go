// Code generated by candi v1.17.16.

package auth

import (
	// "go-api-service-with-candi/internal/modules/auth/delivery/graphqlhandler"
	// "go-api-service-with-candi/internal/modules/auth/delivery/grpchandler"
	"go-api-service-with-candi/internal/modules/auth/delivery/resthandler"
	// "go-api-service-with-candi/internal/modules/auth/delivery/workerhandler"
	"go-api-service-with-candi/pkg/shared/usecase"

	"github.com/golangid/candi/codebase/factory/dependency"
	"github.com/golangid/candi/codebase/factory/types"
	"github.com/golangid/candi/codebase/interfaces"
)

const (
	moduleName types.Module = "Auth"
)

// Module model
type Module struct {
	restHandler    interfaces.RESTHandler
	grpcHandler    interfaces.GRPCHandler
	graphqlHandler interfaces.GraphQLHandler

	workerHandlers map[types.Worker]interfaces.WorkerHandler
	serverHandlers map[types.Server]interfaces.ServerHandler
}

// NewModule module constructor
func NewModule(deps dependency.Dependency) *Module {
	var mod Module
	mod.restHandler = resthandler.NewRestHandler(usecase.GetSharedUsecase(), deps)
	// mod.grpcHandler = grpchandler.NewGRPCHandler(usecase.GetSharedUsecase(), deps)
	// mod.graphqlHandler = graphqlhandler.NewGraphQLHandler(usecase.GetSharedUsecase(), deps)

	mod.workerHandlers = map[types.Worker]interfaces.WorkerHandler{
		// types.Kafka:           workerhandler.NewKafkaHandler(usecase.GetSharedUsecase(), deps),
		// types.Scheduler:       workerhandler.NewCronHandler(usecase.GetSharedUsecase(), deps),
		// types.RedisSubscriber: workerhandler.NewRedisHandler(usecase.GetSharedUsecase(), deps),
		// types.TaskQueue:       workerhandler.NewTaskQueueHandler(usecase.GetSharedUsecase(), deps),
		// types.PostgresListener: workerhandler.NewPostgresListenerHandler(usecase.GetSharedUsecase(), deps),
		// types.RabbitMQ: workerhandler.NewRabbitMQHandler(usecase.GetSharedUsecase(), deps),
	}

	mod.serverHandlers = map[types.Server]interfaces.ServerHandler{
		// 
	}

	return &mod
}

// RESTHandler method
func (m *Module) RESTHandler() interfaces.RESTHandler {
	return m.restHandler
}

// GRPCHandler method
func (m *Module) GRPCHandler() interfaces.GRPCHandler {
	return m.grpcHandler
}

// GraphQLHandler method
func (m *Module) GraphQLHandler() interfaces.GraphQLHandler {
	return m.graphqlHandler
}

// WorkerHandler method
func (m *Module) WorkerHandler(workerType types.Worker) interfaces.WorkerHandler {
	return m.workerHandlers[workerType]
}

// ServerHandler additional server type (another rest framework, p2p, and many more)
func (m *Module) ServerHandler(serverType types.Server) interfaces.ServerHandler {
	return m.serverHandlers[serverType]
}

// Name get module name
func (m *Module) Name() types.Module {
	return moduleName
}
