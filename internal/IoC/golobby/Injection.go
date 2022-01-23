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

type golobbyInjection struct{}

func InjectionConstructor() *golobbyInjection {
	return &golobbyInjection{}
}

func (i *golobbyInjection) Inject() {
	injectLogger() // Logger is everywhere, so is must be top
	injectKafka()
	injectJsonParser()
	injectCache()

	injectAdvEvent()
	injectClient()
	injectBuyingEvent()
	injectChurnBlockerMlResult()
	injectChurnPredictionResult()
	injectEnemyBaseLevelFail()
	injectEnemyBaseLoginLevel()
	injectGameSession()
	injectHardware()
	injectInventory()
	injectLevelBaseSession()
	injectLocation()
	injectManuelFlow()
	injectOfferBehavior()
	injectScreenClick()
	injectScreenSwipe()
}

func injectScreenSwipe() {
	if err := container.Singleton(func() service.IScreenSwipeService {
		return manager.ScreenSwipeManagerConstructor()
	}); err != nil {
		panic(err)
	}

	if err := container.Singleton(func() repository.IScreenSwipeDal {
		return mongodb_driver.MDbDScreenSwipeDalConstructor()
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.ScreenSwipeService); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.ScreenSwipeDal); err != nil {
		panic(err)
	}
}

func injectScreenClick() {
	if err := container.Singleton(func() service.IScreenClickService {
		return manager.ScreenClickManagerConstructor()
	}); err != nil {
		panic(err)
	}
	if err := container.Singleton(func() repository.IScreenClickDal {
		return mongodb_driver.MDbDScreenClickDalConstructor()
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.ScreenClickService); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.ScreenClickDal); err != nil {
		panic(err)
	}
}

func injectOfferBehavior() {
	if err := container.Singleton(func() service.IOfferBehaviorService {
		return manager.OfferBehaviorManagerConstructor()
	}); err != nil {
		panic(err)
	}
	if err := container.Singleton(func() repository.IOfferBehaviorDal {
		return mongodb_driver.MDbDOfferBehaviorDalConstructor()
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.OfferBehaviorService); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.OfferBehaviorDal); err != nil {
		panic(err)
	}
}

func injectManuelFlow() {
	if err := container.Singleton(func() service.IManuelFlowService {
		return manager.ManuelFlowManagerConstructor()
	}); err != nil {
		panic(err)
	}
	if err := container.Singleton(func() repository.IManuelFlowDal {
		return mongodb_driver.MDbDManuelFlowDalConstructor()
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.ManuelFlowService); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.ManuelFlowDal); err != nil {
		panic(err)
	}
}

func injectLocation() {
	if err := container.Singleton(func() service.ILocationService {
		return manager.LocationManagerConstructor()
	}); err != nil {
		panic(err)
	}
	if err := container.Singleton(func() repository.ILocationDal {
		return mongodb_driver.MDbDLocationDalConstructor()
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.LocationService); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.LocationDal); err != nil {
		panic(err)
	}
}

func injectLevelBaseSession() {
	if err := container.Singleton(func() service.ILevelBaseSessionService {
		return manager.LevelBaseSessionManagerConstructor()
	}); err != nil {
		panic(err)
	}
	if err := container.Singleton(func() repository.ILevelBaseSessionDal {
		return mongodb_driver.MDbDLevelBaseSessionDalConstructor()
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.LevelBaseSessionService); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.LevelBaseSessionDal); err != nil {
		panic(err)
	}
}

func injectInventory() {
	if err := container.Singleton(func() service.IInventoryService {
		return manager.InventoryManagerConstructor()
	}); err != nil {
		panic(err)
	}
	if err := container.Singleton(func() repository.IInventoryDal {
		return mongodb_driver.MDbDInventoryDalConstructor()
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.InventoryService); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.InventoryDal); err != nil {
		panic(err)
	}
}

func injectHardware() {
	if err := container.Singleton(func() service.IHardwareService {
		return manager.HardwareManagerConstructor()
	}); err != nil {
		panic(err)
	}
	if err := container.Singleton(func() repository.IHardwareDal {
		return mongodb_driver.MDbDHardwareDalConstructor()
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.HardwareService); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.HardwareDal); err != nil {
		panic(err)
	}
}

func injectGameSession() {
	if err := container.Singleton(func() service.IGameSessionService {
		return manager.GameSessionManagerConstructor()
	}); err != nil {
		panic(err)
	}
	if err := container.Singleton(func() repository.IGameSessionDal {
		return mongodb_driver.MDbDGameSessionDalConstructor()
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.GameSessionService); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.GameSessionDal); err != nil {
		panic(err)
	}
}

func injectEnemyBaseLoginLevel() {
	if err := container.Singleton(func() service.IEnemyBaseLoginLevelService {
		return manager.EnemyBaseLoginLevelManagerConstructor()
	}); err != nil {
		panic(err)
	}
	if err := container.Singleton(func() repository.IEnemyBaseLoginLevelDal {
		return mongodb_driver.MDbDEnemyBaseLoginLevelDalConstruct()
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.EnemyBaseLoginLevelService); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.EnemyBaseLoginLevelDal); err != nil {
		panic(err)
	}
}

func injectEnemyBaseLevelFail() {
	if err := container.Singleton(func() service.IEnemyBaseLevelFailService {
		return manager.EnemyBaseLevelFailManagerConstructor()
	}); err != nil {
		panic(err)
	}
	if err := container.Singleton(func() repository.IEnemyBaseLevelFailDal {
		return mongodb_driver.MDbDEnemyBaseLevelFailDalConstructor()
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.EnemyBaseLevelFailService); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.EnemyBaseLevelFailDal); err != nil {
		panic(err)
	}
}

func injectChurnPredictionResult() {
	if err := container.Singleton(func() service.IChurnPredictionMlResultService {
		return manager.ChurnPredictionMlResultManagerConstructor()
	}); err != nil {
		panic(err)
	}
	if err := container.Singleton(func() repository.IChurnPredictionMlResultDal {
		return mongodb_driver.MDbDChurnPredictionMlResultDalConstructor()
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.ChurnPredictionMlResultService); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.ChurnPredictionMlResultDal); err != nil {
		panic(err)
	}
}

func injectChurnBlockerMlResult() {
	if err := container.Singleton(func() service.IChurnBlockerMlResultService {
		return manager.ChurnBlockerMlResultManagerConstructor()
	}); err != nil {
		panic(err)
	}
	if err := container.Singleton(func() repository.IChurnBlockerMlResultDal {
		return mongodb_driver.MDbDChurnBlockerMlResultDalConstructor()
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.ChurnBlockerMlResultService); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.ChurnBlockerMlResultDal); err != nil {
		panic(err)
	}
}

func injectBuyingEvent() {
	if err := container.Singleton(func() service.IAdvBuyingService {
		return manager.BuyingEventManagerConstructor()
	}); err != nil {
		panic(err)
	}
	if err := container.Singleton(func() repository.IBuyingEventDal {
		return mongodb_driver.MDbDBuyingEventDalConstructor()
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.AdvBuyingService); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.BuyingEventDal); err != nil {
		panic(err)
	}
}

func injectClient() {
	if err := container.Singleton(func() service.IClientService {
		return manager.ClientManagerConstructor()
	}); err != nil {
		panic(err)
	}
	if err := container.Singleton(func() repository.IClientDal {
		return mongodb_driver.MDbDClientDalConstructor()
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.ClientDal); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.ClientService); err != nil {
		panic(err)
	}
}

func injectAdvEvent() {
	if err := container.Singleton(func() service.IAdvEventService {
		return manager.AdvEventManagerConstructor()
	}); err != nil {
		panic(err)
	}
	if err := container.Singleton(func() repository.IAdvEventDal {
		return mongodb_driver.MDbDAdvEventDalConstructor()
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.AdvEventDal); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.AdvEventService); err != nil {
		panic(err)
	}

}

func injectJsonParser() {
	if err := container.Singleton(func() jsonParser.IJsonParser {
		return gojson.GoJsonConstructor()
	}); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.JsonParser); err != nil {
		panic(err)
	}
}

func injectKafka() {
	if err := container.Singleton(func() kafka.IKafka {
		return kafkago.KafkaGoConstructor(&IoC.Logger)
	}); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.Kafka); err != nil {
		panic(err)
	}
}

func injectLogger() {
	if err := container.Singleton(func() logger.ILog {
		return logrus_logstash_hook.LogrusToLogstashLOGConstructor()
	}); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.Logger); err != nil {
		panic(err)
	}
}

func injectCache() {
	if err := container.Singleton(func() cache.ICache {
		return rediscachev8.RedisCacheConstructor(&IoC.Logger)
	}); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.RedisCache); err != nil {
		panic(err)
	}
}
