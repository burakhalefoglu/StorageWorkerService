package abstract

type IBuyingEventService interface{
	AddBuyingEventData(data *[]byte)(success bool,message string)
}