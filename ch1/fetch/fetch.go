package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
			os.Exit(1)
		}
		b, err1 := ioutil.ReadAll(resp.Body)
		c := os.Stdout
		bytes, err2 := io.Copy(c, resp.Body)

		resp.Body.Close()
		if err1 != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s:%v\n", url, err1)
			os.Exit(1)
		}
		if err2 != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s:%v\n", url, err2)
			os.Exit(1)
		}
		fmt.Printf("ccccc%d", bytes)
		fmt.Printf("%s", b)
	}
}
