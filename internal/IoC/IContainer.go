package IoC

import (
	repository "StorageWorkerService/internal/repository/abstract"
	service "StorageWorkerService/internal/service/abstract"
	jsonParser "StorageWorkerService/pkg/jsonParser"
	"StorageWorkerService/pkg/kafka"
	"StorageWorkerService/pkg/logger"
	cache "StorageWorkerService/pkg/redis"
)

type IContainer interface {
	Inject()
}

func InjectContainers(container IContainer){
	container.Inject()
}

var RedisCache cache.ICache
var Logger logger.ILog
var Kafka kafka.IKafka
var JsonParser jsonParser.IJsonParser

var AdvEventService service.IAdvEventService
var AdvBuyingService service.IAdvBuyingService
var ChurnBlockerMlResultService service.IChurnBlockerMlResultService
var ChurnPredictionMlResultService service.IChurnPredictionMlResultService
var ClientService service.IClientService
var EnemyBaseLevelFailService service.IEnemyBaseLevelFailService
var EnemyBaseLoginLevelService service.IEnemyBaseLoginLevelService
var GameSessionService service.IGameSessionService
var HardwareService service.IHardwareService
var InventoryService service.IInventoryService
var LevelBaseSessionService service.ILevelBaseSessionService
var LocationService service.ILocationService
var ManuelFlowService service.IManuelFlowService
var OfferBehaviorService service.IOfferBehaviorService
var ScreenClickService service.IScreenClickService
var ScreenSwipeService service.IScreenSwipeService

var AdvEventDal repository.IAdvEventDal
var BuyingEventDal repository.IBuyingEventDal
var ChurnBlockerMlResultDal repository.IChurnBlockerMlResultDal
var ChurnPredictionMlResultDal repository.IChurnPredictionMlResultDal
var ClientDal repository.IClientDal
var EnemyBaseLevelFailDal repository.IEnemyBaseLevelFailDal
var EnemyBaseLoginLevelDal repository.IEnemyBaseLoginLevelDal
var GameSessionDal repository.IGameSessionDal
var HardwareDal repository.IHardwareDal
var InventoryDal repository.IInventoryDal
var LevelBaseSessionDal repository.ILevelBaseSessionDal
var LocationDal repository.ILocationDal
var ManuelFlowDal repository.IManuelFlowDal
var OfferBehaviorDal repository.IOfferBehaviorDal
var ScreenClickDal repository.IScreenClickDal
var ScreenSwipeDal repository.IScreenSwipeDal

