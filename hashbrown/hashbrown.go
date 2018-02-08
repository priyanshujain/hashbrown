package hashbrown

import (
	"fmt"
	"io"
	"os"
)

func Generate(config, service, salt string, length int) {
	var file, err = os.OpenFile(config, os.O_RDWR, 0644)
	CheckErr(err)
	defer file.Close()

	// read file, line by line
	var text = make([]byte, 1024)
	for {
		_, err = file.Read(text)

		// break if finally arrived at end of file
		if err == io.EOF {
			break
		}

		// break if error occured
		if err != nil && err != io.EOF {
			CheckErr(err)
			break
		}
	}
	// h1 := sha256.New()
	// h1.Write([]byte(text))
	// temp_key := h1.Sum(nil)
	temp_key := text[:16]
	key := []byte(temp_key)

	plain_text := service + salt
	secret := Encrypt(key, plain_text)

	generated_password := string(secret[:length])

	fmt.Println("\n***************************")
	fmt.Printf("\nGenerated password for %s is\n", service)
	fmt.Println(generated_password)
}
