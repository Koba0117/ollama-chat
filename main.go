package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

type RequestData struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type ResponseData struct {
	Response string `json:"response"`
}

func main() {
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, nil)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		prompt := r.PostFormValue("form")
		data := &RequestData{Model: "mistral", Prompt: prompt, Stream: false}

		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "chat", ollama(*data))
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/submit-prompt/", h2)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ollama(data RequestData) string {
	jsonValue, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return ""
	}

	req, err := http.NewRequest("POST", "http://localhost:11434/api/generate", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return ""
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return ""
	}

	var responseData ResponseData
	err = json.Unmarshal(body, &responseData) // Decode response data into the struct
	if err != nil {
		fmt.Println("JSON parsing failed:", err)
		return ""
	}

	return responseData.Response
}

// func MarkdownToHTML(markdown string) (string, error) {
// 	processor := commonmark.NewProcessor()
// 	node, err := processor.Process(markdown, nil)
// 	if err != nil {
// 		return "", err
// 	}
// 	var buf bytes.Buffer
// 	err = html.Render(&buf, node, nil)
// 	if err != nil {
// 		return "", err
// 	}
// 	return buf.String(), nil
// }
