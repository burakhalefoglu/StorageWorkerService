package golobby

import (
	"StorageWorkerService/internal/IoC"
	repository "StorageWorkerService/internal/repository/abstract"
	"StorageWorkerService/internal/repository/concrete/Cassandra"
	service "StorageWorkerService/internal/service/abstract"
	manager "StorageWorkerService/internal/service/concrete"
	cassandra "StorageWorkerService/pkg/database/Cassandra"
	jsonParser "StorageWorkerService/pkg/jsonParser"
	"StorageWorkerService/pkg/jsonParser/gojson"
	"StorageWorkerService/pkg/kafka"
	"StorageWorkerService/pkg/kafka/kafkago"
	cache "StorageWorkerService/pkg/redis"
	rediscachev8 "StorageWorkerService/pkg/redis/redisv8"

	"github.com/golobby/container/v3"
)

type golobbyInjection struct{}

func InjectionConstructor() *golobbyInjection {
	return &golobbyInjection{}
}

func (i *golobbyInjection) Inject() {
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
	injectChurnPredictionSuccessRateResult()
	injectAddAdvStrategyBehavior()
}

func injectScreenSwipe() {
	if err := container.Singleton(func() service.IScreenSwipeService {
		return manager.ScreenSwipeManagerConstructor()
	}); err != nil {
		panic(err)
	}

	if err := container.Singleton(func() repository.IScreenSwipeDal {
		return Cassandra.NewScreenSwipeDal(cassandra.ScreenSwipeModel)
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
		return Cassandra.NewScreenClickDal(cassandra.ScreenClickModel)
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
		return Cassandra.NewOfferBehaviorDal(cassandra.OfferBehaviorModel)
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
		return Cassandra.NewManuelFlowDal(cassandra.ManuelFlowModel)
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
		return Cassandra.NewLocationDal(cassandra.LocationModel)
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
		return Cassandra.NewLevelBaseSessionDal(cassandra.LevelBaseSessionModel)
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

	if err := container.Singleton(func() service.ITemporaryAbilityService {
		return manager.NewTemporaryAbilityManager()
	}); err != nil {
		panic(err)
	}
	if err := container.Singleton(func() repository.ITemporaryAbilityDal {
		return Cassandra.NewTemporaryAbilityDal(cassandra.TemporaryAbilityModel)
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.TemporaryAbilityService); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.TemporaryAbilityDal); err != nil {
		panic(err)
	}

	if err := container.Singleton(func() service.IItemService {
		return manager.NewItemManager()
	}); err != nil {
		panic(err)
	}
	if err := container.Singleton(func() repository.IItemDal {
		return Cassandra.NewItemDal(cassandra.ItemModel)
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.ItemService); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.ItemDal); err != nil {
		panic(err)
	}

	if err := container.Singleton(func() service.ISkillService {
		return manager.NewSkillManager()
	}); err != nil {
		panic(err)
	}
	if err := container.Singleton(func() repository.ISkillDal {
		return Cassandra.NewSkillDal(cassandra.SkillModel)
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.SkillService); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.SkillDal); err != nil {
		panic(err)
	}

	if err := container.Singleton(func() service.IInventoryService {
		return manager.NewInventoryManager()
	}); err != nil {
		panic(err)
	}
	if err := container.Singleton(func() repository.IInventoryDal {
		return Cassandra.NewInventoryDal(cassandra.InventoryModel)
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
		return Cassandra.NewHardwareDal(cassandra.HardwareModel)
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
		return Cassandra.NewGameSessionDal(cassandra.GameSessionModel)
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
		return Cassandra.NewEnemyBaseLoginLevelDal(cassandra.EnemyBaseLoginLevelModel)
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
		return Cassandra.NewEnemyBaseLevelFailDal(cassandra.EnemyBaseLevelFailModel)
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
		return Cassandra.NewCassChurnPredictionMlResultDal(cassandra.ChurnPredictionMlResultModel)
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
		return Cassandra.NewCassChurnBlockerMlResultDal(cassandra.ChurnBlockerMlResultModel)
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
	if err := container.Singleton(func() service.IBuyingEventService {
		return manager.BuyingEventManagerConstructor()
	}); err != nil {
		panic(err)
	}
	if err := container.Singleton(func() repository.IBuyingEventDal {
		return Cassandra.NewCassBuyingEventDal(cassandra.BuyingEventModel)
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
		return Cassandra.NewCassClientDal(cassandra.ClientDataModel)
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
		return Cassandra.NewCassAdvEventDal(cassandra.AdvEventDataModel)
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

func injectChurnPredictionSuccessRateResult() {
	if err := container.Singleton(func() service.IChurnPredictionSuccessRateService {
		return manager.NewChurnPredictionSuccessRate()
	}); err != nil {
		panic(err)
	}
	if err := container.Singleton(func() repository.IChurnPredictionSuccessRateDal {
		return Cassandra.NewCassChurnPredictionSuccessRateDal(cassandra.ChurnPredictionSuccessRateModel)
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.ChurnPredictionSuccessRateDal); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.ChurnPredictionSuccessRateService); err != nil {
		panic(err)
	}

}

func injectAddAdvStrategyBehavior() {
	if err := container.Singleton(func() service.IAdvStrategyBehaviorService {
		return manager.NewAdvStrategyBehaviorManager()
	}); err != nil {
		panic(err)
	}
	if err := container.Singleton(func() repository.IAdvStrategyBehaviorDal {
		return Cassandra.NewCassAdvStrategyBehaviorDal(cassandra.AdvStrategyBehaviorModel)
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.AdvStrategyBehaviorDal); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.AdvStrategyBehaviorService); err != nil {
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
		return kafkago.KafkaGoConstructor()
	}); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.Kafka); err != nil {
		panic(err)
	}
}

func injectCache() {
	if err := container.Singleton(func() cache.ICache {
		return rediscachev8.RedisCacheConstructor()
	}); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.RedisCache); err != nil {
		panic(err)
	}
}
