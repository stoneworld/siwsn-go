package siwsn

import (
	"fmt"
	"github.com/dontpanicdao/caigo"
	"github.com/dontpanicdao/caigo/types"
	"math/big"
	"testing"
)

func TestSignMessage(t *testing.T) {
	privateKey := "12345"
	x, y, _ := caigo.Curve.PrivateToPoint(types.SNValToBN(privateKey))
	fmt.Printf("PublicKey: %v\n", x)

	//address := "0x03A01f3409EfD747b9B6629C2FFdBC63976Ff1EE173648F437069bd33A9424AB"
	hash, _ := caigo.Curve.PedersenHash(
		[]*big.Int{
			//types.HexToBN(address),
			types.StrToBig("4321"),
		})

	fmt.Printf("hash: %v\n", hash)
	r, s, err := SignMessage(hash, privateKey)
	fmt.Printf("r: %v\ns: %v\n", r, s)

	if err != nil {
		t.Errorf("Error signing message %v\n", err)
	}
	if !caigo.Curve.Verify(hash, r, s, x, y) {
		t.Errorf("Error verifying signature %v\n", err)
	}
}
