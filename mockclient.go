package telegramclient

type MockClient struct{}

func NewMockClient() *MockClient {
	return &MockClient{}
}

func (c MockClient) SendMessage(chatID int64, messageOut MessageStruct) error {
	return nil
}
