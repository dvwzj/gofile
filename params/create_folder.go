package params

type CreateFolderParams struct {
	FolderName *string
}

func (p CreateFolderParams) Body(parentFolderId string) map[string]interface{} {
	data := map[string]interface{}{
		"parentFolderId": parentFolderId,
	}
	if p.FolderName != nil {
		data["folderName"] = *p.FolderName
	}
	return data
}

type CreateFolderOption func(*CreateFolderParams)

func WithFolderName(folderName string) CreateFolderOption {
	return func(params *CreateFolderParams) {
		params.FolderName = &folderName
	}
}
