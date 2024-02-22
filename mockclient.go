package telegramclient

type MockClient struct{}

func NewMockClient() *MockClient {
	return &MockClient{}
}

func (c MockClient) SendMessage(_ int64, _ MessageStruct) error {
	return nil
}
