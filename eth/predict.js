import web3 from './web3';
import predict from './build/:PricePredict';

const instance = new web3.eth.Contract(
    JSON.parse(predict.interface),
    '0xd67caF6E55ADF71005fa0E9Cc0D2A5F7E8522EF9'
);

export default instance;
