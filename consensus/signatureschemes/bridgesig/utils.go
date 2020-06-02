package bridgesig

import (
	"math/big"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

const CBridgeSigSz = 65

func DecodeECDSASig(sig []byte) (
	v byte,
	r []byte,
	s []byte,
	err error,
) {
	if len(sig) != CBridgeSigSz {
		err = errors.New("wrong input")
		return
	}
	v = byte(sig[64] + 27)
	r = sig[:32]
	s = sig[32:64]
	return
}

// B2ImN is Bytes to Int mod N, with N is secp256k1 curve order
func B2ImN(bytes []byte) *big.Int {
	x := big.NewInt(0)
	x.SetBytes(ethcrypto.Keccak256Hash(bytes).Bytes())
	for x.Cmp(ethcrypto.S256().Params().N) != -1 {
		x.SetBytes(ethcrypto.Keccak256Hash(x.Bytes()).Bytes())
	}
	return x
}
