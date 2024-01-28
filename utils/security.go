package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"merchant/constants"
	"merchant/controllers/models"
	"merchant/protogen/merchant"

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
	// go DecryptArray(req.Code,chanCode)

	res.ID=<-chanID
	res.Name=<-chanName
	res.Quantity=<-chanQuantity
	res.CC=<-chanCC
	res.Code=<-chanCode

	return res,nil
}

func EncryptTransItemResGrpc(req models.DecTransItem)(string,error){
	res:=""
	bytes,err:=json.Marshal(req)
	if err!=nil{
		return res,errors.New(constants.ERROR_DB)
	}
	chanReq:=make(chan string)

	go EncryptFunc(string(bytes),chanReq)

	res=<-chanReq

	return res,nil
}

func EncryptTransItemRes(req models.DecTransItem)(models.DecReqTransItem,error){
	res:=models.DecReqTransItem{}
	// fmt.Println("req:",req)
	bytes,err:=json.Marshal(req)
	if err!=nil{
		return res,errors.New(constants.ERROR_DB)
	}
	chanReq:=make(chan string)

	go EncryptFunc(string(bytes),chanReq)

	encrypted:=<-chanReq
	// fmt.Println("encrypted",encrypted)
	res.Req=encrypted

	return res,nil
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
func DecryptTransItemGrpc(req *merchant.ReqTransItemsModel)(models.ReqTransItem,error){
	res:=models.ReqTransItem{}
	chanReq:=make(chan string)
	go DecryptFunc(req.Request,chanReq)

	decrypted:=<-chanReq
	// fmt.Println("decrypted:",decrypted)
	err:=json.Unmarshal([]byte(decrypted),&res)
	if err!=nil{
		fmt.Println("err decrypted:",err)
	}

	return res,nil
}

func DecryptTransItem(req models.DecReqTransItem)(models.ReqTransItem,error){
	res:=models.ReqTransItem{}
	chanReq:=make(chan string)
	go DecryptFunc(req.Req,chanReq)

	decrypted:=<-chanReq
	// fmt.Println("decrypted:",decrypted)
	err:=json.Unmarshal([]byte(decrypted),&res)
	if err!=nil{
		fmt.Println("err decrypted:",err)
	}

	return res,nil
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