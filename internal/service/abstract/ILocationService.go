package abstract

type ILocationService interface{
	AddLocationData(data *[]byte)(success bool,message string)
}
