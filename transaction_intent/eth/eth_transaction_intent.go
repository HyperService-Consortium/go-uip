package EthProcessor

type Processor struct {
}

func (p *Processor) CheckAddress(addr []byte) bool {
	return len(addr) == 32
}

var p = new(Processor)

func GetProcessor() *Processor {
	return p
}
