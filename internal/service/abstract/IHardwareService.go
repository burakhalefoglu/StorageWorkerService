package abstract

type IHardwareService interface{
	AddHardwareData(data *[]byte)(success bool,message string)
}

