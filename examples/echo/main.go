package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	url   = "https://localhost:8080"
	topic = "test_topic"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		sc.Scan()
		if err := sc.Err(); err != nil {
			log.Fatalf("an error occurred: %v", err)
		}

		input := sc.Text()

		if input == "q" {
			log.Println("bye!")
			os.Exit(0)
		}

		res, err := http.Post(
			fmt.Sprintf("%s/publish/%s", url, topic),
			"application/json",
			strings.NewReader(input),
		)
		if err != nil {
			log.Printf("failed to publish: %v", err)
			continue
		}
		if res.StatusCode != http.StatusCreated {
			log.Printf("failed to publish, received status code: %d", res.StatusCode)
		}

		fmt.Printf("Published message %s to topic %s\n", input, topic)

		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		enc.Encode("INIT")
		enc.Encode("ACK")

		res, err = http.Post(fmt.Sprintf("%s/subscribe/%s", url, topic), "application/json", &buf)
		if err != nil {
			log.Printf("failed to consume: %v", err)
			continue
		}
		if res.StatusCode != http.StatusOK {
			log.Printf("failed to consume, received status code: %d", res.StatusCode)
			continue
		}

		var subRes struct {
			Msg   string `json:"msg"`
			Error string `json:"error"`
		}
		if err := json.NewDecoder(res.Body).Decode(&subRes); err != nil {
			log.Printf("failed decode response body: %v", err)
		}

		res.Body.Close()

		fmt.Printf("Consumed message: %s\n", mustBase64Decode(subRes.Msg))
	}
}

func mustBase64Decode(b string) string {
	s, err := base64.StdEncoding.DecodeString(string(b))
	if err != nil {
		log.Fatal(err)
	}

	return string(s)
}
