compile-contract:
	solc --optimize --abi ./contracts/solidity/AuctionPlace.sol -o contracts/bin

generate-contract-interface:
	abigen --sol=contracts/solidity/AuctionPlace.sol --pkg=contracts --out=contracts/interfaces/AuctionPlace.go

ganache:
	ganache-cli --account="0x0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef,1000000000000000000000000000000000000000" --mnemonic="range pear quit paddle harvest glory insect tissue erupt spray sport child" 

deploy:
	go run cmd/main.go deploy

# go run cmd/main.go run auction --action user --account 0xFCAd0B19bB29D4674531d6f115237E16AfCE377c
# go run cmd/main.go run auction --action create --name pg1 --description des1 --min 2
# go run cmd/main.go run offer --action create --auctionId 1 --amount 5
# go run cmd/main.go run offer --action user --account 0xFCAd0B19bB29D4674531d6f115237E16AfCE377c
# go run cmd/main.go run trade --auctionId 1

.PHONY: compile-contract, generate-contract-interface, deploy