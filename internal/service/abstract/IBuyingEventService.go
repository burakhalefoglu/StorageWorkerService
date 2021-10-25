package abstract

type IAdvBuyingService interface{
	AddBuyingEventData(data *[]byte)(success bool,message string)
}