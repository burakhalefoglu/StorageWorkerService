package abstract

type IEnemyBaseLevelFailService interface{
	AddEnemyBaseLevelFailData(data *[]byte)(success bool,message string)
}
