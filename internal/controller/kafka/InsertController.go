package KafkaInsertController

import (
	"StorageWorkerService/internal/service/abstract"
	Ikafka "StorageWorkerService/pkg/kafka"
	"sync"
)

type InsertController struct {
	Kafka Ikafka.IKafka
	AdvM abstract.IAdvEventService
	BuyingM  abstract.IAdvBuyingService
	HardwareM abstract.IHardwareService
	LocM abstract.ILocationService
	ScrSwipeM abstract.IScreenSwipeService
	ScrClickM abstract.IScreenClickService
	LvlBaseSessionM abstract.ILevelBaseSessionService
	GameSessionM abstract.IGameSessionService
	EnemyBaseLoginLM abstract.IEnemyBaseLoginLevelService
	EnemyBaseLevelFM abstract.IEnemyBaseLevelFailService
	ManuelFlowM abstract.IManuelFlowService
	OfferBehaviorM abstract.IOfferBehaviorService
	CBMlResultM abstract.IChurnBlockerMlResultService
	CPMResultM abstract.IChurnPredictionMlResultService
	InventoryM abstract.IInventoryService
}


func (controller *InsertController) StartListen() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(11)

	go controller.Kafka.Consume("AdvEventDataModel",
		"AdvEventDataModel_ConsumerGroup",
		&waitGroup,
		controller.AdvM.AddAdvEventData)

	go 	controller.Kafka.Consume("BuyingEventDataModel",
		"BuyingEventDataModel_ConsumerGroup",
		&waitGroup,
		controller.BuyingM.AddBuyingEventData)

	go 	controller.Kafka.Consume("HardwareInformationModel",
		"HardwareInformationModel_ConsumerGroup",
		&waitGroup,
		controller.HardwareM.AddHardwareData)

	go 	controller.Kafka.Consume("LocationDataModel",
		"LocationDataModel_ConsumerGroup",
		&waitGroup,
		controller.LocM.AddLocationData)

	go 	controller.Kafka.Consume("ScreenSwipeDataModel",
		"ScreenSwipeDataModel_ConsumerGroup",
		&waitGroup,
		controller.ScrSwipeM.AddScreenSwipeData)

	go 	controller.Kafka.Consume("ScreenClickDataModel",
		"ScreenClickDataModel_ConsumerGroup",
		&waitGroup,
		controller.ScrClickM.AddScreenClickData)

	go 	controller.Kafka.Consume("LevelBaseSessionDataModel",
		"LevelBaseSessionDataModel_ConsumerGroup",
		&waitGroup,
		controller.LvlBaseSessionM.AddLevelBaseSessionData)

	go 	controller.Kafka.Consume("GameSessionEveryLoginDataModel",
		"GameSessionEveryLoginDataModel_ConsumerGroup",
		&waitGroup,
		controller.GameSessionM.AddGameSessionData)

	go 	controller.Kafka.Consume("InventoryDataModel",
		"InventoryModel_ConsumerGroup",
		&waitGroup,
		controller.InventoryM.AddInventoryData)


	//go 	controller.Kafka.Consume("EnemyBaseEveryLoginLevelDataModel",
	//	"EnemyBaseEveryLoginLevelDataModel_ConsumerGroup",
	//	&waitGroup,
	//	controller.EnemyBaseLoginLM.AddEnemyBaseLoginLevelData)
	//
	//go 	controller.Kafka.Consume("EnemyBaseWithLevelFailDataModel",
	//	"EnemyBaseWithLevelFailDataModel_ConsumerGroup",
	//	&waitGroup,
	//	controller.EnemyBaseLevelFM.AddEnemyBaseLevelFailData)

	//go 	controller.Kafka.Consume("ManuelFlowModel",
	//	"ManuelFlowModel_ConsumerGroup",
	//	&waitGroup,
	//	controller.ManuelFlowM.AddManuelFlowData)

	go 	controller.Kafka.Consume("OfferBehaviorModel",
		"OfferBehaviorModel_ConsumerGroup",
		&waitGroup,
		controller.OfferBehaviorM.AddOfferBehaviorData)

	go 	controller.Kafka.Consume("ChurnPredictionResultModel",
		"ChurnPredictionMlResultModel_ConsumerGroup",
		&waitGroup,
		controller.CPMResultM.AddChurnPredictionMlResultData)

	//go 	controller.Kafka.Consume("ChurnBlockerResultModel",
	//	"ChurnBlockerMlResultModel_ConsumerGroup",
	//	&waitGroup,
	//	controller.CBMlResultM.AddChurnBlockerMlResultData)

	waitGroup.Wait()
}


