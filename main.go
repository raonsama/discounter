package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

const apiVersion = "v10"

func main() {
	displayASCIIArt()

	var channelId string
	var number uint
	var delay uint

	token := os.Getenv("TOKEN")
	if token == "" {
		fmt.Println("Discord Token not found.")
		os.Exit(1)
	}

	flag.StringVar(&channelId, "c", "", "Specify channel id.")
	flag.UintVar(&delay, "d", 5, "Specify number for delay.")
	flag.UintVar(&number, "n", 0, "Specify number for counter.")
	flag.Parse()

	if channelId == "" {
		fmt.Println("Please provide channel id.")
		os.Exit(1)
	}

	apiEndpoint := fmt.Sprintf("https://discord.com/api/%s/channels/%s/messages", apiVersion, channelId)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println("Bye.")
		os.Exit(0)
	}()

	for {
		number++
		payload := []byte(fmt.Sprintf(`{"content": "%d"}`, number))

		req, err := http.NewRequest("POST", apiEndpoint, bytes.NewBuffer(payload))
		req.Header.Add("Authorization", token)
		req.Header.Set("Content-Type", "application/json")

		if err != nil {
			fmt.Println("Error creating request:", err)
			os.Exit(1)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			os.Exit(1)
		}

		fmt.Printf("Response Status: %d\t| Counter: %d\n", resp.StatusCode, number)

		resp.Body.Close()

		time.Sleep(time.Second * time.Duration(delay))
	}
}

func displayASCIIArt() {
	asciiArt := `
	██████╗  █████╗  ██████╗ ███╗   ██╗
	██╔══██╗██╔══██╗██╔═══██╗████╗  ██║
	██████╔╝███████║██║   ██║██╔██╗ ██║
	██╔══██╗██╔══██║██║   ██║██║╚██╗██║
	██║  ██║██║  ██║╚██████╔╝██║ ╚████║
	╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═══╝
	Raymond M      @raonsama/discounter
	`
	fmt.Println(asciiArt)
}
