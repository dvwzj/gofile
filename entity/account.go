package entity

const (
	AccountTierStandard AccountTier = "standard"
	AccountTierPremium  AccountTier = "premium"
	AccountTierGuest    AccountTier = "guest"
)

type GetId struct {
	Id string `json:"id"`
}

func (a GetId) Ptr() *GetId {
	return &a
}

type CreatedAccount struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}

func (a CreatedAccount) Ptr() *CreatedAccount {
	return &a
}

type AccountTier string

type AccountStatsCurrent struct {
	FileCount            int  `json:"fileCount"`
	FolderCount          int  `json:"folderCount"`
	Storage              int  `json:"storage"`
	TrafficWebDownloaded *int `json:"trafficWebDownloaded,omitempty"`
}

type Account struct {
	Id           string              `json:"id"`
	Email        string              `json:"email"`
	Tier         AccountTier         `json:"tier"`
	Token        string              `json:"token"`
	RootFolder   string              `json:"rootFolder"`
	StatsCurrent AccountStatsCurrent `json:"statsCurrent"`
	Icon         *string             `json:"icon,omitempty"`
}

func (a Account) Ptr() *Account {
	return &a
}

type UniversalAccount struct {
	Id           *string              `json:"id,omitempty"`
	Email        *string              `json:"email,omitempty"`
	Tier         *AccountTier         `json:"tier,omitempty"`
	Token        *string              `json:"token,omitempty"`
	RootFolder   *string              `json:"rootFolder,omitempty"`
	StatsCurrent *AccountStatsCurrent `json:"statsCurrent,omitempty"`
	Icon         *string              `json:"icon,omitempty"`
}

func (a *UniversalAccount) Account() Account {
	account := Account{}
	if a.Id != nil {
		account.Id = *a.Id
	}
	if a.Email != nil {
		account.Email = *a.Email
	}
	if a.Tier != nil {
		account.Tier = *a.Tier
	}
	if a.Token != nil {
		account.Token = *a.Token
	}
	if a.RootFolder != nil {
		account.RootFolder = *a.RootFolder
	}
	if a.StatsCurrent != nil {
		account.StatsCurrent = *a.StatsCurrent
	}
	if a.Icon != nil {
		account.Icon = a.Icon
	}
	return account
}

func (a *UniversalAccount) GetId() GetId {
	detId := GetId{}
	if a.Id != nil {
		detId.Id = *a.Id
	}
	return detId
}

func (a *UniversalAccount) CreatedAccount() CreatedAccount {
	createdAccount := CreatedAccount{}
	if a.Id != nil {
		createdAccount.Id = *a.Id
	}
	if a.Token != nil {
		createdAccount.Token = *a.Token
	}
	return createdAccount
}
