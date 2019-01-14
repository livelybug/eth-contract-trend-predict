import React from 'react';
import Head from 'next/head';
import Header from './Header';

export default props => {
    return (
        <div>
            <Head>
                <title>Prediction</title>
                <meta name="viewport" content="initial-scale=1.0, width=device-width" />
            </Head>

            <Header/>
            {props.children}
        </div>
    );
};
