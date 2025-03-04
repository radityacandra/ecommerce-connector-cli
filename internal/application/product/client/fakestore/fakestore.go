package fakestore

type Fakestore struct {
	client *Client
}

func NewFakestore(baseUrl string) *Fakestore {
	client, _ := NewClient(baseUrl)
	return &Fakestore{
		client: client,
	}
}
