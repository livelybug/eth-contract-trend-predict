const HDWalletProvider = require('truffle-hdwallet-provider');
const Web3 = require('web3');
const predict = require('./build/:PricePredict');

const provider = new HDWalletProvider(
    'decline sauce shallow fame mention cause loop maximum force blossom element ethics',
    'https://rinkeby.infura.io/v3/26f18c507ae445e487f4761f0beae803'
);

const web3 = new Web3(provider);

const deploy = async () => {
    accounts = await web3.eth.getAccounts();
    console.log('To deploy via account: ', accounts[0]);
    const result = await new web3.eth.Contract(JSON.parse(predict.interface))
        .deploy({data: '0x' + predict.bytecode, arguments: [2, 100]})
        .send({from: accounts[0]});

    console.log('Deployed to address: ', result.options.address);
};

deploy();
