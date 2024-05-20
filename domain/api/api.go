package api

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"

	"github.com/dvwzj/gofile/entity"
	"github.com/dvwzj/gofile/params"
	"github.com/go-resty/resty/v2"
)

type API interface {
	HttpClient() *resty.Client
	Repository
}

type Domain struct {
	httpClient *resty.Client
}

func (d *Domain) HttpClient() *resty.Client {
	return d.httpClient
}

func (d Domain) GetServers() (*entity.Response[entity.Servers], error) {
	resp, err := d.httpClient.R().
		SetError(entity.Response[entity.Servers]{}).
		SetResult(entity.Response[entity.Servers]{}).
		Get("/servers")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(*entity.Response[entity.Servers]).Error()
	}
	return resp.Result().(*entity.Response[entity.Servers]), nil
}

func (d Domain) UploadFile(file params.UploadFile, options ...params.UploadFileOption) (*entity.Response[entity.UploadedFile], error) {
	params := &params.UploadFileParams{}
	if err := file(params); err != nil {
		return nil, err
	}
	for _, option := range options {
		if err := option(params); err != nil {
			return nil, err
		}
	}
	server := ""
	if params.Server != nil {
		server = *params.Server
	} else {
		serversResp, err := d.GetServers()
		if err != nil {
			return nil, err
		}
		if len(serversResp.Data.Servers) == 0 {
			return nil, fmt.Errorf("no server available")
		}
		avaiableServers := make(chan string, len(serversResp.Data.Servers))
		errCh := make(chan error, len(serversResp.Data.Servers))
		wg := sync.WaitGroup{}
		for _, server := range serversResp.Data.Servers {
			wg.Add(1)
			go func(server string) {
				defer wg.Done()
				resp, err := d.httpClient.R().Head(fmt.Sprintf("https://%s.gofile.io", server))
				if err != nil {
					errCh <- err
					return
				}
				if resp.IsError() {
					errCh <- resp.Error().(error)
					return
				}
				errCh <- nil
				avaiableServers <- server
			}(server.Name)
		}
		wg.Wait()
		close(avaiableServers)
		close(errCh)
		for err := range errCh {
			if err != nil {
				return nil, err
			}
		}
		servers := []string{}
		for s := range avaiableServers {
			servers = append(servers, s)
		}
		if len(servers) == 0 {
			return nil, fmt.Errorf("no server available")
		}
		rand.New(rand.NewSource(0)).Shuffle(len(servers), func(i, j int) {
			servers[i], servers[j] = servers[j], servers[i]
		})
		server = servers[0]
	}
	req := d.httpClient.R().
		SetError(entity.Response[entity.UploadedFile]{}).
		SetResult(entity.Response[entity.UploadedFile]{}).
		SetFileReader("file", *params.FileName, params.FileReader)
	if params.FolderId != nil {
		req.SetFormData(map[string]string{
			"folderId": *params.FolderId,
		})
	}
	resp, err := req.Post(fmt.Sprintf("https://%s.gofile.io/contents/uploadfile", server))
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(*entity.Response[entity.UploadedFile]).Error()
	}
	return resp.Result().(*entity.Response[entity.UploadedFile]), nil
}

func (d Domain) CreateFolder(parentFolderId string, options ...params.CreateFolderOption) (*entity.Response[entity.CreatedFolder], error) {
	params := &params.CreateFolderParams{}
	for _, option := range options {
		option(params)
	}
	resp, err := d.httpClient.R().
		SetError(entity.Response[entity.CreatedFolder]{}).
		SetResult(entity.Response[entity.CreatedFolder]{}).
		SetBody(params.Body(parentFolderId)).
		Post("/contents/createFolder")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(*entity.Response[entity.CreatedFolder]).Error()
	}
	return resp.Result().(*entity.Response[entity.CreatedFolder]), nil
}

func (d Domain) UpdateContent(contentId string, option params.UpdateContentOption) (*entity.EmptyDataResponse, error) {
	params := &params.UpdateContentParams{}
	option(params)
	if params.Attribute == "" {
		return nil, fmt.Errorf("no attribute provided")
	}
	resp, err := d.httpClient.R().
		SetError(entity.EmptyDataResponse{}).
		SetResult(entity.EmptyDataResponse{}).
		SetBody(params.Body()).
		Put(fmt.Sprintf("/contents/%s/update", contentId))
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(*entity.EmptyDataResponse).Error()
	}
	return resp.Result().(*entity.EmptyDataResponse), nil
}

func (d Domain) DeleteContents(contentsId []string) (*entity.Response[map[string]entity.EmptyDataResponse], error) {
	resp, err := d.httpClient.R().
		SetError(entity.Response[map[string]entity.EmptyDataResponse]{}).
		SetResult(entity.Response[map[string]entity.EmptyDataResponse]{}).
		SetBody(map[string]interface{}{
			"contentsId": strings.Join(contentsId, ","),
		}).
		Delete("/contents")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(*entity.Response[map[string]entity.EmptyDataResponse]).Error()
	}
	return resp.Result().(*entity.Response[map[string]entity.EmptyDataResponse]), nil
}

func (d Domain) DeleteContent(contentId string) (*entity.EmptyDataResponse, error) {
	resp, err := d.httpClient.R().
		SetError(entity.EmptyDataResponse{}).
		SetResult(entity.EmptyDataResponse{}).
		Delete(fmt.Sprintf("/contents/%s", contentId))
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(*entity.EmptyDataResponse).Error()
	}
	return resp.Result().(*entity.EmptyDataResponse), nil
}

func (d Domain) GetContent(contentId string) (*entity.Response[entity.Content], error) {
	resp, err := d.httpClient.R().
		SetError(entity.Response[interface{}]{}).
		SetResult(entity.Response[interface{}]{}).
		Get(fmt.Sprintf("/contents/%s", contentId))
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(*entity.Response[interface{}]).Error()
	}
	result := resp.Result().(*entity.Response[interface{}])
	content := entity.Content{}
	if err := content.Unmarshal(result.Data); err != nil {
		return nil, err
	}
	return &entity.Response[entity.Content]{
		Status: result.Status,
		Data:   content,
	}, nil
}

func (d Domain) CreateDirectLink(contentId string, directLink entity.DirectLink) (*entity.Response[entity.DirectLink], error) {
	resp, err := d.httpClient.R().
		SetError(entity.Response[entity.DirectLink]{}).
		SetResult(entity.Response[entity.DirectLink]{}).
		SetBody(directLink).
		Post(fmt.Sprintf("/contents/%s/directlinks", contentId))
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(*entity.Response[entity.DirectLink]).Error()
	}
	return resp.Result().(*entity.Response[entity.DirectLink]), nil
}

func (d Domain) UpdateDirectLink(contentId, directLinkId string, directLink entity.DirectLink) (*entity.Response[entity.DirectLink], error) {
	resp, err := d.httpClient.R().
		SetError(entity.Response[entity.DirectLink]{}).
		SetResult(entity.Response[entity.DirectLink]{}).
		SetBody(directLink).
		Put(fmt.Sprintf("/contents/%s/directlinks/%s", contentId, directLinkId))
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(*entity.Response[entity.DirectLink]).Error()
	}
	return resp.Result().(*entity.Response[entity.DirectLink]), nil
}

func (d Domain) DeleteDirectLink(contentId, directLinkId string) (*entity.EmptyDataResponse, error) {
	resp, err := d.httpClient.R().
		SetError(entity.EmptyDataResponse{}).
		SetResult(entity.EmptyDataResponse{}).
		Delete(fmt.Sprintf("/contents/%s/directlinks/%s", contentId, directLinkId))
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(*entity.Response[entity.EmptyDataResponse]).Error()
	}
	return resp.Result().(*entity.EmptyDataResponse), nil
}

func (d Domain) CopyContents(folderId string, contentsId []string) (*entity.EmptyDataResponse, error) {
	resp, err := d.httpClient.R().
		SetError(entity.EmptyDataResponse{}).
		SetResult(entity.EmptyDataResponse{}).
		SetBody(map[string]interface{}{
			"folderId":   folderId,
			"contentsId": strings.Join(contentsId, ","),
		}).
		Post("/contents/copy")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(*entity.EmptyDataResponse).Error()
	}
	return resp.Result().(*entity.EmptyDataResponse), nil
}

func (d Domain) CopyContent(folderId, contentId string) (*entity.EmptyDataResponse, error) {
	resp, err := d.httpClient.R().
		SetError(entity.EmptyDataResponse{}).
		SetResult(entity.EmptyDataResponse{}).
		SetBody(map[string]interface{}{
			"folderId": folderId,
		}).
		Post(fmt.Sprintf("/contents/%s/copy", contentId))
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(*entity.EmptyDataResponse).Error()
	}
	return resp.Result().(*entity.EmptyDataResponse), nil
}

func (d Domain) MoveContents(folderId string, contentsId []string) (*entity.EmptyDataResponse, error) {
	resp, err := d.httpClient.R().
		SetError(entity.EmptyDataResponse{}).
		SetResult(entity.EmptyDataResponse{}).
		SetBody(map[string]interface{}{
			"folderId":   folderId,
			"contentsId": strings.Join(contentsId, ","),
		}).
		Put("/contents/move")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(*entity.EmptyDataResponse).Error()
	}
	return resp.Result().(*entity.EmptyDataResponse), nil
}

func (d Domain) MoveContent(folderId, contentId string) (*entity.EmptyDataResponse, error) {
	resp, err := d.httpClient.R().
		SetError(entity.EmptyDataResponse{}).
		SetResult(entity.EmptyDataResponse{}).
		SetBody(map[string]interface{}{
			"folderId": folderId,
		}).
		Put(fmt.Sprintf("/contents/%s/move", contentId))
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(*entity.EmptyDataResponse).Error()
	}
	return resp.Result().(*entity.EmptyDataResponse), nil
}

func (d Domain) GetAccountId() (*entity.Response[entity.GetId], error) {
	resp, err := d.httpClient.R().
		SetError(entity.Response[entity.GetId]{}).
		SetResult(entity.Response[entity.GetId]{}).
		Get("/accounts/getid")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(*entity.Response[entity.GetId]).Error()
	}
	return resp.Result().(*entity.Response[entity.GetId]), nil
}

func (d Domain) GetAccount(accountId string) (*entity.Response[entity.Account], error) {
	resp, err := d.httpClient.R().
		SetError(entity.Response[entity.Account]{}).
		SetResult(entity.Response[entity.Account]{}).
		Get(fmt.Sprintf("/accounts/%s", accountId))
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(*entity.Response[entity.Account]).Error()
	}
	return resp.Result().(*entity.Response[entity.Account]), nil
}

func (d Domain) ResetAccountToken(accountId string) (*entity.EmptyDataResponse, error) {
	resp, err := d.httpClient.R().
		SetError(entity.EmptyDataResponse{}).
		SetResult(entity.EmptyDataResponse{}).
		Post(fmt.Sprintf("/accounts/%s/resettoken", accountId))
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(*entity.EmptyDataResponse).Error()
	}
	return resp.Result().(*entity.EmptyDataResponse), nil
}

func NewAPI() API {
	return &Domain{
		httpClient: resty.New().SetBaseURL("https://api.gofile.io"),
	}
}
