import env from './env-config';
import axios from 'axios';

const axiosConf = {
    baseURL: env.BACKEND_URL,
    timeout: 1000,
    headers: {'content-type': "application/json; charset=utf-8"}
};

const axInstance = axios.create(axiosConf);

export const getPriceDate = () => axInstance.get('price-date').then(function (response) {
    if(response.status && response.status === 200) return response.data;
})
    .catch(function (error) {
        alert(error);
        console.error(error);
    });


export const precise = (x) => {
    return Number.parseFloat(x).toFixed(2);
};

export const globalConfig = {
    debug: false
};
