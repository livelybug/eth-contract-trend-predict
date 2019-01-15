const assert = require('assert');
const ganache = require('ganache-cli');
const Web3 = require('web3');
const compiledBetStorage = require('../eth/build/:BetStorage.json');
const compiledPricePredict = require('../eth/build/:PricePredict.json');

const web3 = new Web3(ganache.provider(options));
const blockGasLimit = 900000000;
const options = { gasLimit: blockGasLimit };

let accounts = undefined;
let predict = undefined;
let maxRange = 2;

beforeEach(async () => {
    accounts = await web3.eth.getAccounts();

    let balance = await web3.eth.getBalance(accounts[0]);
    console.log(balance);

    predict = await new web3.eth.Contract(JSON.parse(compiledPricePredict.interface))
        .deploy({
            data: compiledPricePredict.bytecode,
            arguments: [maxRange, 100]
        })
        .send({ from: accounts[0], gas: blockGasLimit });

});

describe('Predict Contract', () => {
    it('deploy a predict contract', () => {
        assert.ok(predict.options.address)
    });

    it('manager should be accout0', async () => {
        const manager = await predict.methods.manager().call();
        assert.equal(manager, accounts[0]);
    });

    it('verify storage max range', async () => {
        const storeAdd = await predict.methods.storeAdds(maxRange).call();
        const betStorage = await new web3.eth.Contract(
            JSON.parse(compiledBetStorage.interface),
            storeAdd
        );
        const _range = await betStorage.methods.range().call();
        assert.equal(_range, maxRange);
    });

    it('bet on a prediction and verify the bet on storage', async () => {
        await predict.methods.bet(maxRange).send({
            from: accounts[0],
            value: web3.utils.toWei('0.02', "ether"),
            gas: blockGasLimit
        });

        const storeAdd = await predict.methods.storeAdds(maxRange).call();
        const betStorage = await new web3.eth.Contract(
            JSON.parse(compiledBetStorage.interface),
            storeAdd
        );
        const wager0 = await betStorage.methods.betMapping(accounts[0]).call();
        assert.equal(wager0, web3.utils.toWei('0.02', "ether"));

        const bettor = await betStorage.methods.bettors(0).call();
        assert.equal(bettor, accounts[0]);
    });

    it('verify winner\' balance after resolving', async () => {
        let balanceBefore = await web3.eth.getBalance(accounts[0]);

        await predict.methods.bet(maxRange).send({
            from: accounts[0],
            value: web3.utils.toWei('0.02', "ether"),
            gas: blockGasLimit
        });

        await predict.methods.bet(maxRange - 1).send({
            from: accounts[1],
            value: web3.utils.toWei('0.3', "ether"),
            gas: blockGasLimit
        });

        await predict.methods.bet(maxRange - 2).send({
            from: accounts[2],
            value: web3.utils.toWei('0.4', "ether"),
            gas: blockGasLimit
        });

        await predict.methods.resolve(maxRange).send({
            from: accounts[0],
            gas: blockGasLimit
        });


        let balanceAfter = await web3.eth.getBalance(accounts[0]);

        assert( (balanceAfter - balanceBefore) > web3.utils.toWei('0.61', "ether"));
    });
});
