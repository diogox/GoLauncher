package main

type Result interface {
	Title() string
	Description() string
	IconPath() string
}
