package helpers

import (
        "crypto/sha256"
        "fmt"
        "io"
)

func EncryptPassword(data []string) string {
        hash := sha256.New()
        for _, v := range data {
                io.WriteString(hash, v)
        }
        return fmt.Sprintf("%x", hash.Sum(nil))
}
