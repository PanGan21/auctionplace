# VERSION=0.0.1

compile-contract:
	solc --optimize --abi ./contracts/solidity/AuctionPlace.sol -o contracts/bin

generate-contract-interface:
	abigen --sol=contracts/solidity/AuctionPlace.sol --pkg=contracts --out=contracts/interfaces/AuctionPlace.go

clean:
  @rm -rf ./build

build: clean
  @$(GOPATH)/bin/goxc \
    -bc="darwin,amd64" \
    -pv=$(VERSION) \
    -d=build \
    -build-ldflags "-X main.VERSION=$(VERSION)"

# version:
#   @echo $(VERSION)

.PHONY: compile-contract, generate-contract-interface, clean, version, build