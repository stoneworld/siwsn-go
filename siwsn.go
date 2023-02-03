package siwsn

import (
	"github.com/dontpanicdao/caigo"
	"github.com/dontpanicdao/caigo/types"
	"math/big"
)

func SignMessage(messageHash *big.Int, privateKey string) (x, y *big.Int, err error) {
	private := types.SNValToBN(privateKey)
	x, y, err = caigo.Curve.Sign(messageHash, private)
	return
}
