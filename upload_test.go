package gofile_test

import (
	"testing"

	"github.com/dvwzj/gofile"
	"github.com/dvwzj/gofile/entity"
	"github.com/dvwzj/gofile/params"
)

func TestUpload(t *testing.T) {
	client, err := gofile.NewClient()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if client == nil {
		t.Fatalf("unexpected nil client")
	}
	token := client.GetToken()
	if token != "" {
		t.Fatal("unexpected token", token)
	}
	uplodedFile, err := client.UploadFile(params.WithBytes([]byte("ok"), "test.txt")) // upload as guest (without token)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if uplodedFile == nil {
		t.Fatalf("unexpected nil uploaded file")
	}
	err = client.DeleteContent(uplodedFile.FileId) // not able to delete as guest (without token)
	if err == nil {
		t.Fatalf("unexpected nil error")
	}
	if err != entity.ErrToken {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestAsGuest(t *testing.T) {
	client, err := gofile.NewClient(gofile.WithNewGuestAccount)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if client == nil {
		t.Fatalf("unexpected nil client")
	}
	token := client.GetToken()
	if token == "" {
		t.Fatalf("unexpected empty token")
	}
	uplodedFile, err := client.UploadFile(params.WithBytes([]byte("ok"), "test.txt")) // upload as guest (with token)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if uplodedFile == nil {
		t.Fatalf("unexpected nil uploaded file")
	}
	err = client.DeleteContent(uplodedFile.FileId) // able to delete as guest (with token)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestUploadWithToken(t *testing.T) {
	client, err := gofile.NewClient(gofile.WithToken("7JyBbtDTF7yakfcfmTWYlXNTL4j5r9HV"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if client == nil {
		t.Fatalf("unexpected nil client")
	}
	uplodedFile, err := client.UploadFile(params.WithPath("https://gofile.io/dist/img/logo-small-70.png"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if uplodedFile == nil {
		t.Fatalf("unexpected nil uploaded file")
	}
	accountId, err := client.GetAccountId()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if accountId == "" {
		t.Fatalf("unexpected empty account id")
	}
	account, err := client.GetAccount()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	content, err := client.GetContent(uplodedFile.ParentFolder)
	if account.Tier == entity.AccountTierPremium {
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		err = client.UpdateContent(uplodedFile.FileId, params.WithName("new-name.txt"))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		newFolder1, err := client.CreateFolder(content.Id, params.WithFolderName("new-folder-1"))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		newFolder2, err := client.CreateFolder(newFolder1.FolderId, params.WithFolderName("new-folder-2"))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		err = client.CopyContent(newFolder2.FolderId, uplodedFile.FileId)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		err = client.CopyContents(newFolder1.FolderId, []string{uplodedFile.FileId})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		err = client.MoveContent(newFolder1.FolderId, uplodedFile.FileId)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		err = client.MoveContents(newFolder2.FolderId, []string{uplodedFile.FileId})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		err = client.DeleteContent(uplodedFile.FileId)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		err = client.DeleteContent(newFolder2.FolderId)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		_, err = client.DeleteContents([]string{newFolder1.FolderId, uplodedFile.FileId, uplodedFile.ParentFolder})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	} else {
		if err != entity.ErrNotPremium {
			t.Fatalf("unexpected error: %v", err)
		}
		err = client.UpdateContent(uplodedFile.FileId, params.WithName("new-name.txt"))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		err = client.DeleteContent(uplodedFile.FileId)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}
}
