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

func KeyValue(Size int) string {
  //Get a new hex string (return value is 32 characters in length)
  token := HexString()
  //only return the requested size from the 32 character hex value
  //  the below statement says return substring from index 0 to Size-1
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
  //incomplete method
  //this method saves the value and the key into a file
  fOut, writeError := os.Create(TokenFile)
  defer fOut.Close()
  if writeError != nil {
    return
  }
}

func booltostr(val bool) string {
  if val == true {
    return "True"
  } else {
    return "False"
  }
}

func main() {
  //Creates a reader from the standard input (in this case from the os command line)
  reader := bufio.NewReader(os.Stdin)
  fmt.Println("Enter value to save")
  //reads a full string up to the carriage return from the standard input
  value, _ := reader.ReadString('\n')
  fmt.Printf("You entered %s",value)
  //create a key to use as the replacement value from the input
  key := KeyValue(9)
  fmt.Println(key+" ",(len(key)))
  //Determine if the "key" (token) has already been used
  fmt.Println("Does key exists? ",TokenExists(key))
}
