package gofile

import (
	"github.com/dvwzj/gofile/entity"
	"github.com/go-resty/resty/v2"
)

func NewGuestAccount() (*entity.CreatedAccount, error) {
	createdAccount, err := RegisterNewAccount("")
	if err != nil {
		return nil, err
	}
	return createdAccount, nil
}

func RegisterNewAccount(email string) (*entity.CreatedAccount, error) {
	httpCLient := resty.New()
	req := httpCLient.
		R().
		SetError(entity.Response[entity.CreatedAccount]{}).
		SetResult(entity.Response[entity.CreatedAccount]{})
	if email != "" {
		req.SetBody(map[string]interface{}{
			"email": email,
		})
	}
	resp, err := req.Post("https://api.gofile.io/accounts")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(*entity.Response[entity.CreatedAccount]).Error()
	}
	if createdAccount, ok := resp.Result().(*entity.Response[entity.CreatedAccount]); ok {
		return &createdAccount.Data, nil
	}
	return nil, entity.ErrAccount
}
