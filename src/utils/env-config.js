const prod = process.env.NODE_ENV === 'production';

module.exports = {
    'BACKEND_URL': prod ? 'https://api.example.com' : 'http://localhost:8000'
};
