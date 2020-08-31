package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func generatPrivateAndPublicKey(privateKey, publicKey io.Writer, bits int) error {

	// 生成私匙，提供一个随机数和私匙的长度，目前主流的长度为1024、2048、3072、4096，
	// 但1024已经不在推荐使用了。
	priKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	// 将rsa私钥序列化为ASN.1 PKCS#1 DER编码
	data := x509.MarshalPKCS1PrivateKey(priKey)

	block := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: data,
	}
	// 将私匙做pem数据编码，然后写入文件
	err = pem.Encode(privateKey, &block)
	if err != nil {
		return err
	}

	// 生成公匙
	pubKey := priKey.PublicKey
	// 将rsa公钥序列化为ASN.1 PKCS#1 DER编码
	pubKeyData := x509.MarshalPKCS1PublicKey(&pubKey)
	// 将公匙做pem数据编码，然后写入文件
	err = pem.Encode(publicKey, &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubKeyData,
	})
	if err != nil {
		return err
	}

	return nil

}

// 解析公匙
func parsingRsaPublicKey(file string) (*rsa.PublicKey, error) {
	// 读取公匙文件
	pubByte, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	// pem解码
	b, _ := pem.Decode(pubByte)
	if b == nil {
		return nil, errors.New("error public key")
	}
	// der解码，最终返回一个公匙对象
	pubKey, err := x509.ParsePKCS1PublicKey(b.Bytes)
	if err != nil {
		return nil, err
	}
	return pubKey, nil
}

// 解析私匙
func parsingRsaPrivateKey(file string) (*rsa.PrivateKey, error) {
	// 读取私匙
	priByte, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	// pem解码
	b, _ := pem.Decode(priByte)
	if b == nil {
		return nil, errors.New("error private key")
	}
	// der加密，返回一个私匙对象
	prikey, err := x509.ParsePKCS1PrivateKey(b.Bytes)
	if err != nil {
		return nil, err
	}
	return prikey, nil
}

//  rsa公匙加密
func rsaPublicKeyEncrypt(src []byte, publickey *rsa.PublicKey) ([]byte, error) {
	// 使用公匙加密数据，需要一个随机数生成器和公匙和需要加密的数据
	data, err := rsa.EncryptPKCS1v15(rand.Reader, publickey, src)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// rsa私匙解密
func rsaPrivateKeyDecrypt(src []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	// 使用私匙解密数据，需要一个随机数生成器和私匙和需要解密的数据
	data, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, src)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func main() {
	pri, _ := os.Create("pri.key")
	pub, _ := os.Create("pub.key")
	defer pri.Close()
	defer pub.Close()
	generatPrivateAndPublicKey(pri, pub, 2048)

	pubKey, _ := parsingRsaPublicKey("pub.key")                        // 解密公匙
	encryData, _ := rsaPublicKeyEncrypt([]byte("hello world"), pubKey) // 加密数据
	fmt.Printf("%x\n", encryData)

	priKey, _ := parsingRsaPrivateKey("pri.key")            // 解密私匙
	decryData, _ := rsaPrivateKeyDecrypt(encryData, priKey) // 解密数据
	fmt.Printf("%s\n", decryData)

}
