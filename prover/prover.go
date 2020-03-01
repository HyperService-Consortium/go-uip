package prover

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"

	rlp "github.com/HyperService-Consortium/go-rlp"
	crypto "github.com/HyperService-Consortium/go-uip/crypto"
)

const (
	// test-db dir
	EDB_PATH = "D:/Go Ethereum/data/geth/chaindata"
	// size of a secure-MPT's short-node
	SHORTNODE = 2
	// size of a secure-MPT's full-node
	FULLNODE = 17
	// length of a keccak256hash in string (16^64 = 256^32)
	HASHSTRINGLENGTH = 64
	// bytes of a char* (C pointer)
	CHAR_P_SIZE = 8
	// max number of onload-dbs
	MXONLOADDB = 64
	// Hash byte-Array's length
	HashLength = 32
)

var (
	hexmaps [256]uint64
)

func ProveEthProof() (bool, error) {
	return false, nil
}

// get value of key on the trie
func findPath(roothash string, path string, storagepath []string, consumed int) (string, error) {

	//key consumed
	if len(path) == 0 {
		return "", errors.New("consumed")
	}

	var node [][]byte
	fmt.Println(storagepath)
	if len(storagepath[0]) >= 2 && storagepath[0][0:2] == "0x" {
		storagepath[0] = storagepath[0][2:]
	}
	var w, err = hex.DecodeString(storagepath[0])
	if err != nil {
		return "", err
	}

	if hex.EncodeToString(crypto.Keccak256(w)) != roothash {
		fmt.Println(hex.EncodeToString(crypto.Keccak256(w)), roothash)
		return "", errors.New("not good roothash")
	}

	err = rlp.DecodeBytes(w, &node)
	if err != nil {
		return "", err
	}

	fmt.Println(node)

	switch len(node) {
	case SHORTNODE:
		firstvar, secondvar := hex.EncodeToString(node[0]), hex.EncodeToString(node[1])
		flag := firstvar[0]
		fmt.Println(firstvar, flag)
		switch flag {
		case '0':
			//compare prefix of the path with node[0]
			firstvarlen := len(firstvar)
			if (firstvarlen > len(path)) || (firstvar != path[0:firstvarlen]) {
				return "", errors.New("No exists(no this branch), in depth" + strconv.FormatInt(int64(consumed), 10) + ", querying" + roothash)
			}

			return findPath(secondvar, path[firstvarlen:], storagepath[1:], consumed+firstvarlen)
		case '1':
			//compare prefix of the path with node[0]
			firstvarlen := len(firstvar)
			if (firstvarlen > len(path)) || (firstvar != path[0:firstvarlen]) {
				return "", errors.New("No exists(no this branch), in depth" + strconv.FormatInt(int64(consumed), 10) + ", querying" + roothash)
			}

			return findPath(secondvar, path[firstvarlen:], storagepath[1:], consumed+firstvarlen)
		case '2':
			//end of proofpath
			if len(storagepath) != 1 {
				return "", errors.New("bad proof encoding")
			}

			if path != firstvar[consumed:] {
				return "", errors.New("No exists(no this branch), in depth" + strconv.FormatInt(int64(consumed), 10) + ", querying" + roothash)
			}
			return secondvar, nil
		case '3':
			//end of proofpath
			if len(storagepath) != 1 {
				return "", errors.New("bad proof encoding")
			}

			if path != firstvar[consumed:] {
				return "", errors.New("No exists(no this branch), in depth" + strconv.FormatInt(int64(consumed), 10) + ", querying" + roothash)
			}
			return secondvar, nil
		default:
			return "", errors.New("unknown flags of merkle tree node")
		}
	case FULLNODE:
		{
			fmt.Println(node[int(hexmaps[path[0]])], path)
			tryquery := hex.EncodeToString(node[int(hexmaps[path[0]])])

			fmt.Println(len(tryquery))

			if len(tryquery) == HASHSTRINGLENGTH {
				return findPath(tryquery, path[1:], storagepath[1:], consumed+1)
			} else {
				return "", errors.New("No exists, in depth" + strconv.FormatInt(int64(consumed), 10) + ", querying" + roothash)
			}
		}
	default:
		{
			return "", errors.New("Unknown node types")
		}
	}
}

func VerifyProof(rootHashStr string, key string, value string, storagepath []string) bool {
	if len(key) >= 2 && key[0:2] == "0x" {
		key = key[2:]
	}
	if len(rootHashStr) >= 2 && rootHashStr[0:2] == "0x" {
		rootHashStr = rootHashStr[2:]
	}
	for _, str := range storagepath {
		if len(str) >= 2 && str[0:2] == "0x" {
			str = str[2:]
		}
	}

	// key = append(make([]byte,0),2,9)
	toval, err := findPath(rootHashStr, key, storagepath, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		if hexstringequal(value, toval) == true {
			fmt.Println("Proved")
			return true
		} else {
			fmt.Println("key maps to", "0x"+toval+", not", "0x"+value)
			return false
		}
	}
	return false
}

// compare to hexstring without leading-zero
func hexstringequal(hexx string, hexy string) bool {
	var ofsx, ofsy = 0, 0
	if len(hexx) >= 2 && hexx[0:2] == "0x" {
		ofsx = 2
	}
	if len(hexy) >= 2 && hexy[0:2] == "0x" {
		ofsy = 2
	}
	var lenx, leny int
	for lenx = len(hexx) - 1; ofsx < lenx && (hexx[ofsx] == '0'); ofsx++ {

	}

	for leny = len(hexy) - 1; ofsy < leny && (hexy[ofsy] == '0'); ofsy++ {

	}
	return hexx[ofsx:] == hexy[ofsy:]
}

func init() {
	for idx := '0'; idx <= '9'; idx++ {
		hexmaps[idx] = uint64(idx - '0')
	}
	for idx := 'a'; idx <= 'f'; idx++ {
		hexmaps[idx] = uint64(idx - 'a' + 10)
	}
	for idx := 'A'; idx <= 'F'; idx++ {
		hexmaps[idx] = uint64(idx - 'A' + 10)
	}
}
