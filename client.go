package gofile

import (
	"github.com/dvwzj/gofile/entity"
	"github.com/dvwzj/gofile/services"
	"github.com/go-resty/resty/v2"
)

type Client interface {
	HttpClient() *resty.Client
	GetToken() string
	services.Service
}

type Gofile struct {
	services.Service
}

func (g *Gofile) HttpClient() *resty.Client {
	return g.Service.HttpClient()
}

func (g *Gofile) GetToken() string {
	return g.Service.HttpClient().Token
}

type ClientOption func(Client) error

func WithToken(token string) ClientOption {
	return func(client Client) error {
		client.HttpClient().SetAuthToken(token)
		return nil
	}
}

func WithAccount(account *entity.Account) ClientOption {
	return func(client Client) error {
		if account == nil {
			return entity.ErrAccount
		}
		if account.Token == "" {
			return entity.ErrToken
		}
		client.HttpClient().SetAuthToken(account.Token)
		return nil
	}
}

func WithNewGuestAccount(client Client) error {
	createdAccount, err := NewGuestAccount()
	if err != nil {
		return err
	}
	client.HttpClient().SetAuthToken(createdAccount.Token)
	return nil
}

func NewClient(options ...ClientOption) (Client, error) {
	client := &Gofile{
		Service: services.NewAPI(),
	}
	for _, option := range options {
		if err := option(client); err != nil {
			return nil, err
		}
	}
	return client, nil
}
