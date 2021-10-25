package abstract

type IEnemyBaseLoginLevelService interface{
	AddEnemyBaseLoginLevelData(data *[]byte)(success bool,message string)
}
