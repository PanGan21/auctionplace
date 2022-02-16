compile-contract:
	solc --optimize --abi ./contracts/solidity/AuctionPlace.sol -o contracts/bin

generate-contract-interface:
	abigen --sol=contracts/solidity/AuctionPlace.sol --pkg=contracts --out=contracts/interfaces/AuctionPlace.go

deploy:
	go run cmd/main.go deploy

.PHONY: compile-contract, generate-contract-interface, deploy