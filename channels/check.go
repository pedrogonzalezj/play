package channels

import "sync"

// WebsiteChecker checks a url, returning a bool.
type WebsiteChecker func(string) bool
type result struct {
	url    string
	status bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls {
		go func(u string) {
			status := wc(u)
			resultChannel <- result{u, status}
			wg.Done()
		}(url)
	}
	// Launch a goroutine to monitor when all the work is done.
	go func() {
		// Wait to everything to be processed.
		wg.Wait()
		// close the results channel.
		close(resultChannel)
	}()
	for r := range resultChannel {
		results[r.url] = r.status
	}
	return results
}
