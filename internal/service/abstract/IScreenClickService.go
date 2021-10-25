package abstract

type IScreenClickService interface{
	AddScreenClickData(data *[]byte)(success bool,message string)
}

