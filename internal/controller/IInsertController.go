package controller

type IInsertController interface {
	StartListen()
}

func StartInsertListener(listener IInsertController){
	listener.StartListen()
}