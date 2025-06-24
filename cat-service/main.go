package main

import (
	"fmt"

	server "github.com/sai7xp/go-microservice-template/cat-service/cmd"
	commonutils "github.com/sai7xp/go-microservice-template/common-utils/crypto/aes"
)

func main() {
	key, _ := commonutils.GenerateSymmetricKey()
	encryptedString := commonutils.EncryptData(key, []byte("hello"))
	fmt.Println(encryptedString)
	decryptedData := commonutils.DecryptData("swOwWa9MfseTK7q2FkbRjseWCntJGoUypcI+8w91Ico=", "5b/QzGVZY/4xtsVwTU9HQu+XVhzA6jXm+QDAUvOLpMSC")
	fmt.Println(string(decryptedData))

	server.Serve()
}
