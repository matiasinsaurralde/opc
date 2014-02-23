package main

// open proxy crawler

import "net/http"
import "net/url"
import "fmt"

func pageHandler( w http.ResponseWriter, r *http.Request ){
  filePath := "./public" + r.RequestURI
  if( filePath == "./" ) {
    http.ServeFile( w, r, "./public/index.html" )
  } else {
    http.ServeFile( w, r, filePath )
  }
}

func crawl() {
  fmt.Println( "i'm crawling the world" )
  someProxy := "201.217.55.97:3128"
  testProxy( someProxy )
}

func testProxy(proxyAddress string) {
  proxyUrl, _ := url.Parse("http://"+proxyAddress)
  testClient := &http.Client{ Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)} }
  response, err := testClient.Get("http://www.google.com/robots.txt")
  fmt.Println( err )
  fmt.Println( response )
}

func main() {
  fmt.Println( "up and running!" )
  go crawl()
  http.HandleFunc( "/", pageHandler )
  http.ListenAndServe( ":8000", nil )
}

