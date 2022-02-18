# auctionplace
A cli tool that interacts with an Ethereum smart contract for creating and trading auctions

#### Run ganache or any other provider
`make ganache`

#### Deploy the smart contract locally
`make deploy`

#### Build the project
`make build`

#### Create an auction
`./auctionplace run auction --action create --name pg1 --description des1 --min 2`

#### Fetch user auctions
`./auctionplace run auction --action user --account 0xFCAd0B19bB29D4674531d6f115237E16AfCE377c`

#### Create an offer
`./auctionplace run offer --action create --auctionId 1 --amount 5`

#### Fetch user offers
`./auctionplace run offer --action user --account 0xFCAd0B19bB29D4674531d6f115237E16AfCE377c`

#### Trade an auction with the best offer
`./auctionplace run trade --auctionId 1`
