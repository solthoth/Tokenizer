package main

import (
  "fmt"
  "bufio"
  "os"
  "crypto/rand"
  "encoding/hex"
  )

var
  TokenFile = "Tokens.TXT"

func HexString() string {
  u := make([]byte,16)
  _, err := rand.Read(u)
  if err != nil {
    return ""
  }

  u[8] = (u[8] | 0x80) & 0xBF
  u[6] = (u[6] | 0x40) & 0x4F
  return hex.EncodeToString(u)
}

func KeyValue(Size int) string {
  token := HexString()
  return token[:Size]
}

func TokenExists(key string) bool {
  fRead, err := os.Open(TokenFile)
  defer fRead.Close()
  if err != nil {
    return false
  }
  fReader := bufio.NewReader(fRead)
  for {
    sValue, readErr := fReader.ReadString('\n')
    if sValue == key {
      return true
    }
    if readErr != nil {
      break//all done
    }
  }
  return false
}

func SaveToken(value, key string) {
  fOut, writeError := os.Create(TokenFile)
  defer fOut.Close()
}

func booltostr(val bool) string {
  if val == true {
    return "True"
  } else {
    return "False"
  }
}

func main() {
  reader := bufio.NewReader(os.Stdin)
  fmt.Println("Enter value to save")
  value, _ := reader.ReadString('\n')
  fmt.Printf("You entered %s",value)
  key := KeyValue(9)
  fmt.Println(key)
  fmt.Println("Does key exists? ",TokenExists(key))
}
