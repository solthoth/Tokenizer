package main

import (
  "fmt"
  "log"
  "net/http"
  "encoding/json"
  "crypto/rand"
  "encoding/hex"
  "github.com/gorilla/mux"
  )

  // Token definition
type Token struct {
  SSN string
  Value string
}

func HexString() string {
  // The following function was copied from stack-overflow
  // http://stackoverflow.com/questions/15130321/is-there-a-method-to-generate-a-uuid-with-go-language
  u := make([]byte,16)
  _, err := rand.Read(u)
  if err != nil {
    return ""
  }

  u[8] = (u[8] | 0x80) & 0xBF
  u[6] = (u[6] | 0x40) & 0x4F
  return hex.EncodeToString(u)
}

func NewToken(length int) string {
  if length != 0 {
    return HexString()[:length]
  } else {
    return HexString()
  }
}
// ---------
// Routers
//----------
func HomePage(writer http.ResponseWriter, request *http.Request) {
  fmt.Fprintf(writer,"Welcome home")
}

func SSN2Token(writer http.ResponseWriter, request *http.Request) {
  vars := mux.Vars(request)
  value := vars["ssn"]
  myToken := Token{
    SSN: value,
    Value: NewToken(9),
  }
  json.NewEncoder(writer).Encode(myToken)
}

func Token2SSN(writer http.ResponseWriter, request *http.Request) {
  vars := mux.Vars(request)
  ssn := vars["token"]
  fmt.Fprintf(writer,"You sent %s",ssn)
}

func Token2Last4SSN(writer http.ResponseWriter, request *http.Request) {
  vars := mux.Vars(request)
  ssn := vars["token"]
  fmt.Fprintf(writer,"You sent %s",ssn)
}

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/", HomePage)
  router.HandleFunc("/SSN2Token/{ssn}",SSN2Token)
  router.HandleFunc("/Token2SSN/{token}",Token2SSN)
  router.HandleFunc("/Token2Last4SSN/{token}",Token2Last4SSN)
  http.Handle("/", router)

  log.Fatal(http.ListenAndServe(":3000",router))
}
