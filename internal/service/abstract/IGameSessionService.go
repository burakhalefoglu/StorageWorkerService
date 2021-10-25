package abstract

type IGameSessionService interface{
	AddGameSessionData(data *[]byte)(success bool,message string)
}

