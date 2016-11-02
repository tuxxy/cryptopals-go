package main

import (
    "strings"
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

func FreqAnalysis(s string) int {
    ETAOIN := "etaoinshrdlcumwfgypbvkjxqz"

    score := 0
    for _, char := range strings.ToLower(s) {
        freqIndex := strings.Index(ETAOIN, string(char))
        if freqIndex != -1 {
            score += 26 - freqIndex
        }
    }
    return score
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

func Basics_Chall3() (string, int) {
    ciphertext_bytes := Bytify("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

    high_val := 0
    var high_val_str string

    key_length := len(*ciphertext_bytes)
    for i := 0; i < 256; i++ {
        // Generate a key from single char
        key := make([]byte, key_length)
        for h := 0; h < key_length; h++ {
            key[h] = byte(rune(i))
        }

        // XOR ciphertext with key
        xor_text := make([]byte, key_length)
        XOR_Bytes(*ciphertext_bytes, key, xor_text)
        str_xor_text := string(xor_text[:])

        new_high_val := FreqAnalysis(str_xor_text)
        if new_high_val > high_val {
            high_val = new_high_val
            high_val_str = str_xor_text
        }
    }
    return high_val_str, high_val
}

func Basics_Chall4() (string, int) {
    chall_file := GetBody("http://cryptopals.com/static/challenge-data/4.txt")
    cipher_strings := strings.Split(string(chall_file), "\n")

    highest_val := 0
    var highest_val_str string

    for _, ciphertext := range cipher_strings {
        high_val := 0
        var high_val_str string

        key_length := len(ciphertext)
        for i := 0; i < 256; i++ {
            // Generate key from single char
            key := make([]byte, key_length)
            for h := 0; h < key_length; h++ {
                key[h] = byte(rune(i))
            }

            // XOR ciphertext with key
            xor_text := make([]byte, key_length)
            XOR_Bytes(*Bytify(ciphertext), key, xor_text)
            str_xor_text := string(xor_text[:])

            new_high_val := FreqAnalysis(str_xor_text)
            if new_high_val > high_val {
                high_val = new_high_val
                high_val_str = str_xor_text
            }
        }
        
        if high_val > highest_val {
            highest_val = high_val
            highest_val_str = high_val_str
        }
    }
    return highest_val_str, highest_val
}
