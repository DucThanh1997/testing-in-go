// package main

// import (
// 	"backend/app"
// )


package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

func main() {
	fmt.Println("alooo")
}

type MessageService interface {
	SendMessageToUser(string) bool
}

type MyService struct {
	messageService MessageService
}

// smsServiceMock để giả định cho một MessageService
type smsServiceMock struct {
	mock.Mock
}

// và đương nhiên là chúng ta phải implement lại service MessageService để giả định cho việc gửi message tới user
// như vậy thay vì phải gọi tới SMSService để gửi message thì bây giờ chúng ta sẽ gọi smsServiceMock
func (m *smsServiceMock) SendMessageToUser(message string) bool {
	fmt.Println("Mocked function")
  return true
}

func TestSendMessage(t *testing.T) {
	// new một smsServiceMock để thay thế cho SMSService
	smsService := new(smsServiceMock)

  	// giả định cho một hàm
	smsService.On("SendMessageToUser", "hello").Return(true)

	// sử dụng smsServiceMock để giả định
	myService := MyService{smsService}
	// hàm chính 
	myService.messageService.SendMessageToUser("hello")

	smsService.AssertExpectations(t)
}