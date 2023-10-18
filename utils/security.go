package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"merchant/controllers/models"

	"io"
)

func serializeStruct(input interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(input)
	if err != nil {
		return nil, err
	}
	
	return buffer.Bytes(), nil
}

func DecryptTransItemRes(req models.ResTransItem)(models.DecTransItem,error){
	res:=models.DecTransItem{}
	chanID:=make(chan string)
	chanName:=make(chan string)
	chanQuantity:=make(chan string)
	chanCC:=make(chan string)
	chanCode:=make(chan []string)

	go DecryptFunc(req.ID,chanID)
	go DecryptFunc(req.Name,chanName)
	go DecryptFunc(req.Quantity,chanQuantity)
	go DecryptFunc(req.CC,chanCC)
	go DecryptArray(req.Code,chanCode)

	res.ID=<-chanID
	res.Name=<-chanName
	res.Quantity=<-chanQuantity
	res.CC=<-chanCC
	res.Code=<-chanCode

	return res,nil
}

func EncryptTransItemRes(req models.ResTransItem,codes []string)(models.ResTransItem,error){
	chanID:=make(chan string)
	chanName:=make(chan string)
	chanQuantity:=make(chan string)
	chanCC:=make(chan string)
	chanCode:=make(chan string)

	go EncryptFunc(req.ID,chanID)
	go EncryptFunc(req.Name,chanName)
	go EncryptFunc(req.Quantity,chanQuantity)
	go EncryptFunc(req.CC,chanCC)
	go EncryptArray(codes,chanCode)

	req.ID=<-chanID
	req.Name=<-chanName
	req.Quantity=<-chanQuantity
	req.CC=<-chanCC
	req.Code=<-chanCode

	return req,nil
}

func EncryptArray(arr []string,ch chan string){
	serialized, err := serializeStruct(arr)
	if err != nil {
		fmt.Println("err:",err)
	}
	res,err:=EncryptionAES(string(serialized))
	if err!=nil{
		fmt.Println("Encrypt Fails")
	}
	ch<-res
}
func DecryptArray(arr string,ch chan []string){
	res,err:=DecryptionAES(arr)
	if err!=nil{
		fmt.Println("Encrypt Fails")
	}
	var result []string
	buffer := bytes.NewBuffer([]byte(res))
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(&result)
	if err != nil {
		fmt.Println("fail decr",err)
	}

	ch<-result
}

func EncryptTransItem(req models.ReqTransItem)(models.ReqTransItem,error){
	chanID:=make(chan string)
	chanDiscount:=make(chan string)
	chanQuantity:=make(chan string)
	chanCCNumber:=make(chan string)
	// chanCVV:=make(chan string)
	chanAmount:=make(chan string)
	// chanPrice:=make(chan string)
	// chanName:=make(chan string)
	// chanType:=make(chan string)
	// chanPercentage:=make(chan string)

	// itemId:=strconv.Itoa(req.ItemID)
	// qty:=strconv.Itoa(req.Quantity)
	// amount:=strconv.Itoa(req.Amount)
	// price:=strconv.Itoa(req.Price)
	// percent:=strconv.Itoa(req.Percentage)
	go EncryptFunc(req.ID,chanID)
	go EncryptFunc(req.Discount,chanDiscount)
	go EncryptFunc(req.Quantity,chanQuantity)
	go EncryptFunc(req.CC,chanCCNumber)
	// go EncryptFunc(req.CVV,chanCVV)
	go EncryptFunc(req.Amount,chanAmount)
	// go EncryptFunc(req.Price,chanPrice)
	// go EncryptFunc(req.Name,chanName)
	// go EncryptFunc(req.Type,chanType)
	// go EncryptFunc(req.Percentage,chanPercentage)

	req.ID=<-chanID
	req.Discount=<-chanDiscount
	req.Quantity=<-chanQuantity
	req.CC=<-chanCCNumber
	// req.CVV=<-chanCVV
	req.Amount=<-chanAmount
	// req.Price=<-chanPrice
	// req.Name=<-chanName
	// req.Type=<-chanType
	// req.Percentage=<-chanPercentage

	return req,nil
}
func DecryptTransItem(req models.ReqTransItem)(models.ReqTransItem,error){
	chanID:=make(chan string)
	chanDiscount:=make(chan string)
	chanQuantity:=make(chan string)
	chanCCNumber:=make(chan string)
	// chanCVV:=make(chan string)
	chanAmount:=make(chan string)
	// chanPrice:=make(chan string)
	// chanName:=make(chan string)
	// chanType:=make(chan string)
	// chanPercentage:=make(chan string)

	// itemId:=strconv.Itoa(req.ItemID)
	// qty:=strconv.Itoa(req.Quantity)
	// amount:=strconv.Itoa(req.Amount)
	// price:=strconv.Itoa(req.Price)
	// percent:=strconv.Itoa(req.Percentage)
	go DecryptFunc(req.ID,chanID)
	go DecryptFunc(req.Discount,chanDiscount)
	go DecryptFunc(req.Quantity,chanQuantity)
	go DecryptFunc(req.CC,chanCCNumber)
	// go EncryptFunc(req.CVV,chanCVV)
	go DecryptFunc(req.Amount,chanAmount)
	// go EncryptFunc(req.Price,chanPrice)
	// go EncryptFunc(req.Name,chanName)
	// go EncryptFunc(req.Type,chanType)
	// go EncryptFunc(req.Percentage,chanPercentage)

	req.ID=<-chanID
	req.Discount=<-chanDiscount
	req.Quantity=<-chanQuantity
	req.CC=<-chanCCNumber
	// req.CVV=<-chanCVV
	req.Amount=<-chanAmount
	// req.Price=<-chanPrice
	// req.Name=<-chanName
	// req.Type=<-chanType
	// req.Percentage=<-chanPercentage

	return req,nil
}


func EncryptFunc(input string, ch chan string){
	res,err:=EncryptionAES(input)
	if err!=nil{
		fmt.Println("Encrypt Fails")
	}
	ch<-res
}
func DecryptFunc(input string, ch chan string){
	res,err:=DecryptionAES(input)
	if err!=nil{
		fmt.Println("Decrypt Fails")
	}
	ch<-res
}

func EncryptionAES(input string) (string, error) {
	result:=""
	key, _ := hex.DecodeString(GetEnv("KEY"))

	plaintext:=[]byte(input)

	block, err := aes.NewCipher(key)
	if err != nil {
		return result, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return result, err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func DecryptionAES(input string) (string, error) {
	result:=""
	key, _ := hex.DecodeString(GetEnv("KEY"))

	ciphertext, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return result, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return result, err
	}

	if len(ciphertext) < aes.BlockSize {
		return result, fmt.Errorf("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(ciphertext, ciphertext)

	result=string(ciphertext)

	return result, nil
}