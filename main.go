package main

// open proxy crawler

import( "net/http"
	"math/rand"
	"strconv"
	"net/url"
	"time"
	"flag"
	"fmt" )

var mode int

func pageHandler( w http.ResponseWriter, r *http.Request ){
  filePath := "./public" + r.RequestURI
  if( filePath == "./" ) {
    http.ServeFile( w, r, "./public/index.html" )
  } else {
    http.ServeFile( w, r, filePath )
  }
}

func crawl() {
  //fmt.Println( "i'm crawling the world" )
  semaphore := make( chan bool, 5 )
  for {
    semaphore <- true
    someProxy := randomIp()+":3128"
    go testProxy( someProxy )
  }
}

func testProxy(proxyAddress string) {
  proxyUrl, _ := url.Parse("http://"+proxyAddress)
  testClient := &http.Client{ Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)} }
  response, err := testClient.Get("http://www.google.com/robots.txt")
  fmt.Println( err )
  fmt.Println( response )
}

func randomIp() string {
  rand.Seed( time.Now().UnixNano() )
  var someIp string = ""
  for i := 0; i < 4; i++ {
    n := rand.Intn(254)
    someIp += strconv.Itoa( n )
    if i < 3 {
      someIp += "."
    }
  }
  return someIp
}

func main() {
  flag.Int( "mode", 0, "Crazy crawl mode (default)" )

  flag.Parse()

  //fmt.Println( "up and running!" )
  go crawl()  
  http.HandleFunc( "/", pageHandler )
  http.ListenAndServe( ":8000", nil )
}
