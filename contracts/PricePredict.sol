pragma solidity ^0.4.2;

contract PricePredict {

    address public manager;
    uint public minBet;
    uint public rangeMax;
    uint public poolSum;
    address[] public storeAdds; 
    
    constructor(uint range, uint minimunt) public {
        require(range > 0);
        require(minimunt > 0);
        
        manager = msg.sender;
        minBet = minimunt;
        rangeMax = range;

        for(uint i = 0; i <= rangeMax; i++) {
            address betStorageAdd = new BetStorage(i);
            storeAdds.push(betStorageAdd);
        }        
    }

    function bet(uint predict) public payable{
        require(msg.value > minBet);
        require(predict <= rangeMax);

        address storeAdd1 = storeAdds[predict];
        BetStorage betStorage1 = BetStorage(storeAdd1);
        if(betStorage1.betMapping(msg.sender) == 0) {
           betStorage1.addBettor(msg.sender); 
        }
        betStorage1.addBetMapping(msg.sender, msg.value);
        betStorage1.addBetPool(msg.value);
        
        poolSum += msg.value;
    }

    function resolve(uint predict) public payable {
        require(msg.sender == manager);
        require(predict <= rangeMax);

        address storeAdd2 = storeAdds[predict];
        BetStorage betStorage2 = BetStorage(storeAdd2);
        require(betStorage2.betPool() > 0);        

        // loop the bettors of winning group
        uint bettorsLen = betStorage2.getBettorsLen();
        for(uint i = 0; i < bettorsLen; i++ ) {
            address aWinner = betStorage2.bettors(i);
            uint refund = betStorage2.betMapping(aWinner) / uint128(betStorage2.betPool()) * poolSum;
            aWinner.transfer(refund);
        }
        
        // reset
        storeAdds = new address[](0);
        for(i = 0; i <= rangeMax; i++) {
            address betStorageAdd = new BetStorage(i);
            storeAdds.push(betStorageAdd);
        }
        
        poolSum = 0;
    }

}


contract BetStorage {
    
    mapping(address => uint) public betMapping;
    address[] public bettors;
    uint public range;
    uint public betPool;

    constructor(uint i) public {
        bettors = new address[](0);
        range = i;
        betPool = 0;
    }

    function getBettorsLen() public view returns(uint) {
        return bettors.length;
    }
    
    function addBettor(address bettor) public {
        bettors.push(bettor);
    }
    
    function addBetMapping(address bettor, uint wager) public {
        betMapping[bettor] += wager;
    }
    
    function addBetPool(uint wager) public {
        betPool += wager;
    }
}