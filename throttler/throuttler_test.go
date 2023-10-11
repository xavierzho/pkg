package throttler

import (
	"net/http"
	"sync"
	"testing"
)

func TestHandler_RoundTrip(t *testing.T) {
	handler := NewThrottler(6, 1, WithErrContent("exceeded"), WithRoundTripper(http.DefaultTransport))
	client := http.Client{
		Transport: handler,
	}
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := client.Get("https://api.bscscan.com/api?module=logs&action=getLogs&apikey=J5YQDEZWE4RDNF66F4SDX6V7RATP7SWS4V&address=0x888e52605383a0f5fc61e3f23363b28a4d59a6b4&fromBlock=32443600&toBlock=latest&topic0=0xb3d987963d01b2f68493b4bdb130988f157ea43070d4ad840fee0466ed9370d9&page=1&offset=10000")
			if err != nil {
				//fmt.Println(err)
				t.Error(err)
			}
		}()
	}
	wg.Wait()

}
