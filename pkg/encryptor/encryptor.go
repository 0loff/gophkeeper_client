package encryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"github.com/0loff/gophkeeper_client/internal/logger"
	"go.uber.org/zap"
)

func generateRandom(size int) ([]byte, error) {
	// генерируем криптостойкие случайные байты в b
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func Encrypt(data, uid string) ([]byte, error) {
	// src := []byte(data)

	// key := []byte(uid)

	// aesBlock, err := aes.NewCipher(key)
	// if err != nil {
	// 	logger.Log.Error("Error during encrypting", zap.Error(err))
	// 	return nil, err
	// }

	// aesgcm, err := cipher.NewGCM(aesBlock)
	// if err != nil {
	// 	logger.Log.Error("Error during encrypting", zap.Error(err))
	// 	return nil, err
	// }

	// nonce, err := generateRandom(aesgcm.NonceSize())
	// if err != nil {
	// 	logger.Log.Error("Error during encrypting", zap.Error(err))
	// 	return nil, err
	// }

	// dst := aesgcm.Seal(nil, nonce, src, nil)

	// return dst, nil

	block, err := aes.NewCipher([]byte(uid))
	if err != nil {
		logger.Log.Error("Error encrypting block creation", zap.Error(err))
		return nil, err
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		logger.Log.Error("Error during encrypting", zap.Error(err))
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		logger.Log.Error("Error during encrypting", zap.Error(err))
		return nil, err
	}

	ciphertext := aesgcm.Seal(nil, nonce, []byte(data), nil)
	ciphertext = append(ciphertext, ciphertext...)

	return ciphertext, nil
}
