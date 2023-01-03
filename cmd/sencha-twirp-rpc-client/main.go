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

	fmt.Println("ADD THEMES")
	_, err := client.AddThemes(context.Background(), &api.AddThemesRequest{})
	if err != nil {
		fmt.Println("AddThemes Error:", err)
		os.Exit(1)
	}

	fmt.Println("----------------\nGET THEMES")
	themes, err := client.GetThemes(context.Background(), &api.GetThemesRequest{})
	if err != nil {
		fmt.Println("GetThemes Error:", err)
	}
	for i, theme := range themes.ThemeData {
		fmt.Println("theme", i, theme)
	}

	fmt.Println("----------------\nCREATE A THEME")
	createdTheme, err := client.CreateTheme(context.Background(), &api.CreateThemeRequest{Name: "green", Background: "#0C0C0C", Foreground: "#00FF00"})

	if err != nil {
		fmt.Println("CreateTheme Error:", err)
		os.Exit(1)
	}
	fmt.Println("Created Theme:", createdTheme.ThemeData)

	fmt.Println("----------------\nGET A THEME")
	themeResp, err := client.GetTheme(context.Background(), &api.GetThemeRequest{Name: "green"})

	if err != nil {
		fmt.Println("GetTheme Error:", err)
		os.Exit(1)
	}
	fmt.Println("Theme:", themeResp.ThemeData)
}
