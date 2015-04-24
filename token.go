package Tokenizer

// Internal use functions
//
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
