package admin

import (
	"context"
	"time"
)

// User is the firebase auth user
type User struct {
	LocalID           string      `json:"localId,omitempty"`
	Email             string      `json:"email,omitempty"`
	EmailVerified     bool        `json:"emailVerified,omitempty"`
	ProviderUserInfo  []*UserInfo `json:"providerUserInfo,omitempty"`
	PasswordHash      string      `json:"passwordHash,omitempty"`
	PasswordUpdatedAt float64     `json:"passwordUpdatedAt,omitempty"`
	ValidSince        string      `json:"validSince,omitempty"`
	Disabled          bool        `json:"disabled,omitempty"`
	LastLoginAt       string      `json:"lastLoginAt,omitempty"`
	CreatedAt         string      `json:"createdAt,omitempty"`
}

// UserInfo type
type UserInfo struct {
	ProviderID  string `json:"providerId,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	PhotoURL    string `json:"photoUrl,omitempty"`
	FederatedID string `json:"federatedId,omitempty"`
	Email       string `json:"email,omitempty"`
	RawID       string `json:"rawId,omitempty"`
	ScreenName  string `json:"screenName,omitempty"`
}

type apiError struct {
	Message string `json:"message"`
}

type getAccountInfoRequest struct {
	LocalIDs []string `json:"localId,omitempty"`
	Emails   []string `json:"email,omitempty"`
}

type getAccountInfoResponse struct {
	Users []*User `json:"users,omitempty"`
}

var scopes = []string{
	"https://www.googleapis.com/auth/userinfo.email",
	"https://www.googleapis.com/auth/firebase.database",
	"https://www.googleapis.com/auth/identitytoolkit",
}

const (
	baseURL = "https://www.googleapis.com/identitytoolkit/v3/relyingparty/"
	timeout = time.Second * 10000
)

type apiMethod string

const (
	getAccountInfo   apiMethod = "getAccountInfo"
	setAccountInfo   apiMethod = "setAccountInfo"
	deleteAccount    apiMethod = "deleteAccount"
	uploadAccount    apiMethod = "uploadAccount"
	downloadAccount  apiMethod = "downloadAccount"
	getOOBCode       apiMethod = "getOobConfirmationCode"
	getProjectConfig apiMethod = "getProjectConfig"
)

type httpMethod string

const (
	httpGet  httpMethod = "GET"
	httpPost httpMethod = "POST"
)

func getContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), timeout)
}
