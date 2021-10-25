package abstract

type ILevelBaseSessionService interface{
	AddLevelBaseSessionData(data *[]byte)(success bool,message string)
}

