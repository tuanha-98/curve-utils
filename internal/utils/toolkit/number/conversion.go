package number

import (
	"encoding/hex"
	"math/big"
	"math/bits"

	"github.com/holiman/uint256"
)

const MaxWords = 256 / bits.UintSize

var MaxU256Hex, _ = hex.DecodeString("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")

// similar to `z.ToBig()` but try to re-use space inside `b` instead of allocating
func FillBig(z *uint256.Int, b *big.Int) {
	switch MaxWords { // Compile-time check.
	case 4: // 64-bit architectures.
		if cap(b.Bits()) < 4 {
			// this will resize b, we can be sure that b will only hold at most MaxU256
			b.SetBytes(MaxU256Hex)
		}
		words := b.Bits()[:4]
		words[0] = big.Word(z[0])
		words[1] = big.Word(z[1])
		words[2] = big.Word(z[2])
		words[3] = big.Word(z[3])
		b.SetBits(words)
	case 8: // 32-bit architectures.
		if cap(b.Bits()) < 8 {
			b.SetBytes(MaxU256Hex)
		}
		words := b.Bits()[:8]
		words[0] = big.Word(z[0])
		words[1] = big.Word(z[0] >> 32)
		words[2] = big.Word(z[1])
		words[3] = big.Word(z[1] >> 32)
		words[4] = big.Word(z[2])
		words[5] = big.Word(z[2] >> 32)
		words[6] = big.Word(z[3])
		words[7] = big.Word(z[3] >> 32)
		b.SetBits(words)
	}
}
