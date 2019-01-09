pragma solidity ^0.4.2;

contract PricePredict {

    struct BetStruct{
        mapping(address => uint) betMapping;
        address[] bettors;
        uint range;
        uint betpool;
    }

    address public manager;
    BetStruct[] public betStructs;
    uint public minBet;
    uint public rangeMax;
    uint public poolSum;
    
    constructor(uint range, uint minimunt) public {
        manager = msg.sender;
        minBet = minimunt;
        rangeMax = range;
        
        for(uint i = 0; i <= rangeMax; i++) {
            BetStruct memory betStruct = BetStruct({
                bettors: new address[](0),
                range: i,
                betpool: 0
            });
            betStructs.push(betStruct);
            betStructs[i].range = i;
        }        
        
    }

    function bet(uint predict) public payable{
        require(msg.value > minBet);
        require(predict <= rangeMax);

        if(betStructs[predict].betMapping[msg.sender] == 0) {
            betStructs[predict].bettors.push(msg.sender);
        }
        betStructs[predict].betMapping[msg.sender] += msg.value;
        betStructs[predict].betpool += msg.value;
        poolSum += msg.value;
    }

    function resolve(uint predict) public payable {
        require(msg.sender == manager);
        require(predict <= rangeMax);
        require(betStructs[predict].betpool > 0);

        // loop the correct bettors
        for(uint i = 0; i < betStructs[predict].bettors.length; i++) {
            BetStruct storage betStruct = betStructs[predict];
            uint refund = betStruct.betMapping[betStruct.bettors[i]] / betStructs[predict].betpool * poolSum;
            betStruct.bettors[i].transfer(refund);
        }        

        for(i = 0; i <= rangeMax; i++) {
            address[] storage _betters = betStructs[i].bettors; 
            if(_betters.length > 0) delete betStructs[i].betMapping[_betters[0]];
            betStructs[i].bettors = new address[](0);
            betStructs[i].range = i;
            betStructs[i].betpool = 0;
        }
        
        poolSum = 0;
    }

    function getBettors(uint predict) public view returns(address[]) {
        return betStructs[predict].bettors;
    }

}
