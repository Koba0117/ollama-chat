// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"net/url"
// )

// type Chat struct {
// 	Model  string
// 	Prompt string
// }

// func main() {
// 	ollma_url := "http://localhost:11434/api/generate"
// 	data := url.Values{
// 		"model":  {"llama2"},
// 		"prompt": {"what is golang"},
// 		"stream": {"false"},
// 	}

// 	resp, err := http.PostForm(ollma_url, data)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer resp.Body.Close()

// 	var res map[string]interface{}

// 	json.NewDecoder(resp.Body).Decode(&res)

// 	fmt.Println(res["form"])
// }

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Data struct {
	Key1 string `json:"model"`
	Key2 string `json:"prompt"`
	Key3 bool   `json:"stream"`
}

func main() {
	data := &Data{Key1: "mistral", Key2: "print hello world in go", Key3: false}

	jsonValue, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", "http://localhost:11434/api/generate", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var responseData map[string]interface{}
	json.Unmarshal([]byte(body), &responseData) // assuming the API returns a JSON response

	fmt.Println("Response:", responseData)
}
