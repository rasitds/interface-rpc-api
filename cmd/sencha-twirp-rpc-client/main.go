package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	api "sencha-twirp-rpc/rpc/colors"
)

func main() {
	client := api.NewColorsProtobufClient("http://localhost:8080", &http.Client{})

	colorResp, err := client.GetColors(context.Background(), &api.GetColorsRequest{Name: "red"})

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	fmt.Println("Color: ", colorResp.ColorData)
}
