package IoC

import (
	KafkaInsertController "StorageWorkerService/internal/controller/kafka"
	"StorageWorkerService/internal/repository/concrete/mongodb_driver"
	ManagerConcrete "StorageWorkerService/internal/service/concrete"
	"StorageWorkerService/pkg/database/mongodb"
	"StorageWorkerService/pkg/jsonParser/gojson"
	"StorageWorkerService/pkg/kafka/kafkago"
)

var InsertKafkaController = &KafkaInsertController.InsertController{
	Kafka: &kafkago.KafkaGo{},

	AdvM: &ManagerConcrete.AdvEventManager{
		Parser: &gojson.GoJson{},
		AdvEventDal: &mongodb_driver.MDbDAdvEventDal{
			Client: mongodb.Conn.Conn,
		},
	},

	BuyingM: &ManagerConcrete.BuyingEventManager{
		Parser: &gojson.GoJson{},
		BuyingEventDal: &mongodb_driver.MDbDBuyingEventDal{
			Client: mongodb.Conn.Conn,
		},
		ClientService: &ManagerConcrete.ClientManager{
			Parser:    &gojson.GoJson{},
			ClientDal:  &mongodb_driver.MDbDClientDal{
				Client: mongodb.Conn.Conn,
			},
		},
	},

	HardwareM: &ManagerConcrete.HardwareManager{
		Parser: &gojson.GoJson{},
		HardwareDal: &mongodb_driver.MDbDHardwareDal{
			Client: mongodb.Conn.Conn,
		},
	},

	LocM : &ManagerConcrete.LocationManager{
		Parser: &gojson.GoJson{},
		LocationDal: &mongodb_driver.MDbDLocationDal{
			Client: mongodb.Conn.Conn,
		},
	},

	ScrSwipeM : &ManagerConcrete.ScreenSwipeManager{
		Parser: &gojson.GoJson{},
		ScreenSwipeDal: &mongodb_driver.MDbDScreenSwipeDal{
			Client: mongodb.Conn.Conn,
		},
	},

	ScrClickM : &ManagerConcrete.ScreenClickManager{
		Parser: &gojson.GoJson{},
		ScreenClickDal: &mongodb_driver.MDbDScreenClickDal{
			Client: mongodb.Conn.Conn,
		},
	},

	LvlBaseSessionM : &ManagerConcrete.LevelBaseSessionManager{
		Parser: &gojson.GoJson{},
		LevelBaseSessionDal: &mongodb_driver.MDbDLevelBaseSessionDal{
			Client: mongodb.Conn.Conn,
		},
	},

	GameSessionM : &ManagerConcrete.GameSessionManager{
		Parser: &gojson.GoJson{},
		GameSessionDal: &mongodb_driver.MDbDGameSessionDal{
			Client: mongodb.Conn.Conn,
		},
	},

	EnemyBaseLoginLM : &ManagerConcrete.EnemyBaseLoginLevelManager{
		Parser: &gojson.GoJson{},
		EnemyBaseLoginLevelDal: &mongodb_driver.MDbDEnemyBaseLoginLevelDal{
			Client: mongodb.Conn.Conn,
		},
	},

	EnemyBaseLevelFM : &ManagerConcrete.EnemyBaseLevelFailManager{
		Parser: &gojson.GoJson{},
		EnemyBaseLevelFailDal: &mongodb_driver.MDbDEnemyBaseLevelFailDal{
			Client: mongodb.Conn.Conn,
		},
	},

	ManuelFlowM : &ManagerConcrete.ManuelFlowManager{
		Parser: &gojson.GoJson{},
		ManuelFlowDal: &mongodb_driver.MDbDManuelFlowDal{
			Client: mongodb.Conn.Conn,
		},
	},

	OfferBehaviorM : &ManagerConcrete.OfferBehaviorManager{
		Parser: &gojson.GoJson{},
		OfferBehaviorDal: &mongodb_driver.MDbDOfferBehaviorDal{
			Client: mongodb.Conn.Conn,
		},
	},

	CBMlResultM : &ManagerConcrete.ChurnBlockerMlResultManager{
		Parser: &gojson.GoJson{},
		ChurnBlockerMlResultDal: &mongodb_driver.MDbDChurnBlockerMlResultDal{
			Client: mongodb.Conn.Conn,
		},
	},

	CPMResultM : &ManagerConcrete.ChurnPredictionMlResultManager{
		Parser: &gojson.GoJson{},
		ChurnPredictionMlResultDal: &mongodb_driver.MDbDChurnPredictionMlResultDal{
			Client: mongodb.Conn.Conn,
		},
	},

	InventoryM : &ManagerConcrete.InventoryManager{
		Parser: &gojson.GoJson{},
		InventoryDal: &mongodb_driver.MDbDInventoryDal{
			Client: mongodb.Conn.Conn,
		},
	}}
