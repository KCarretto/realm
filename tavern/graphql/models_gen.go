// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql

import (
	"github.com/kcarretto/realm/tavern/ent"
	"github.com/kcarretto/realm/tavern/ent/credential"
)

type CallbackInput struct {
	TargetID    int    `json:"targetID"`
	SessionID   string `json:"sessionID"`
	ConfigName  string `json:"configName"`
	ProcessName string `json:"processName"`
}

type CallbackResponse struct {
	Implant *ent.Implant `json:"implant"`
}

type CreateCredentialInput struct {
	Principal string          `json:"principal"`
	Secret    string          `json:"secret"`
	Kind      credential.Kind `json:"kind"`
	TargetID  int             `json:"targetID"`
}

type CreateDeploymentConfigInput struct {
	Name            string  `json:"name"`
	Cmd             *string `json:"cmd"`
	StartCmd        *bool   `json:"startCmd"`
	FileDst         *string `json:"fileDst"`
	ImplantConfigID *int    `json:"implantConfigID"`
	FileID          *int    `json:"fileID"`
}

type CreateImplantCallbackConfigInput struct {
	URI      string  `json:"uri"`
	ProxyURI *string `json:"proxyURI"`
	Priority *int    `json:"priority"`
	Timeout  *int    `json:"timeout"`
	Interval *int    `json:"interval"`
	Jitter   *int    `json:"jitter"`
}

type CreateImplantConfigInput struct {
	Name              string `json:"name"`
	ServiceConfigIDs  []int  `json:"serviceConfigIDs"`
	CallbackConfigIDs []int  `json:"callbackConfigIDs"`
}

type CreateImplantServiceConfigInput struct {
	Name           string  `json:"name"`
	Description    *string `json:"description"`
	ExecutablePath string  `json:"executablePath"`
}

type CreateTagInput struct {
	Name      string `json:"name"`
	TargetIDs []int  `json:"targetIDs"`
}

type CreateTargetInput struct {
	Name             string `json:"name"`
	ForwardConnectIP string `json:"forwardConnectIP"`
}

type UpdateDeploymentConfigInput struct {
	ID              int     `json:"id"`
	Name            *string `json:"name"`
	Cmd             *string `json:"cmd"`
	StartCmd        *bool   `json:"startCmd"`
	FileDst         *string `json:"fileDst"`
	ImplantConfigID *int    `json:"implantConfigID"`
	FileID          *int    `json:"fileID"`
}

type UpdateImplantCallbackConfigInput struct {
	ID       int     `json:"id"`
	URI      *string `json:"uri"`
	ProxyURI *string `json:"proxyURI"`
	Priority *int    `json:"priority"`
	Timeout  *int    `json:"timeout"`
	Interval *int    `json:"interval"`
	Jitter   *int    `json:"jitter"`
}

type UpdateImplantConfigInput struct {
	ID                      int     `json:"id"`
	Name                    *string `json:"name"`
	AddServiceConfigIDs     []int   `json:"addServiceConfigIDs"`
	RemoveServiceConfigIDs  []int   `json:"removeServiceConfigIDs"`
	AddCallbackConfigIDs    []int   `json:"addCallbackConfigIDs"`
	RemoveCallbackConfigIDs []int   `json:"removeCallbackConfigIDs"`
}

type UpdateImplantServiceConfigInput struct {
	ID             int     `json:"id"`
	Name           *string `json:"name"`
	Description    *string `json:"description"`
	ExecutablePath *string `json:"executablePath"`
}

type UpdateTagInput struct {
	ID              int     `json:"id"`
	Name            *string `json:"name"`
	AddTargetIDs    []int   `json:"addTargetIDs"`
	RemoveTargetIDs []int   `json:"removeTargetIDs"`
}
