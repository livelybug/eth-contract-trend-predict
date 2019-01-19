const prod = process.env.NODE_ENV === 'production';

module.exports = {
    'BACKEND_URL': prod ? 'https://eth-rinkeby-dapp-1024.herokuapp.com' : 'http://localhost:8000'
    // 'BACKEND_URL': prod ? 'http://65.49.229.160:8000/' : 'http://localhost:8000'
};
