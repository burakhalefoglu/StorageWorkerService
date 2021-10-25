package abstract

type IScreenSwipeService interface{
	AddScreenSwipeData(data *[]byte)(success bool,message string)
}

