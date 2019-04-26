//go:generate go build -ldflags="-s -w" -o go-launcher cmd/launcher/main.go
//go:generate upx --brute go-launcher
