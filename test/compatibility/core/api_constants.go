package core

type RuntimeAPIName string

const (
	SetContextAPIName RuntimeAPIName = "SetContext"
	GetContextAPIName                = "GetContext"
	AddServerAPIName                 = "AddServer"
	GetServerAPIName                 = "GetServer"
)
