package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	res := make(map[string]bool)
	resChannel := make(chan result)

	for _, url := range urls {
		go func() {
			resChannel <- result{url, wc(url)}
		}()
	}

	for range urls {
		r := <-resChannel
		res[r.string] = r.bool
	}

	return res
}
