# gofile
A gofile.io API Wrapper Client for golang

# Installation
```bash
go get github.com/dvwzj/gofile
```

# Example
```go
package main

import (
	"github.com/dvwzj/gofile"
	"github.com/dvwzj/gofile/params"
)

func main() {
    client, err := gofile.NewClient(gofile.WithToken("your-token-here"))
    if err != nil {
        panic(err)
    }
    account, err := client.GetAccount()
    if err != nil {
        panic(err)
    }
    folder, err := client.CreateFolder(account.RootFolder, params.WithFolderName("new-folder"))
    if err != nil {
        panic(err)
    }
    err = client.UpdateContent(folder.FolderId, params.WithPublic(true))
    if err != nil {
        panic(err)
    }
    file, err := os.Open("path/to/file")
    if err != nil {
        panic(err)
    }
    defer file.Close()
    uploadedFile, err := client.UploadFile(params.WithFile(file), params.WithFolderId(folder.FolderId))
    if err != nil {
        panic(err)
    }
    err = client.UpdateContent(uploadedFile.FileId, params.WithName("new-file-name"))
    if err != nil {
        panic(err)
    }
    err = client.DeleteContent(uploadedFile.ParentFolder)
    if err != nil {
        panic(err)
    }
}
```

# Usage

## Client

```go
client, err := gofile.NewClient()
// This will return a client without setting token,
// You can upload a file but unable to use other features cause no token provided.
```
```go
client, err := gofile.NewClient(gofile.WithToken("your-token-here"))
// This will return a client with your token,
// You can use any features as long as you have permissions (with your tier).
```
More details about permissions: [API](https://gofile.io/api)

### Server

```go
servers, err := client.GetServers()
// This will return a list of servers
/**
servers = [
    {
        Name   string,
        Zone   string,
    },
    ...
]
*/
```

### Content

#### Upload file

```go
// Attach file with a file system (*os.File):
uploadedFile, err := client.UploadFile(params.WithFile(file))
// This will upload a file with auto select a server, create a new folder and push this file into.

// To select a folder by your self:
uploadedFile, err := client.UploadFile(params.WithFile(file), params.WithFolderId("your-folder-id"))

// To override file name by your self:
uploadedFile, err := client.UploadFile(params.WithFile(file), params.WithFileName("your-file-name"))

// To select a select by your self:
uploadedFile, err := client.UploadFile(params.WithFile(file), params.WithServerName("store1"))
// The server url will be "https://store1.gofile.io/contents/uploadfile"

// Some time you may want to specific all options.
uploadedFile, err := client.UploadFile(
    params.WithFile(file), 
    params.WithServerName("store1"), 
    params.WithFolderId("your-folder-id"), 
    params.WithFileName("your-file-name"),
)
```
```go
// Attach file with a reader (id.Reader):
uploadedFile, err := client.UploadFile(params.WithReader(fileReader, "your-file-name"))
// params.WithFolderId, params.WithServerName and params.WithFileName are available too.
```
```go
// Attach file with a file path (or url):
uploadedFile, err := client.UploadFile(params.WithPath("path/to/file"))
// params.WithFolderId, params.WithServerName and params.WithFileName are available too.
```
```go
// Attach file with bytes ([]byte):
uploadedFile, err := client.UploadFile(params.WithBytes([]byte("your-content"), "your-file-name"))
// params.WithFolderId, params.WithServerName and params.WithFileName are available too.
```

#### Create folder

```go
createdFolder, err := client.CreatedFolder("parent-folder-id")
// This will create a folder with a generated name (from API) inside your parent folder.

// To specific your new folder name:
createdFolder, err := client.CreatedFolder("parent-folder-id", params.WithFolderName("your-folder-name"))
```

#### Update content

```go
// To update the content name:
err := client.UpdateContent("content-id", params.WithName("your-new-content-name")) // File or Folder

// To update the content description:
err := client.UpdateContent("content-id", params.WithDescription("your-new-content-description")) // Folder only

// To update the content tags:
err := client.UpdateContent("content-id", params.WithTags([]string{"your-content-tag-1", "your-content-tag-2"})) // Folder only

// To update the content public or not:
err := client.UpdateContent("content-id", params.WithPublic(false)) // folder only

// To update the content expiration:
err := client.UpdateContent("content-id", params.WithExpiry(time.Now().Add(time.Hour * 24 * 7))) // Folder only

// To update the content password:
err := client.UpdateContent("content-id", params.WithPassword("your-password")) // Folder only
```

#### Update content

```go
// To delete a single content:
err := client.DeleteContent("content-id")
```
```go
// To delete a multiple contents:
err := client.DeleteContents([]string{"content-id-1", "content-id-2"})
```

#### Get content

```go
content, err := client.GetContent("content-id")
/**
content = {
    Id                  string
    Type                string
    Name                string
    ParentFolder        string
    Code                string
    CreateTime          string
    Public              bool
    TotalDownloadCount  int
    TotalSize           int
    ChildrenIds         []string
    Children            *map[string]UniversalContent // File Or Folder
    IsOwner             *bool
    IsRoot              *bool
}
*/

// We have a helper function for your content children:
content.Children.Files() // Return []ChildContentFile
/**
childContentFile = {
    Id              string
    Type            string
    Name            string
    CreateTime      string
    Size            int
    DownloadCount   int
    MD5             string
    Mimetype        string
    ServerSelected  string
    Link            string
    DirectLinks     *map[string]DirectLink
}
*/
content.Children.Folders() // Return []ChildContentFolder
/**
childContentFolder = {
    Id              string
    Type            string
    Name            string
    ParentFolder    string
    Code            string
    CreateTime      int
    Public          bool
    ChildrenIds     []string
}
*/
```

#### Direct link

```go
// To create a direct link:
directLink, err := client.CreateDirectLink("content-id", entity.DirectLink{
    Auth             []string{"user:pass", "root:admin"}, // Example
	DomainsAllowed   []string{"example.com", "test.com"}, // Example
	ExpireTime       1735689600,                          // Example
	SourceIpsAllowed []string{"127.0.0.1", "127.0.0.2"},  // Example
})
```
```go
// To update a direct link:
directLink, err := client.UpdateDirectLink("content-id", "direct-link-id", entity.DirectLink{
    Auth             []string{"user:pass", "root:admin"}, // Example
	DomainsAllowed   []string{"example.com", "test.com"}, // Example
	ExpireTime       1735689600,                          // Example
	SourceIpsAllowed []string{"127.0.0.1", "127.0.0.1"},  // Example
})
```
```go
// To delete a direct link:
err := client.DeleteDirectLink("content-id", "direct-link-id")
```

#### Copy a content

```go
// Single content:
err := client.CopyContent("folder-id", "content-id")
```
```go
// Multiple contents:
err := client.CopyContents("folder-id", []string{"content-id-1", "content-id-2"})
```

#### Move a content

```go
// Single content:
err := client.MoveContent("folder-id", "content-id")
```
```go
// Multiple contents:
err := client.MoveContents("folder-id", []string{"content-id-1", "content-id-2"})
```

## Account

```go
// To get your account id:
accountId, err := client.GetAccoutId()
```
```go
// To get your account information:
account, err := gofile.GetAccout()
/**
account = {
    Id              string
    Email           string
    Tier            string
    Token           string
    RootFolder      string
    StatsCurrent    struct {
        FileCount               int
        FolderCount             int
        Storage                 int
        TrafficWebDownloaded    *int
    }
    Icon            *string
}
*/
```
```go
// To reset your token:
err := gofile.ResetAccountToken()
// Sending login url to your email
```