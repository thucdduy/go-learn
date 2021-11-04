package main

import "fmt"

func main() {
	nginxServer := newNginxServer()

	appStatusURL := "/app/status"
	createuserURL := "/create/user"

	httpCode, body := nginxServer.handleRequest("GET", appStatusURL)
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)
	httpCode, body = nginxServer.handleRequest("GET", appStatusURL)
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)
	httpCode, body = nginxServer.handleRequest("GET", appStatusURL)
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)
	httpCode, body = nginxServer.handleRequest("POST", createuserURL)
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", createuserURL, httpCode, body)
	httpCode, body = nginxServer.handleRequest("GET", createuserURL)
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", createuserURL, httpCode, body)
	httpCode, body = nginxServer.handleRequest("GET", createuserURL)
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", createuserURL, httpCode, body)

}
