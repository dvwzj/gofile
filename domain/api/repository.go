package api

import (
	"github.com/dvwzj/gofile/entity"
	"github.com/dvwzj/gofile/params"
)

type Repository interface {
	// GET
	// https://api.gofile.io/servers
	GetServers() (*entity.Response[entity.Servers], error)

	// POST
	// https://{server}.gofile.io/contents/uploadfile
	UploadFile(file params.UploadFile, options ...params.UploadFileOption) (*entity.Response[entity.UploadedFile], error)

	// POST
	// https://api.gofile.io/contents/createFolder
	CreateFolder(parentFolderId string, options ...params.CreateFolderOption) (*entity.Response[entity.CreatedFolder], error)

	// PUT
	// https://api.gofile.io/contents/{contentId}/update
	UpdateContent(contentId string, option params.UpdateContentOption) (*entity.EmptyDataResponse, error)

	// DELETE
	// https://api.gofile.io/contents
	DeleteContents(contentsId []string) (*entity.Response[map[string]entity.EmptyDataResponse], error)

	// DELETE
	// https://api.gofile.io/contents/{contentId}
	DeleteContent(contentId string) (*entity.EmptyDataResponse, error)

	// GET
	// https://api.gofile.io/contents/{contentId}
	GetContent(contentId string) (*entity.Response[entity.Content], error)

	// POST
	// https://api.gofile.io/contents/{contentId}/directlinks
	CreateDirectLink(contentId string, directLink entity.DirectLink) (*entity.Response[entity.DirectLink], error)

	// PUT
	// https://api.gofile.io/contents/{contentId}/directlinks/{directLinkId}
	UpdateDirectLink(contentId, directLinkId string, directLink entity.DirectLink) (*entity.Response[entity.DirectLink], error)

	// DELETE
	// https://api.gofile.io/contents/{contentId}/directlinks/{directLinkId}
	DeleteDirectLink(contentId, directLinkId string) (*entity.EmptyDataResponse, error)

	// POST
	// https://api.gofile.io/contents/copy
	CopyContents(folderId string, contentsId []string) (*entity.EmptyDataResponse, error)

	// POST
	// https://api.gofile.io/contents/{contentId}/copy
	CopyContent(folderId, contentId string) (*entity.EmptyDataResponse, error)

	// PUT
	// https://api.gofile.io/contents/move
	MoveContents(folderId string, contentsId []string) (*entity.EmptyDataResponse, error)

	// PUT
	// https://api.gofile.io/contents/{contentId}/move
	MoveContent(folderId string, contentId string) (*entity.EmptyDataResponse, error)

	// GET
	// https://api.gofile.io/accounts/getid
	GetAccountId() (*entity.Response[entity.GetId], error)

	// GET
	// https://api.gofile.io/accounts/{accountId}
	GetAccount(accountId string) (*entity.Response[entity.Account], error)

	// POST
	// https://api.gofile.io/accounts/{accountId}/resettoken
	ResetAccountToken(accountId string) (*entity.EmptyDataResponse, error)
}
