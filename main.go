package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethCmn "github.com/ethereum/go-ethereum/common"
)

type Bazooka struct {
	LoadedContracts map[string]Contract
	Config          Config
}

func (b *Bazooka) LoadContract(contract Contract, key string) {
	b.LoadedContracts[key] = contract
}

func (b *Bazooka) LoadConfig(url) {
	b.Config = NewConfig(url)
}

// Call provides the ability to call functions of any loaded smart contract
func (b *Bazooka) Call() {

}

func (b *Bazooka) Fire() {

}

type Contract struct {
	Interface abi.ABI
	Address   ethCmn.Address
	Instance  interface{}
}

type Config struct {
	URL       string
	NetworkID int
}

func NewConfig(url string) Config {
	return Config{URL: url}
}
func main() {

}
