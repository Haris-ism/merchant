package utils

import (
	"crypto/rand"
)

func GenerateRandom(length int)(string,error){
	char:="1234567890abcdefghijklmnopqrstuvwxyz"
	buffer := make([]byte, length)
    _, err := rand.Read(buffer)
    if err != nil {
		return "",err
    }
	for i := 0; i < length; i++ {
        buffer[i] = char[int(buffer[i])%length]
    }
	// result,err:=strconv.Atoi(string(buffer))
	return string(buffer),nil
}