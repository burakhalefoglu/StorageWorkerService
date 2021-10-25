package abstract

type IManuelFlowService interface{
	AddManuelFlowData(data *[]byte)(success bool,message string)
}
