package concurrency

type WebsiteChecker func(string) bool

type results struct {
	string
	bool
}

func checkWebsiteStatus(wc WebsiteChecker, urls []string) map[string]bool {
	result := make(map[string]bool)
	resChannel := make(chan results)

	for _, url := range urls {
		go func(u string) {
			resChannel <- results{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resChannel
		result[r.string] = r.bool
	}
	return result
}
