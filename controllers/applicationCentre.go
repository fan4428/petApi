package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	data "petApi/data"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/blowfish"

	model "petApi/models"
)

func blowfishDecrypt(et, key []byte) []byte {
	// create the cipher
	dcipher, err := blowfish.NewCipher(key)
	if err != nil {
		// fix this. its okay for this tester program, but...
		panic(err)
	}
	// make initialisation vector to be the first 8 bytes of ciphertext.
	// see related note in blowfishEncrypt()
	div := et[:blowfish.BlockSize]
	// check last slice of encrypted text, if it's not a modulus of cipher block size, we're in trouble
	decrypted := et[blowfish.BlockSize:]
	if len(decrypted)%blowfish.BlockSize != 0 {
		panic("decrypted is not a multiple of blowfish.BlockSize")
	}
	// ok, we're good... create the decrypter
	dcbc := cipher.NewCBCDecrypter(dcipher, div)
	// decrypt!
	dcbc.CryptBlocks(decrypted, decrypted)
	return decrypted
}

func blowfishEncrypt(ppt, key []byte) []byte {
	// create the cipher
	ecipher, err := blowfish.NewCipher(key)
	if err != nil {
		// fix this. its okay for this tester program, but ....
		panic(err)
	}
	// make ciphertext big enough to store len(ppt)+blowfish.BlockSize
	ciphertext := make([]byte, blowfish.BlockSize+len(ppt))
	// make initialisation vector to be the first 8 bytes of ciphertext. you
	// wouldn't do this normally/in real code, but this IS example code! :)
	eiv := ciphertext[:blowfish.BlockSize]
	// create the encrypter
	ecbc := cipher.NewCBCEncrypter(ecipher, eiv)
	// encrypt the blocks, because block cipher
	ecbc.CryptBlocks(ciphertext[blowfish.BlockSize:], ppt)
	// return ciphertext to calling function
	return ciphertext
}

func Decrypt(ciphertext, key []byte) ([]byte, error) {
	keyBytes := []byte(key)
	block, err := aes.NewCipher(keyBytes) //选择加密算法
	if err != nil {
		return nil, err
	}
	blockModel := cipher.NewCBCDecrypter(block, keyBytes)
	plantText := make([]byte, len(ciphertext))
	blockModel.CryptBlocks(plantText, ciphertext)
	plantText = PKCS7UnPadding(plantText, block.BlockSize())
	return plantText, nil
}

func PKCS7UnPadding(plantText []byte, blockSize int) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}

//Test 获取user
func Test(c *gin.Context) {

	// aL := []byte("92891f39b5e25844e5e3e205f40205f5")
	// secretkey := []byte("r1Yw4pX9b4WQc6l") // dat key
	// // encryptedtext := blowfishEncrypt(plaintext, secretkey)
	// // decryptedtext := blowfishDecrypt(encryptedtext, secretkey)

	// ad, err := Decrypt(aL, secretkey)
	// sad := string(ad)
	// fmt.Println("plaintext=", sad)
	// fmt.Println("err=", err)

	var key string = "r1Yw4pX9b4WQc6l"

	ci, err := blowfish.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	dst2 := make([]byte, len("92891f39b5e25844e5e3e205f40205f5"))
	ci.Encrypt(dst2, []byte("92891f39b5e25844e5e3e205f40205f5"))

	fmt.Println("plaintext=", string(dst2))
	// fmt.Println("  ad=%s\n", ad)
	// secretkey := []byte("r1Yw4pX9b4WQc6l")
	// aL := []byte("cc5db42c089783aa174c4d7cc4141b5256fb85bd540aa82b")
	// fan := []byte("fanpeng")
	// fande := blowfishEncrypt(fan, aL)
	// zzza := blowfishDecrypt(aL, secretkey)
	// fanEN := blowfishDecrypt(fande, secretkey)
	// z := string(zzza)
	// f := string(fanEN)
	// fmt.Println("  decryptedtext=", z)
	// fmt.Println("  f=", f)

	// secretkey := []byte("r1Yw4pX9b4WQc6l")
	// val := []byte("DDFD58EE2140B130")
	// v, err := Decrypt(val, secretkey)
	// if err != nil {

	// 	fmt.Println(err)
	// 	panic(err)
	// }
	// sv := string(v)
	// fmt.Println(sv)

	// var key string = "r1Yw4pX9b4WQc6l"
	// var IV = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	// ci, err := blowfish.NewCipher([]byte(key))
	// if err != nil {
	// 	panic(err)
	// }

	// s := cipher.NewCFBDecrypter(ci, IV)
	// data, err := base64.URLEncoding.DecodeString("DDFD58EE2140B130")
	// if err != nil {
	// 	panic(err)
	// }
	// dst := make([]byte, len(data))
	// dst2 := make([]byte, len(data))
	// ci.Encrypt(dst2, data)

	// s.XORKeyStream(dst, data)
	// fmt.Println(string(dst))

	// fmt.Println("dst2", string(dst2))
}

//GetUserByname 获取user
func GetUserByname(c *gin.Context) {
	var result = make(chan []model.User)

	go data.GetUserByName("Colin CHIU", result)
	v := <-result
	fmt.Println(v)
}
