/**
* Fetch prints the content found at a URL.
* This program introduces functions from two packages, net/http and io/ioutil. The
* http.Get function makes an HTTP request and, if there is no error, returns the result in the
* response struct resp. The Body field of resp contains the server response as a readable
* stream. Next, ioutil.ReadAll reads the entire response; the result is stored in b
 */

package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// body, err := ioutil.ReadAll(resp.Body) // Reads the whole file into memory

		var buf bytes.Buffer
		io.Copy(&buf, resp.Body)

		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("status: %s\n%s", resp.Status, buf.String())
	}
}
