package KafkaController

import (
	"StorageWorkerService/internal/IoC"
	"StorageWorkerService/internal/service/abstract"
	"StorageWorkerService/pkg/helper"
	"StorageWorkerService/pkg/kafka"
	"sync"
)

type insertController struct {
	Kafka                          *kafka.IKafka
	AdvEventService                *abstract.IAdvEventService
	AdvBuyingService               *abstract.IAdvBuyingService
	HardwareService                *abstract.IHardwareService
	LocationService                *abstract.ILocationService
	ScreenSwipeService             *abstract.IScreenSwipeService
	ScreenClickService             *abstract.IScreenClickService
	LevelBaseSessionService        *abstract.ILevelBaseSessionService
	GameSessionService             *abstract.IGameSessionService
	InventoryService               *abstract.IInventoryService
	EnemyBaseLoginLevelService     *abstract.IEnemyBaseLoginLevelService
	EnemyBaseLevelFailService      *abstract.IEnemyBaseLevelFailService
	OfferBehaviorService           *abstract.IOfferBehaviorService
	ChurnPredictionMlResultService *abstract.IChurnPredictionMlResultService
	ChurnBlockerMlResultService    *abstract.IChurnBlockerMlResultService
	ManuelFlowService              *abstract.IManuelFlowService
}

func InsertControllerConstructor() *insertController {
	return &insertController{Kafka: &IoC.Kafka,
		AdvEventService:                &IoC.AdvEventService,
		AdvBuyingService:               &IoC.AdvBuyingService,
		HardwareService:                &IoC.HardwareService,
		LocationService:                &IoC.LocationService,
		ScreenSwipeService:             &IoC.ScreenSwipeService,
		ScreenClickService:             &IoC.ScreenClickService,
		LevelBaseSessionService:        &IoC.LevelBaseSessionService,
		GameSessionService:             &IoC.GameSessionService,
		InventoryService:               &IoC.InventoryService,
		EnemyBaseLoginLevelService:     &IoC.EnemyBaseLoginLevelService,
		EnemyBaseLevelFailService:      &IoC.EnemyBaseLevelFailService,
		OfferBehaviorService:           &IoC.OfferBehaviorService,
		ChurnPredictionMlResultService: &IoC.ChurnPredictionMlResultService,
		ChurnBlockerMlResultService:    &IoC.ChurnBlockerMlResultService,
		ManuelFlowService:              &IoC.ManuelFlowService}
}

func (controller *insertController) StartListen(waitGroup *sync.WaitGroup) {
	waitGroup.Add(15)
	helper.CreateHealthFile()
	go (*controller.Kafka).Consume("AdvEventDataModel",
		"AdvEventDataModel_ConsumerGroup",
		waitGroup,
		(*controller.AdvEventService).AddAdvEventData)

	go (*controller.Kafka).Consume("BuyingEventDataModel",
		"BuyingEventDataModel_ConsumerGroup",
		waitGroup,
		(*controller.AdvBuyingService).AddBuyingEventData)

	go (*controller.Kafka).Consume("HardwareInformationModel",
		"HardwareInformationModel_ConsumerGroup",
		waitGroup,
		(*controller.HardwareService).AddHardwareData)

	go (*controller.Kafka).Consume("LocationDataModel",
		"LocationDataModel_ConsumerGroup",
		waitGroup,
		(*controller.LocationService).AddLocationData)

	go (*controller.Kafka).Consume("ScreenSwipeDataModel",
		"ScreenSwipeDataModel_ConsumerGroup",
		waitGroup,
		(*controller.ScreenSwipeService).AddScreenSwipeData)

	go (*controller.Kafka).Consume("ScreenClickDataModel",
		"ScreenClickDataModel_ConsumerGroup",
		waitGroup,
		(*controller.ScreenClickService).AddScreenClickData)

	go (*controller.Kafka).Consume("LevelBaseSessionDataModel",
		"LevelBaseSessionDataModel_ConsumerGroup",
		waitGroup,
		(*controller.LevelBaseSessionService).AddLevelBaseSessionData)

	go (*controller.Kafka).Consume("GameSessionEveryLoginDataModel",
		"GameSessionEveryLoginDataModel_ConsumerGroup",
		waitGroup,
		(*controller.GameSessionService).AddGameSessionData)

	go (*controller.Kafka).Consume("InventoryDataModel",
		"InventoryModel_ConsumerGroup",
		waitGroup,
		(*controller.InventoryService).AddInventoryData)

	go (*controller.Kafka).Consume("EnemyBaseEveryLoginLevelDataModel",
		"EnemyBaseEveryLoginLevelDataModel_ConsumerGroup",
		waitGroup,
		(*controller.EnemyBaseLoginLevelService).AddEnemyBaseLoginLevelData)

	go (*controller.Kafka).Consume("EnemyBaseWithLevelFailDataModel",
		"EnemyBaseWithLevelFailDataModel_ConsumerGroup",
		waitGroup,
		(*controller.EnemyBaseLevelFailService).AddEnemyBaseLevelFailData)

	go (*controller.Kafka).Consume("ManuelFlowModel",
		"ManuelFlowModel_ConsumerGroup",
		waitGroup,
		(*controller.ManuelFlowService).AddManuelFlowData)

	go (*controller.Kafka).Consume("OfferBehaviorModel",
		"OfferBehaviorModel_ConsumerGroup",
		waitGroup,
		(*controller.OfferBehaviorService).AddOfferBehaviorData)

	go (*controller.Kafka).Consume("ChurnPredictionResultModel",
		"ChurnPredictionMlResultModel_ConsumerGroup",
		waitGroup,
		(*controller.ChurnPredictionMlResultService).AddChurnPredictionMlResultData)

	go (*controller.Kafka).Consume("ChurnBlockerResultModel",
		"ChurnBlockerMlResultModel_ConsumerGroup",
		waitGroup,
		(*controller.ChurnBlockerMlResultService).AddChurnBlockerMlResultData)
}
