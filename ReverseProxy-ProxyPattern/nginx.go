package main

type nginx struct {
	rateLimiter     map[string]int
	maxAllowRequest int
	backend         *app
}

func newNginxServer() *nginx {
	return &nginx{
		rateLimiter:     make(map[string]int),
		maxAllowRequest: 2,
		backend:         &app{},
	}

}

func (n *nginx) handleRequest(method, url string) (int, string) {

	if ok := n.checkRateLimiting(url); !ok {
		return 403, "Not Allowed"
	}

	return n.backend.handleRequest(method, url)
}

func (n *nginx) checkRateLimiting(url string) bool {

	n.rateLimiter[url]++

	return n.rateLimiter[url] <= n.maxAllowRequest

}
