package abstract

type IAdvStrategyBehaviorService interface {
	AddAdvStrategyBehaviorData(data *[]byte) (success bool, message string)
}
