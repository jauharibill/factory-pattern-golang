package Shapes

type ErrorMessage struct {
	Message string
}

func (errorMessage ErrorMessage) Draw() string {
	return errorMessage.Message
}
