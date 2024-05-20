package entity

import "encoding/json"

const (
	ContentTypeFolder ContentType = "folder"
	ContentTypeFile   ContentType = "file"
)

type ContentType string

type DirectLink struct {
	Auth             []string `json:"auth,omitempty"`
	DomainsAllowed   []string `json:"domainsAllowed,omitempty"`
	ExpireTime       int      `json:"expireTime,omitempty"`
	IsReqLink        bool     `json:"isReqLink,omitempty"`
	SourceIpsAllowed []string `json:"sourceIpsAllowed,omitempty"`
	DirectLink       string   `json:"directLink,omitempty"`
}

type CreatedFolder struct {
	FolderId     string      `json:"folderId"`
	Type         ContentType `json:"type"`
	Name         string      `json:"name"`
	ParentFolder string      `json:"parentFolder"`
	CreateTime   int         `json:"createTime"`
	Code         string      `json:"code"`
}

type UploadedFile struct {
	Code         string `json:"code"`
	DownloadPage string `json:"downloadPage"`
	FileId       string `json:"fileId"`
	FileName     string `json:"fileName"`
	MD5          string `json:"md5"`
	ParentFolder string `json:"parentFolder"`
}

type Content struct {
	Id                 string        `json:"id"`
	Type               ContentType   `json:"type"`
	Name               string        `json:"name"`
	ParentFolder       string        `json:"parentFolder"`
	Code               string        `json:"code"`
	CreateTime         int           `json:"createTime"`
	Public             bool          `json:"public"`
	TotalDownloadCount int           `json:"totalDownloadCount"`
	TotalSize          int           `json:"totalSize"`
	ChildrenIds        []string      `json:"childrenIds"`
	Children           *ChildContent `json:"children,omitempty"`
	IsOwner            *bool         `json:"isOwner,omitempty"`
	IsRoot             *bool         `json:"isRoot,omitempty"`
}

func (c *Content) Unmarshal(v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, c)
}

func (c *Content) Child() ChildContent {
	if c.Children == nil {
		return ChildContent{}
	}
	return *c.Children
}

type ChildContentFolder struct {
	Id           string      `json:"id"`
	Type         ContentType `json:"type"`
	Name         string      `json:"name"`
	ParentFolder string      `json:"parentFolder"`
	Code         string      `json:"code"`
	CreateTime   int         `json:"createTime"`
	Public       bool        `json:"public"`
	ChildrenIds  []string    `json:"childrenIds"`
}

type ChildContentFile struct {
	Id             string                 `json:"id"`
	Type           ContentType            `json:"type"`
	Name           string                 `json:"name"`
	CreateTime     int                    `json:"createTime"`
	Size           int                    `json:"size"`
	DownloadCount  int                    `json:"downloadCount"`
	MD5            string                 `json:"md5"`
	Mimetype       string                 `json:"mimetype"`
	ServerSelected string                 `json:"serverSelected"`
	Link           string                 `json:"link"`
	DirectLinks    *map[string]DirectLink `json:"directLinks,omitempty"`
}

type ChildContent map[string]UniversalContent

func (c *ChildContent) Folders() []ChildContentFolder {
	if c == nil {
		return nil
	}
	contents := []ChildContentFolder{}
	for _, v := range *c {
		if v.Type == nil {
			continue
		}
		if *v.Type == ContentTypeFolder {
			content := ChildContentFolder{}
			if v.Id != nil {
				content.Id = *v.Id
			}
			if v.Type != nil {
				content.Type = *v.Type
			}
			if v.Name != nil {
				content.Name = *v.Name
			}
			if v.ParentFolder != nil {
				content.ParentFolder = *v.ParentFolder
			}
			if v.Code != nil {
				content.Code = *v.Code
			}
			if v.CreateTime != nil {
				content.CreateTime = *v.CreateTime
			}
			if v.Public != nil {
				content.Public = *v.Public
			}
			if v.ChildrenIds != nil {
				content.ChildrenIds = *v.ChildrenIds
			}
			contents = append(contents, content)
		}
	}
	return contents
}

func (c *ChildContent) Files() []ChildContentFile {
	if c == nil {
		return nil
	}
	contents := []ChildContentFile{}
	for _, v := range *c {
		if v.Type == nil {
			continue
		}
		if *v.Type == ContentTypeFile {
			content := ChildContentFile{}
			if v.Id != nil {
				content.Id = *v.Id
			}
			if v.Type != nil {
				content.Type = *v.Type
			}
			if v.Name != nil {
				content.Name = *v.Name
			}
			if v.CreateTime != nil {
				content.CreateTime = *v.CreateTime
			}
			if v.Size != nil {
				content.Size = *v.Size
			}
			if v.DownloadCount != nil {
				content.DownloadCount = *v.DownloadCount
			}
			if v.MD5 != nil {
				content.MD5 = *v.MD5
			}
			if v.Mimetype != nil {
				content.Mimetype = *v.Mimetype
			}
			if v.ServerSelected != nil {
				content.ServerSelected = *v.ServerSelected
			}
			if v.Link != nil {
				content.Link = *v.Link
			}
			if v.DirectLinks != nil {
				content.DirectLinks = v.DirectLinks
			}
		}
	}
	return contents
}

type UniversalContent struct {
	Id                 *string                `json:"id,omitempty"`
	Type               *ContentType           `json:"type,omitempty"`
	Name               *string                `json:"name,omitempty"`
	ParentFolder       *string                `json:"parentFolder,omitempty"`
	Code               *string                `json:"code,omitempty"`
	CreateTime         *int                   `json:"createTime,omitempty"`
	Public             *bool                  `json:"public,omitempty"`
	TotalDownloadCount *int                   `json:"totalDownloadCount,omitempty"`
	TotalSize          *int                   `json:"totalSize,omitempty"`
	ChildrenIds        *[]string              `json:"childrenIds,omitempty"`
	Children           *ChildContent          `json:"children,omitempty"`
	IsOwner            *bool                  `json:"isOwner,omitempty"`
	IsRoot             *bool                  `json:"isRoot,omitempty"`
	Size               *int                   `json:"size,omitempty"`
	DownloadCount      *int                   `json:"downloadCount,omitempty"`
	MD5                *string                `json:"md5,omitempty"`
	Mimetype           *string                `json:"mimetype,omitempty"`
	ServerSelected     *string                `json:"serverSelected,omitempty"`
	Link               *string                `json:"link,omitempty"`
	Thumbnail          *string                `json:"thumbnail,omitempty"`
	DirectLinks        *map[string]DirectLink `json:"directLinks,omitempty"`
	FileId             *string                `json:"fileId,omitempty"`
	FileName           *string                `json:"fileName,omitempty"`
	FolderId           *string                `json:"folderId,omitempty"`
	DownloadPage       *string                `json:"downloadPage,omitempty"`
}

func (c *UniversalContent) Content() Content {
	content := Content{}
	if c == nil {
		return content
	}
	if c.Id != nil {
		content.Id = *c.Id
	}
	if c.Type != nil {
		content.Type = *c.Type
	}
	if c.Name != nil {
		content.Name = *c.Name
	}
	if c.ParentFolder != nil {
		content.ParentFolder = *c.ParentFolder
	}
	if c.Code != nil {
		content.Code = *c.Code
	}
	if c.CreateTime != nil {
		content.CreateTime = *c.CreateTime
	}
	if c.Public != nil {
		content.Public = *c.Public
	}
	if c.TotalDownloadCount != nil {
		content.TotalDownloadCount = *c.TotalDownloadCount
	}
	if c.TotalSize != nil {
		content.TotalSize = *c.TotalSize
	}
	if c.ChildrenIds != nil {
		content.ChildrenIds = *c.ChildrenIds
	}
	if c.Children != nil {
		content.Children = c.Children
	}
	if c.IsOwner != nil {
		content.IsOwner = c.IsOwner
	}
	if c.IsRoot != nil {
		content.IsRoot = c.IsRoot
	}
	return content
}

func (c *UniversalContent) CreatedFolder() CreatedFolder {
	content := CreatedFolder{}
	if c == nil {
		return content
	}
	if c.FolderId != nil {
		content.FolderId = *c.FolderId
	}
	if c.Type != nil {
		content.Type = *c.Type
	}
	if c.Name != nil {
		content.Name = *c.Name
	}
	if c.ParentFolder != nil {
		content.ParentFolder = *c.ParentFolder
	}
	if c.CreateTime != nil {
		content.CreateTime = *c.CreateTime
	}
	if c.Code != nil {
		content.Code = *c.Code
	}
	return content
}

func (c *UniversalContent) UploadedFile() UploadedFile {
	content := UploadedFile{}
	if c == nil {
		return content
	}
	if c.Code != nil {
		content.Code = *c.Code
	}
	if c.DownloadPage != nil {
		content.DownloadPage = *c.DownloadPage
	}
	if c.FileId != nil {
		content.FileId = *c.FileId
	}
	if c.FileName != nil {
		content.FileName = *c.FileName
	}
	if c.MD5 != nil {
		content.MD5 = *c.MD5
	}
	if c.ParentFolder != nil {
		content.ParentFolder = *c.ParentFolder
	}
	return content
}
