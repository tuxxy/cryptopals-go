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

func XOR_Bytes(src_a, src_b, dest []byte) {
    for i := 0; i < len(src_a); i++ {
        dest[i] = src_a[i] ^ src_b[i]
    }
}

func Basics_Chall1() string {
    hex_string := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
    b64_msg := Bytes_b64(Bytify(hex_string))
    return b64_msg
}

func Basics_Chall2() string {
    byte_msg_a := Bytify("1c0111001f010100061a024b53535009181c")
    byte_msg_b := Bytify("686974207468652062756c6c277320657965")

    byte_c := make([]byte, len(*byte_msg_a))
    XOR_Bytes(*byte_msg_a, *byte_msg_b, byte_c)

    return Stringify(byte_c)
}

func Basics_Chall3() string {
    
}
