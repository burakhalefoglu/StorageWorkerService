package abstract

type IChurnPredictionMlResultService interface{
	AddChurnPredictionMlResultData(data *[]byte)(success bool,message string)
}
