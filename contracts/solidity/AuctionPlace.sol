// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.0;

contract AuctionPlace {
    struct Auction {
        uint256 id;
        address payable seller;
        string name;
        string description;
        uint256 min;
        uint256 bestOfferId;
        uint256[] offerIds;
    }

    struct Offer {
        uint256 id;
        uint256 auctionId;
        address payable buyer;
        uint256 price;
    }

    mapping(uint256 => Auction) private idsToAuctions;
    mapping(uint256 => Offer) private idsToOffers;
    mapping(address => uint256[]) private usersToAuctionIds;
    mapping(address => uint256[]) private usersToOfferIds;
    uint256 private currentAuctionId = 1;
    uint256 private currentOfferId = 1;

    modifier auctionExists(uint256 _auctionId) {
        require(
            _auctionId > 0 && _auctionId < currentAuctionId,
            "Auction does not exist"
        );
        _;
    }

    function createAuction(
        string calldata _name,
        string calldata _description,
        uint256 _min
    ) external {
        require(_min > 0, "_min must be > 0");
        uint256[] memory offerIds = new uint256[](0);
        idsToAuctions[currentAuctionId] = Auction(
            currentAuctionId,
            payable(msg.sender),
            _name,
            _description,
            _min,
            0,
            offerIds
        );

        usersToAuctionIds[msg.sender].push(currentAuctionId);
        currentAuctionId++;
    }

    function createOffer(uint256 _auctionId)
        external
        payable
        auctionExists(_auctionId)
    {
        Auction storage auction = idsToAuctions[_auctionId];
        Offer storage bestOffer = idsToOffers[auction.bestOfferId];
        require(
            msg.value >= auction.min && msg.value > bestOffer.price,
            "msg.value must be greater than min and bestOffer"
        );
        auction.bestOfferId = currentOfferId;
        auction.offerIds.push(currentOfferId);
        idsToOffers[currentOfferId] = Offer(
            currentOfferId,
            auction.id,
            payable(msg.sender),
            msg.value
        );
        usersToOfferIds[msg.sender].push(currentOfferId);
        currentOfferId++;
    }

    function trade(uint256 _auctionId) external auctionExists(_auctionId) {
        Auction storage auction = idsToAuctions[_auctionId];
        Offer storage bestOffer = idsToOffers[auction.bestOfferId];
        for (uint256 i = 0; i < auction.offerIds.length; i++) {
            uint256 offerId = auction.offerIds[i];
            if (offerId != auction.bestOfferId) {
                Offer storage offer = idsToOffers[offerId];
                offer.buyer.transfer(offer.price);
            }
        }
        auction.seller.transfer(bestOffer.price);
    }

    function getAuctions() external view returns (Auction[] memory) {
        Auction[] memory _auctions = new Auction[](currentAuctionId - 1);
        if (currentAuctionId == 1) {
            return _auctions;
        }
        for (uint256 i = 1; i < currentAuctionId; i++) {
            _auctions[i - 1] = idsToAuctions[i];
        }
        return _auctions;
    }

    function getUserAuctions(address _user)
        external
        view
        returns (Auction[] memory)
    {
        uint256[] storage userAuctionIds = usersToAuctionIds[_user];
        Auction[] memory _auctions = new Auction[](userAuctionIds.length);

        for (uint256 i = 0; i < userAuctionIds.length; i++) {
            _auctions[i] = idsToAuctions[i + 1];
        }
        return _auctions;
    }

    function getUserOffers(address _user)
        external
        view
        returns (Offer[] memory)
    {
        uint256[] storage userOfferIds = usersToOfferIds[_user];
        Offer[] memory _offers = new Offer[](userOfferIds.length);
        for (uint256 i = 0; i < userOfferIds.length; i++) {
            _offers[i] = idsToOffers[i + 1];
        }
        return _offers;
    }
}
