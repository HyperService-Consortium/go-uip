package prover

import (
	"encoding/hex"
	"fmt"
	"testing"

	crypto "github.com/HyperService-Consortium/go-uip/internal/crypto"
)

func TestEthProve(t *testing.T) {
	var accountProof = []string{"0xf90151a049e4cb91861a702fa596a23ed9b4de7f8f461d6c10065abf13329a54b50ea99e80a0ea73ca602695062343891de51c9c2aa59fa12f5aa7e1d65a24b4a8fada74b1baa03c8a9b18b3b4d5f0f20231b5bdca1c79205b49720a9b9ee256e1effcfa1debcea0597e25b6fb91a724eb9a6f046d3e4224232d0a055460cf39e793d62b2c4b414fa0814fb016bb65113e7e9b80254b615caf41d6dcd7aeece6987dc28781f3762ea9a033ca2637b94b6eb8df0e1deffd6f4b3dc7db6bde632eb95c98e1fc1cbb0ee76b8080a0f3e1ad665e0b629a2e914b971291b9822b2d3dd31dfe6e5b5fb714c03d5dc2d1a0b369852eb3019dfb2da8270b9052feef1b69b30c8da64d5a555c8b9241eb4ebea01b18478eb0fd59ccea5248bc6c52039c1537053fd36d13b0951060c42264342680a00812d65fb02a188643591533a178a323dedea2e9e388e7a7ac69c3342e3ee724808080", "0xf87180a0314d7dba2b562fb3777e4efa09ce1a92d4b2d9cf01b8d413f2d4f047984e6cec808080a04ed389f42aa6995eca849e6fcf96a32ced81fe1e1bd2f1d2cdb03c34154a7a22808080808080a025f8c45eb48ad1d050bbd03e269ad41869e85de4a533cb515b5a2f436db320d980808080", "0xf873a0201653ff4a6d0660d11912ce77346791b671daab4947be800a7da66978666d54b850f84e378a01233cdb0fbbf668c0b3a056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421a0c5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470"}

	var key = "0ac45f1e6b8d47ac4c73aee62c52794b5898da9f"

	var w, err = hex.DecodeString(key)
	if err != nil {
		return
	}

	key = hex.EncodeToString(crypto.Keccak256(w))

	fmt.Println(key)

	fmt.Println(VerifyProof("0xea165aeb2a28fbaef6c538444f9cc33f7190f5205273451ad348ca0162dc863e", key, "0xf84e378a01233cdb0fbbf668c0b3a056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421a0c5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470", accountProof))
}
