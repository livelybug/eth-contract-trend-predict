import web3 from './web3';
import predict from './predict';

const compiledBetStorage = require('../eth/build/:BetStorage.json');

const getBetStorage = async (range) => {
    const storeAdd = await predict.methods.storeAdds(range).call();
    const betStorage = await new web3.eth.Contract(
        JSON.parse(compiledBetStorage.interface),
        storeAdd
    );
    return betStorage;
}

export default getBetStorage;
