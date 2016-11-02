package basics

import (
    "encoding/hex"
    "encoding/base64"
)

func Bytify(s string) []byte {
    byte_msg, err := hex.DecodeString(s)
    if err != nil {
        panic(err)
    }
    return byte_msg
}

func bytes_b64(src []byte) string {
    enc := base64.NewEncoding(base64.StdEncoding)
    str_msg := enc.EncodeToString(&src)
    return str_msg
}
