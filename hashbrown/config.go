package hashbrown

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

func Setup(config string) {
	reader := bufio.NewReader(os.Stdin)

	host_name, err := os.Hostname()
	CheckErr(err)
	h := sha1.New()
	h.Write([]byte(host_name))
	bs := h.Sum(nil)
	temp_key := bs[4:]
	// fmt.Println(host_name)
	// fmt.Println(bs)
	// fmt.Println(temp_key)
	key := []byte(temp_key)
	fmt.Printf("file %s not exists creating..\n", config)
	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')
	secret := Encrypt(key, password)
	file, err := os.Create(config)
	if err != nil {
		fmt.Printf("error: %v in creating file %s\n", err, config)
	}
	defer file.Close()
	_, err = file.WriteString(secret)
	if err != nil {
		fmt.Printf("error: %v in writing config\n", err)
	}
	defer file.Close()
}

func Encrypt(key []byte, text string) string {
	plaintext := []byte(text)
	block, err := aes.NewCipher(key)
	CheckErr(err)

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext)
}

func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}
