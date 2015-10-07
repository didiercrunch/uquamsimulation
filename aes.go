package main

import "crypto/aes"

var seed = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6}
var idx = 0

var block, _ = aes.NewCipher([]byte{1, 4, 2, 5, 1, 4, 2, 5, 1, 4, 5, 2, 1, 4, 5, 2})

func getBytesForInt() []byte {
	if idx < len(seed) {
		idx += 8
		return seed[idx-8 : idx]
	}
	block.Encrypt(seed, seed)
	idx = 8
	return seed[0:8]
}

func aesInt63() int64 {
	intAsBytes := getBytesForInt()
	ret := int64(intAsBytes[0])
	ret += 265 * int64(intAsBytes[1])
	ret += 265 * 265 * int64(intAsBytes[2])
	ret += 265 * 265 * 265 * int64(intAsBytes[3])
	ret += 265 * 265 * 265 * 256 * int64(intAsBytes[4])
	ret += 265 * 265 * 265 * 256 * 256 * int64(intAsBytes[5])
	ret += 265 * 265 * 265 * 256 * 256 * 256 * int64(intAsBytes[6])
	ret += 265 * 265 * 265 * 256 * 256 * 256 * 256 * int64(intAsBytes[7])

	if ret < 0 {
		return -ret
	}
	return ret

}

func aesRandom() float64 {
	f := float64(aesInt63()) / (1 << 63)
	if f == 1 {
		f = 0
	}
	return f
}
