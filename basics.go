package main

import (
    "encoding/hex"
    "encoding/base64"
)

func Bytify(s string) *[]byte {
    byte_msg, err := hex.DecodeString(s)
    if err != nil {
        panic(err)
    }
    return &byte_msg
}

func Stringify(src []byte) string {
    return hex.EncodeToString(src)
}

func Bytes_b64(src *[]byte) string {
    str_msg := base64.StdEncoding.EncodeToString(*src)
    return str_msg
}

func Basics_Chall1() string {
    hex_string := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
    b64_msg := Bytes_b64(Bytify(hex_string))
    return b64_msg
}

func Basics_Chall2() {

}
