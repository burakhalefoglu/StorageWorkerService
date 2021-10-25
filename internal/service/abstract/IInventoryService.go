package abstract

type IInventoryService interface{
	AddInventoryData(data *[]byte)(success bool,message string)
}

