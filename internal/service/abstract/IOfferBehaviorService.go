package abstract

type IOfferBehaviorService interface{
	AddOfferBehaviorData(data *[]byte)(success bool,message string)
}
