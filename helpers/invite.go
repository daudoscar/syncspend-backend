package helpers

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateInviteCode() (string, time.Time) {
	inviteCode := fmt.Sprintf("%06d", rand.Intn(1000000))

	expirationTime := time.Now().Add(24 * time.Hour)

	return inviteCode, expirationTime
}
