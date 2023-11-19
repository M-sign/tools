package main

var (
	signToFile     bool
	privateFile    string
	publicFile     string
	forceOverwrite bool
)

const (
	msign_Env_Private = "MSIGN_PRIVATE"
	msign_Env_Public  = "MSIGN_PUBLIC"
)
