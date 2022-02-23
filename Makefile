compile-contract:
	solc --optimize --abi ./contracts/solidity/AuctionPlace.sol -o contracts/bin

generate-contract-interface:
	abigen --sol=contracts/solidity/AuctionPlace.sol --pkg=contracts --out=contracts/interfaces/AuctionPlace.go

ganache:
	ganache-cli --account="0x0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef,1000000000000000000000000000000000000000" --mnemonic="range pear quit paddle harvest glory insect tissue erupt spray sport child" 

deploy:
	go run cmd/main.go deploy

build: 
	go build .

.PHONY: compile-contract, generate-contract-interface, deploy
