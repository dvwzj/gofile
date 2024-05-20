package services

import (
	"github.com/dvwzj/gofile/domain/api"
	"github.com/dvwzj/gofile/entity"
	"github.com/dvwzj/gofile/params"
	"github.com/go-resty/resty/v2"
)

type Service interface {
	HttpClient() *resty.Client
	// GET
	// https://api.gofile.io/servers
	GetServers() (*entity.Servers, error)

	// POST
	// https://{server}.gofile.io/contents/uploadfile
	UploadFile(file params.UploadFile, options ...params.UploadFileOption) (*entity.UploadedFile, error)

	// POST
	// https://api.gofile.io/contents/createFolder
	CreateFolder(parentFolderId string, options ...params.CreateFolderOption) (*entity.CreatedFolder, error)

	// PUT
	// https://api.gofile.io/contents/{contentId}/update
	UpdateContent(contentId string, option params.UpdateContentOption) error

	// DELETE
	// https://api.gofile.io/contents
	DeleteContents(contentsId []string) (*map[string]entity.EmptyDataResponse, error)

	// DELETE
	// https://api.gofile.io/contents/{contentId}
	DeleteContent(contentId string) error

	// GET
	// https://api.gofile.io/contents/{contentId}
	GetContent(contentId string) (*entity.Content, error)

	// POST
	// https://api.gofile.io/contents/{contentId}/directlinks
	CreateDirectLink(contentId string, directLink entity.DirectLink) (*entity.DirectLink, error)

	// PUT
	// https://api.gofile.io/contents/{contentId}/directlinks/{directLinkId}
	UpdateDirectLink(contentId, directLinkId string, directLink entity.DirectLink) (*entity.DirectLink, error)

	// DELETE
	// https://api.gofile.io/contents/{contentId}/directlinks/{directLinkId}
	DeleteDirectLink(contentId, directLinkId string) error

	// POST
	// https://api.gofile.io/contents/copy
	CopyContents(folderId string, contentsId []string) error

	// POST
	// https://api.gofile.io/contents/{contentId}/copy
	CopyContent(folderId, contentId string) error

	// PUT
	// https://api.gofile.io/contents/move
	MoveContents(folderId string, contentsId []string) error

	// PUT
	// https://api.gofile.io/contents/{contentId}/move
	MoveContent(folderId string, contentId string) error

	// GET
	// https://api.gofile.io/accounts/getid
	GetAccountId() (string, error)

	// GET
	// https://api.gofile.io/accounts/{accountId}
	GetAccount() (*entity.Account, error)

	// POST
	// https://api.gofile.io/accounts/{accountId}/resettoken
	ResetAccountToken() error
}

type API struct {
	api.Repository
}

func (a API) HttpClient() *resty.Client {
	return a.Repository.(*api.Domain).HttpClient()
}

func (a API) GetServers() (*entity.Servers, error) {
	resp, err := a.Repository.GetServers()
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

func (a API) UploadFile(file params.UploadFile, options ...params.UploadFileOption) (*entity.UploadedFile, error) {
	resp, err := a.Repository.UploadFile(file, options...)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

func (a API) CreateFolder(parentFolderId string, options ...params.CreateFolderOption) (*entity.CreatedFolder, error) {
	resp, err := a.Repository.CreateFolder(parentFolderId, options...)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

func (a API) UpdateContent(contentId string, option params.UpdateContentOption) error {
	_, err := a.Repository.UpdateContent(contentId, option)
	if err != nil {
		return err
	}
	return nil
}

func (a API) DeleteContents(contentsId []string) (*map[string]entity.EmptyDataResponse, error) {
	resp, err := a.Repository.DeleteContents(contentsId)
	if err != nil {
		return nil, err
	}
	for _, v := range resp.Data {
		if v.Status != "ok" {
			return nil, entity.ErrorResponseStatus(v.Status)
		}
	}
	return &resp.Data, nil
}

func (a API) DeleteContent(contentId string) error {
	_, err := a.Repository.DeleteContent(contentId)
	if err != nil {
		return err
	}
	return nil
}

func (a API) GetContent(contentId string) (*entity.Content, error) {
	resp, err := a.Repository.GetContent(contentId)
	if err != nil {
		return nil, err
	}
	if resp.Data.Id == "" && !resp.Data.Public {
		return nil, entity.ErrorResponseStatus("error-privateContent")
	}
	return &resp.Data, nil
}

func (a API) CreateDirectLink(contentId string, directLink entity.DirectLink) (*entity.DirectLink, error) {
	resp, err := a.Repository.CreateDirectLink(contentId, directLink)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

func (a API) UpdateDirectLink(contentId, directLinkId string, directLink entity.DirectLink) (*entity.DirectLink, error) {
	resp, err := a.Repository.UpdateDirectLink(contentId, directLinkId, directLink)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

func (a API) DeleteDirectLink(contentId, directLinkId string) error {
	_, err := a.Repository.DeleteDirectLink(contentId, directLinkId)
	if err != nil {
		return err
	}
	return nil
}

func (a API) CopyContents(folderId string, contentsId []string) error {
	_, err := a.Repository.CopyContents(folderId, contentsId)
	if err != nil {
		return err
	}
	return nil
}

func (a API) CopyContent(folderId, contentId string) error {
	_, err := a.Repository.CopyContent(folderId, contentId)
	if err != nil {
		return err
	}
	return nil
}

func (a API) MoveContents(folderId string, contentsId []string) error {
	_, err := a.Repository.MoveContents(folderId, contentsId)
	if err != nil {
		return err
	}
	return nil
}

func (a API) MoveContent(folderId string, contentId string) error {
	_, err := a.Repository.MoveContent(folderId, contentId)
	if err != nil {
		return err
	}
	return nil
}

func (a API) GetAccountId() (string, error) {
	resp, err := a.Repository.GetAccountId()
	if err != nil {
		return "", err
	}
	return resp.Data.Id, nil
}

func (a API) GetAccount() (*entity.Account, error) {
	accountId, err := a.GetAccountId()
	if err != nil {
		return nil, err
	}
	resp, err := a.Repository.GetAccount(accountId)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

func (a API) ResetAccountToken() error {
	accountId, err := a.GetAccountId()
	if err != nil {
		return err
	}
	_, err = a.Repository.ResetAccountToken(accountId)
	if err != nil {
		return err
	}
	return nil
}

func NewAPI() Service {
	service := &API{}
	service.Repository = api.NewAPI()
	return service
}
