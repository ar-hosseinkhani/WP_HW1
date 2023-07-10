package main

import (
	"WP/pkg/redis"
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"github.com/fxtlabs/primes"
	"math"
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

const PrimeBound = 14

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetSHA1EncodedString(input string) string {
	h := sha1.New()
	h.Write([]byte(input))
	shaHash := hex.EncodeToString(h.Sum(nil))
	return shaHash
}

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GeneratePG() (int32, int32) {
	// TODO: some of these constraints are not necessary
	p := generatePrimeNumber(PrimeBound)
	for p < 3 {
		p = generatePrimeNumber(PrimeBound)
	}
	g := rand.Intn(p)
	for g == 0 {
		g = rand.Intn(p)
	}
	return int32(p), int32(g)
}

func generatePrimeNumber(bound int) int {
	primeNums := primes.Sieve(bound)
	index := rand.Intn(bound) % len(primeNums)
	return primeNums[index]
}

func SetRedis(ctx context.Context, redisDB *redis.Redis, key string, value interface{}, expiration time.Duration) error {
	valueBytes, err := json.Marshal(value)
	if err == nil {
		err = redisDB.Set(ctx, key, string(valueBytes), expiration).Err()
	}
	return err
}

func CalculatePublic(g int32, p int32, b int32) int32 {
	return int32(math.Pow(float64(g), float64(b))) % p
}

func CalculatePrivate(p int32, A int32, b int32) int32 {
	return int32(math.Pow(float64(A), float64(b))) % p
}
