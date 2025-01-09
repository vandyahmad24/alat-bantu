package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"math/big"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// GenerateHashFromString takes a string and generates a bcrypt hash from it,
// returning the hash string and error (if any).
func GenerateHashFromString(str string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// ComparePassword compares a bcrypt hashed password with its possible plaintext equivalent.
// Returns true if the passwords match, otherwise false.

func ComparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	return err == nil
}

// createHash takes a string and creates a SHA256 hash from it, returning the hash as a string.
func createHash(key string) string {
	hash := sha256.Sum256([]byte(key))
	return fmt.Sprintf("%x", hash[:])
}

// NewEncryptAES256 encrypts a plaintext string using AES-256-CFB.
// It returns the base64-encoded ciphertext as a string.
// The aesSecret is the secret key used for encryption, and must be 32 bytes long.
// An error string is returned if the encryption process encounters an error.
func NewEncryptAES256(aesSecret, PlainText string) string {

	plaintext := []byte(PlainText)

	block, err := aes.NewCipher([]byte(aesSecret))
	if err != nil {
		return err.Error()
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err.Error()
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.StdEncoding.EncodeToString(ciphertext)
}

// NewDecryptAES256 decrypts a base64-encoded ciphertext string using AES-256-CFB.
// The aesSecret is the secret key used for decryption, and must be 32 bytes long.
// An error string is returned if the decryption process encounters an error.
func NewDecryptAES256(aesSecret, ChiperText string) string {

	ciphertext, err := base64.StdEncoding.DecodeString(ChiperText)
	if err != nil {
		return err.Error()
	}

	block, err := aes.NewCipher([]byte(aesSecret))
	if err != nil {
		return err.Error()
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return err.Error()
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)
	return fmt.Sprintf("%s", ciphertext)
}

// GenerateRandomString returns a string of length n consisting of randomly
// generated characters from the alphabet of digits and uppercase letters.
// It returns an error if the random number generator fails.
func GenerateRandomString(n int) (string, error) {
	const characters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(characters))))
		if err != nil {
			return "", err
		}
		result[i] = characters[num.Int64()]
	}
	return string(result), nil
}

/*************  ✨ Codeium Command ⭐  *************/
// ChangeFormatTime converts a time string in RFC3339 format to a string in the format
// "2006-01-02 15:04:05". It returns an empty string if the input is invalid.
/******  0fe4236e-dacf-4152-b41e-79f35aea8fa8  *******/
func ChangeFormatTime(input string) string {
	parsedTime, err := time.Parse(time.RFC3339, input)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return ""
	}

	// Mengubah waktu menjadi format yang diinginkan
	formattedTime := parsedTime.Format("2006-01-02 15:04:05")

	return formattedTime

}
