package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	api "sencha-twirp-rpc/rpc/themes"
)

func main() {
	client := api.NewThemesProtobufClient("http://localhost:8080", &http.Client{})

	themeResp, err := client.GetTheme(context.Background(), &api.GetThemeRequest{Name: "pink"})

	if err != nil {
		fmt.Println("Client main.go themeResp Error:", err)
		os.Exit(1)
	}

	fmt.Println("Theme Colors: ", themeResp.ThemeData)
}
