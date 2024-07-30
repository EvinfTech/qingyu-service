package kit

//
//import (
//  "crypto"
//  "crypto/rand"
//  "crypto/rsa"
//  "crypto/x509"
//  "encoding/base64"
//  "encoding/pem"
//  "fmt"
//)
//var privateKey *rsa.PrivateKey
//var publicKey rsa.PublicKey
//func aaa()  {
//  var err error
//  privateKey, err = rsa.GenerateKey(rand.Reader, 2048)
//  if err != nil {
//    panic(err)
//  }
//
//  // The public key is a part of the *rsa.PrivateKey struct
//  publicKey = privateKey.PublicKey
//
//}
//func getPrivRSA() *rsa.PrivateKey {
//  block, _ := pem.Decode([]byte(privateKey))
//
//  if prk, err := x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
//    return nil // 获取失败
//  } else {
//    return prk // 读取成功
//  }
//}
//
//func getPubRSA(pubKey string) *rsa.PublicKey {
//  block, _ := pem.Decode([]byte(pubKey))
//
//  if pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes); err != nil {
//    return nil // 获取失败
//  } else {
//    return pubInterface.(*rsa.PublicKey) // 读取成功
//  }
//}
