package abstract

type IChurnPredictionSuccessRateService interface {
	AddChurnPredictionSuccessRate(data *[]byte) (success bool, message string)
}
