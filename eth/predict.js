import web3 from './web3';
import predict from './build/:PricePredict';

const instance = new web3.eth.Contract(
    JSON.parse(predict.interface),
    '0xe465B7734a10F74ed6499635c13F213261843Ab6'
);

export default instance;
