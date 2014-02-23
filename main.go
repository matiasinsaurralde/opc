package main

// open proxy crawler

import "net/http"
import "fmt"

func pageHandler( w http.ResponseWriter, r *http.Request ){
  filePath := "./public" + r.RequestURI
  if( filePath == "./" ) {
    http.ServeFile( w, r, "./public/index.html" )
  } else {
    http.ServeFile( w, r, filePath )
  }
}

func main() {
  fmt.Println( "up and running!" )
  http.HandleFunc( "/", pageHandler )
  http.ListenAndServe( ":8000", nil )
}

