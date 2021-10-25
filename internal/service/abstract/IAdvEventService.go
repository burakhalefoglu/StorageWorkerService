package abstract

type IAdvEventService interface{
	AddAdvEventData(data *[]byte)(success bool,message string)
}
