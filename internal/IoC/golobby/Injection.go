package golobby

import (
	"StorageWorkerService/internal/IoC"
	repository "StorageWorkerService/internal/repository/abstract"
	"StorageWorkerService/internal/repository/concrete/mongodb_driver"
	service "StorageWorkerService/internal/service/abstract"
	manager "StorageWorkerService/internal/service/concrete"
	jsonParser "StorageWorkerService/pkg/jsonParser"
	"StorageWorkerService/pkg/jsonParser/gojson"
	"StorageWorkerService/pkg/kafka"
	"StorageWorkerService/pkg/kafka/kafkago"
	"StorageWorkerService/pkg/logger"
	"StorageWorkerService/pkg/logger/logrus_logstash_hook"
	"StorageWorkerService/pkg/redis"
	"StorageWorkerService/pkg/redis/redisv8"
)
import "github.com/golobby/container/v3"

type golobbyInjection struct {}

func InjectionConstructor() *golobbyInjection {
	return &golobbyInjection{}
}

func (i *golobbyInjection) Inject(){
	i.bind()
	i.resolve()
}

func (i *golobbyInjection) bind(){
	_ = container.Transient(func() cache.ICache {
		return rediscachev8.RedisCacheConstructor()
	})

	_ = container.Transient(func() logger.ILog {
		return logrus_logstash_hook.LogrusToLogstashLOGConstructor()
	})

	_ = container.Transient(func() kafka.IKafka {
		return kafkago.KafkaGoConstructor()
	})

	_ = container.Transient(func() jsonParser.IJsonParser {
		return gojson.GoJsonConstructor()
	})

	_ = container.Singleton(func() service.IAdvEventService {
		return manager.AdvEventManagerConstructor(IoC.JsonParser, IoC.AdvEventDal)
	})

	_ = container.Singleton(func() service.IAdvBuyingService{
		return manager.BuyingEventManagerConstructor(IoC.JsonParser, IoC.BuyingEventDal, IoC.ClientService)
	})

	_ = container.Singleton(func() service.IChurnBlockerMlResultService{
		return manager.ChurnBlockerMlResultManagerConstructor(IoC.JsonParser, IoC.ChurnBlockerMlResultDal)
	})

	_ = container.Singleton(func() service.IChurnPredictionMlResultService{
		return manager.ChurnPredictionMlResultManagerConstructor(IoC.JsonParser, IoC.ChurnPredictionMlResultDal)
	})

	_ = container.Singleton(func() service.IClientService{
		return manager.ClientManagerConstructor(IoC.JsonParser, IoC.ClientDal)
	})

	_ = container.Singleton(func() service.IEnemyBaseLevelFailService{
		return manager.EnemyBaseLevelFailManagerConstructor(IoC.JsonParser, IoC.EnemyBaseLevelFailDal)
	})

	_ = container.Singleton(func() service.IEnemyBaseLoginLevelService{
		return manager.EnemyBaseLoginLevelManagerConstructor(IoC.JsonParser, IoC.EnemyBaseLoginLevelDal)
	})

	_ = container.Singleton(func() service.IGameSessionService{
		return manager.GameSessionManagerConstructor(IoC.JsonParser, IoC.GameSessionDal)
	})

	_ = container.Singleton(func() service.IHardwareService{
		return manager.HardwareManagerConstructor(IoC.JsonParser, IoC.HardwareDal)
	})

	_ = container.Singleton(func() service.IInventoryService{
		return manager.InventoryManagerConstructor(IoC.JsonParser, IoC.InventoryDal)
	})

	_ = container.Singleton(func() service.ILevelBaseSessionService{
		return manager.LevelBaseSessionManagerConstructor(IoC.JsonParser, IoC.LevelBaseSessionDal)
	})

	_ = container.Singleton(func() service.ILocationService{
		return manager.LocationManagerConstructor(IoC.JsonParser, IoC.LocationDal)
	})

	_ = container.Singleton(func() service.IManuelFlowService{
		return manager.ManuelFlowManagerConstructor(IoC.JsonParser, IoC.ManuelFlowDal)
	})

	_ = container.Singleton(func() service.IOfferBehaviorService{
		return manager.OfferBehaviorManagerConstructor(IoC.JsonParser, IoC.OfferBehaviorDal)
	})

	_ = container.Singleton(func() service.IScreenClickService{
		return manager.ScreenClickManagerConstructor(IoC.JsonParser, IoC.ScreenClickDal)
	})

	_ = container.Singleton(func() service.IScreenSwipeService{
		return manager.ScreenSwipeManagerConstructor(IoC.JsonParser, IoC.ScreenSwipeDal)
	})



	_ = container.Singleton(func() repository.IAdvEventDal{
		return mongodb_driver.MDbDAdvEventDalConstructor()
	})

	_ = container.Singleton(func() repository.IBuyingEventDal{
		return mongodb_driver.MDbDBuyingEventDalConstructor()
	})

	_ = container.Singleton(func() repository.IChurnBlockerMlResultDal{
		return mongodb_driver.MDbDChurnBlockerMlResultDalConstructor()
	})

	_ = container.Singleton(func() repository.IChurnPredictionMlResultDal{
		return mongodb_driver.MDbDChurnPredictionMlResultDalConstructor()
	})

	_ = container.Singleton(func() repository.IClientDal{
		return mongodb_driver.MDbDClientDalConstructor()
	})

	_ = container.Singleton(func() repository.IEnemyBaseLevelFailDal{
		return mongodb_driver.MDbDEnemyBaseLevelFailDalConstructor()
	})

	_ = container.Singleton(func() repository.IEnemyBaseLoginLevelDal{
		return mongodb_driver.MDbDEnemyBaseLoginLevelDalConstruct()
	})

	_ = container.Singleton(func() repository.IGameSessionDal{
		return mongodb_driver.MDbDGameSessionDalConstructor()
	})

	_ = container.Singleton(func() repository.IHardwareDal{
		return mongodb_driver.MDbDHardwareDalConstructor()
	})

	_ = container.Singleton(func() repository.IInventoryDal{
		return mongodb_driver.MDbDInventoryDalConstructor()
	})

	_ = container.Singleton(func() repository.ILevelBaseSessionDal{
		return mongodb_driver.MDbDLevelBaseSessionDalConstructor()
	})

	_ = container.Singleton(func() repository.ILocationDal{
		return mongodb_driver.MDbDLocationDalConstructor()
	})

	_ = container.Singleton(func() repository.IManuelFlowDal{
		return mongodb_driver.MDbDManuelFlowDalConstructor()
	})

	_ = container.Singleton(func() repository.IOfferBehaviorDal{
		return mongodb_driver.MDbDOfferBehaviorDalConstructor()
	})

	_ = container.Singleton(func() repository.IScreenClickDal{
		return mongodb_driver.MDbDScreenClickDalConstructor()
	})

	_ = container.Singleton(func() repository.IScreenSwipeDal{
		return mongodb_driver.MDbDScreenSwipeDalConstructor()
	})


}

func (i *golobbyInjection) resolve(){
	_ = container.Resolve(&IoC.RedisCache)
	_ = container.Resolve(&IoC.Logger)
	_ = container.Resolve(&IoC.Kafka)
	_ = container.Resolve(&IoC.JsonParser)

	_ = container.Resolve(&IoC.AdvEventService)
	_ = container.Resolve(&IoC.AdvBuyingService)
	_ = container.Resolve(&IoC.ChurnBlockerMlResultService)
	_ = container.Resolve(&IoC.ChurnPredictionMlResultService)
	_ = container.Resolve(&IoC.ClientService)
	_ = container.Resolve(&IoC.EnemyBaseLevelFailService)
	_ = container.Resolve(&IoC.EnemyBaseLoginLevelService)
	_ = container.Resolve(&IoC.GameSessionService)
	_ = container.Resolve(&IoC.HardwareService)
	_ = container.Resolve(&IoC.InventoryService)
	_ = container.Resolve(&IoC.LevelBaseSessionService)
	_ = container.Resolve(&IoC.LocationService)
	_ = container.Resolve(&IoC.ManuelFlowService)
	_ = container.Resolve(&IoC.OfferBehaviorService)
	_ = container.Resolve(&IoC.ScreenClickService)
	_ = container.Resolve(&IoC.ScreenSwipeService)

	_ = container.Resolve(&IoC.AdvEventDal)
	_ = container.Resolve(&IoC.BuyingEventDal)
	_ = container.Resolve(&IoC.ChurnBlockerMlResultDal)
	_ = container.Resolve(&IoC.ChurnPredictionMlResultDal)
	_ = container.Resolve(&IoC.ClientDal)
	_ = container.Resolve(&IoC.EnemyBaseLevelFailDal)
	_ = container.Resolve(&IoC.EnemyBaseLoginLevelDal)
	_ = container.Resolve(&IoC.GameSessionDal)
	_ = container.Resolve(&IoC.HardwareDal)
	_ = container.Resolve(&IoC.InventoryDal)
	_ = container.Resolve(&IoC.LevelBaseSessionDal)
	_ = container.Resolve(&IoC.LocationDal)
	_ = container.Resolve(&IoC.ManuelFlowDal)
	_ = container.Resolve(&IoC.OfferBehaviorDal)
	_ = container.Resolve(&IoC.ScreenClickDal)
	_ = container.Resolve(&IoC.ScreenSwipeDal)
}