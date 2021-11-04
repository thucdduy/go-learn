package main

type app struct {
}

func (app *app) handleRequest(method string, url string) (int, string) {
	if url == "/app/status" && method == "GET" {
		return 200, "OK"
	}

	if url == "/create/user" && method == "POST" {
		return 201, "User created"
	}

	return 404, "Not Found"
}
