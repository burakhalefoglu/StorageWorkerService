package abstract

type IChurnBlockerMlResultService interface{
	AddChurnBlockerMlResultData(data *[]byte)(success bool,message string)
}
