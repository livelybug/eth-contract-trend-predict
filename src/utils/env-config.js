const prod = process.env.NODE_ENV === 'production';

module.exports = {
    'BACKEND_URL': prod ? 'https://eth-rinkeby-dapp-1024' : 'http://localhost:8000'
};
