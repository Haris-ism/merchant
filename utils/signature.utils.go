package utils

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/json"
	"errors"
	"fmt"
	"merchant/controllers/models"
	"merchant/protogen/merchant"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

func SignatureValidationGrpc(reqHeader metadata.MD,reqBody *merchant.ReqTransItemsModel)error{
	body,err:=json.Marshal(reqBody)
	hash:=Signature(string(body),reqHeader.Get("timestamp")[0])
	// fmt.Println("hash:",hash)
	// fmt.Println("sig:",reqHeader.Signature)
	if hash!=reqHeader.Get("signature")[0]{
		logrus.Error("err:",err)
		return errors.New("Invalid Signature")
	}
	return nil
}

func SignatureValidation(reqHeader models.ReqHeader,reqBody models.DecReqTransItem)error{

	body,err:=json.Marshal(reqBody)
	hash:=Signature(string(body),reqHeader.TimeStamp)
	fmt.Println("hash:",hash)
	fmt.Println("sig:",reqHeader.Signature)
	if hash!=reqHeader.Signature{
		logrus.Error("err:",err)
		return errors.New("Invalid Signature")
	}
	return nil
}

func Signature(req string,ts string)string{
	key:=GetEnv("SIG_KEY")
	data:=req+"&"+ts+"&"+key
	// fmt.Println("data:",data)
	res:=HashSha512(key,data)
	// fmt.Println("res:",res)
	return res
}
func HashSha512(secret, data string) string {
	hash := hmac.New(sha512.New, []byte(secret))
	hash.Write([]byte(data))
	return fmt.Sprintf("%x",hash.Sum(nil))
}