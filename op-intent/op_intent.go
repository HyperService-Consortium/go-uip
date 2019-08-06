package opintent

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"reflect"
	"unsafe"

	types "github.com/Myriad-Dreamin/go-uip/types"
	gjson "github.com/tidwall/gjson"
)

type OpIntentInitializer struct {
	degPool          *DegreePool
	largerThanLarger uint32
}

func NewOpIntentInitializer() *OpIntentInitializer {
	return &OpIntentInitializer{
		degPool: newDegreePool(),
	}
}

func (ier *OpIntentInitializer) InitOpIntent(
	opIntents types.OpIntents,
) (transactionIntents []*TransactionIntent, proposals []*MerkleProofProposal, err error) {
	contents, rawDependencies := opIntents.GetContents(), opIntents.GetDependencies()
	var intent types.BaseOpIntent
	var rtx [][]*TransactionIntent

	// todo: add merkle proof proposals
	var rps [][]*MerkleProofProposal
	var hacker uint32
	var bn []byte
	var rbn = (*reflect.SliceHeader)(unsafe.Pointer(&bn))
	var sh *reflect.StringHeader
	var nameMap = make(map[uint32]uint32)
	for idx, content := range contents {
		err = json.Unmarshal(content, &intent)
		if err != nil {
			return nil, nil, err
		}
		switch intent.OpType {
		case "Payment":
			if transactionIntents, proposals, err = ier.InitPaymentOpIntent(intent.Name, content); err != nil {
				return nil, nil, err
			} else {
				rtx = append(rtx, transactionIntents)
				rps = append(rps, proposals)
			}

		case "ContractInvocation":
			if transactionIntents, proposals, err = ier.InitContractInvocationOpIntent(intent.Name, content); err != nil {
				return nil, nil, err
			} else {
				rtx = append(rtx, transactionIntents)
				rps = append(rps, proposals)
			}
			// if tx, err = ier.InitContractInvocationOpIntent(intent.Name, intent.SubIntent); err != nil {
			// 	return nil, err
			// } else {
			// 	rtx = append(rtx, tx)
			// }

		default:
			return nil, nil, invalidOpType
		}

		sh = (*reflect.StringHeader)(unsafe.Pointer(&intent.Name))
		rbn.Data = sh.Data
		rbn.Cap = sh.Len
		rbn.Len = sh.Len
		s := md5.Sum(bn)
		hacker = *(*uint32)(unsafe.Pointer(&s[0]))
		nameMap[hacker] = uint32(idx)
	}

	var deps []Dependency
	var dep Dependency
	var res, sres gjson.Result
	var ok bool
	var sn string
	for _, rawDep := range rawDependencies {
		res = gjson.ParseBytes(rawDep)
		if sres = res.Get("left"); !sres.Exists() {
			return nil, nil, errors.New("left property not found in dependency")
		}
		sn = sres.String()
		sh = (*reflect.StringHeader)(unsafe.Pointer(&sn))
		rbn.Data = sh.Data
		rbn.Cap = sh.Len
		rbn.Len = sh.Len
		s := md5.Sum(bn)
		dep.Src = *(*uint32)(unsafe.Pointer(&s[0]))

		if dep.Src, ok = nameMap[dep.Src]; !ok {
			return nil, nil, errors.New("not such name(left)")
		}

		if sres = res.Get("right"); !sres.Exists() {
			return nil, nil, errors.New("right property not found in dependency")
		}

		sn = sres.String()
		sh = (*reflect.StringHeader)(unsafe.Pointer(&sn))
		rbn.Data = sh.Data
		rbn.Cap = sh.Len
		rbn.Len = sh.Len
		s = md5.Sum(bn)
		dep.Dst = *(*uint32)(unsafe.Pointer(&s[0]))

		if dep.Dst, ok = nameMap[dep.Dst]; !ok {
			return nil, nil, errors.New("not such name(right)")
		}

		if sres = res.Get("dep"); !sres.Exists() {
			return nil, nil, errors.New("dep property not found in dependency")
		}
		switch sres.String() {
		case "before":
		case "after":
			dep.Src, dep.Dst = dep.Dst, dep.Src
		default:
			if dep.Dst, ok = nameMap[dep.Src]; !ok {
				return nil, nil, errors.New("not such dependency relationship type")
			}
		}
		deps = append(deps, dep)
	}

	// WARNING: ier.TopologicalSort assume that the size of total intents is <= 2 * len(rtx)
	if err = ier.TopologicalSort(rtx, deps); err != nil {
		return nil, nil, err
	}

	transactionIntents = nil
	proposals = nil

	for _, rt := range rtx {
		transactionIntents = append(transactionIntents, rt...)
	}
	for _, rp := range rps {
		proposals = append(proposals, rp...)
	}
	return
}
