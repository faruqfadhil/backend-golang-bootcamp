package main

import (
	"crypto/aes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"regexp"
)

func main() {
	// example1()
	// example2()
	// example3()
	// example4()
	// cryptoExample1()
	mypassword := "secrett"
	fmt.Println("MD5 Hashed value: ", getMD5Hash(mypassword))
}

func example1() {
	var text = "banana burger soup"
	var regex, err = regexp.Compile(`[a-z]+`)

	if err != nil {
		fmt.Println(err.Error())
	}

	var res1 = regex.FindAllString(text, 2)
	fmt.Printf("%#v \n", res1)
	// []string{"banana", "burger"}

	var res2 = regex.FindAllString(text, -1)
	fmt.Printf("%#v \n", res2)
	// []string{"banana", "burger", "soup"}
}

func example2() {
	var text = "banana burger soup"
	var regex, _ = regexp.Compile(`[a-z]+`)

	var isMatch = regex.MatchString(text)
	fmt.Println(isMatch)
}

func cryptoExample1() {
	plainText := "This is a secret"
	key := "this_must_be_of_32_byte_length!!"

	emsg := encryptMessage(key, plainText)
	dmesg := decryptMessage(key, emsg)

	fmt.Println("Encrypted Message: ", emsg)
	fmt.Println("Decrypted Message: ", dmesg)
}

func encryptMessage(key string, message string) string {
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
	}
	msgByte := make([]byte, len(message))
	c.Encrypt(msgByte, []byte(message))
	return hex.EncodeToString(msgByte)
}

func decryptMessage(key string, message string) string {
	txt, _ := hex.DecodeString(message)
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
	}
	msgByte := make([]byte, len(txt))
	c.Decrypt(msgByte, []byte(txt))

	msg := string(msgByte[:])
	return msg
}

func example3() {
	var text = "banana burger soup"
	var regex, _ = regexp.Compile(`[a-z]+`)

	var str = regex.ReplaceAllStringFunc(text, func(each string) string {
		if each == "burger" {
			return "potato"
		}
		return each
	})
	fmt.Println(str)
	// "banana potato soup"
}

func example4() {
	var text = "banana burger soup"
	var regex, _ = regexp.Compile(`[a-b]+`) // split dengan separator adalah karakter "a" dan/atau "b"

	var str = regex.Split(text, -1)
	fmt.Printf("%#v \n", str)
	// []string{"", "n", "n", " ", "urger soup"}
}

func getMD5Hash(message string) string {
	hash := md5.Sum([]byte(message))
	return hex.EncodeToString(hash[:])
}
