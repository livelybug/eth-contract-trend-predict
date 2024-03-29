/* eslint-disable jsx-a11y/anchor-is-valid */

import React from 'react';
import PropTypes from 'prop-types';
import {withStyles} from '@material-ui/core/styles';
import Grid from "@material-ui/core/Grid";
import ReactLoading from 'react-loading';

import predict from '../eth/predict';
import betStorege from '../eth/betStorage';
import Layout from '../src/components/Layout';
import PredictInit from '../src/components/PredictInit';
import SubmitBet from '../src/components/SubmitBet';
import {getPriceDate} from '../src/utils/utils'

const styles = () => ({
    root: {
        flexGrow: 1
    },
    center: {
        marginLeft: "auto",
        marginRight: "auto",
        width: "50%"
    }
});

const trendsMap = {
    2: {predict: 'Up'},
    1: {predict: 'Unchanged'},
    0: {predict: 'Down'}
};

class Index extends React.Component {
    state = {
        open: false,
        manager: '',
        minBet: 0,
        rangeMax: 0,
        poolSum: 0,
        storeAdds: [],
        submitBets: [],
        isLoading: true,
        priceDate: {}
    };

    async componentDidMount() {
        const manager = await predict.methods.manager().call();
        const minBet = await predict.methods.minBet().call();
        const rangeMax = await predict.methods.rangeMax().call();
        const poolSum = await predict.methods.poolSum().call();
        this.setState({manager, minBet, rangeMax, poolSum});

        await this.getSubmitBets();

        this.setState({priceDate: await getPriceDate()});

        setTimeout(() => this.setState({isLoading: false}), 300);
    }

    async getSubmitBets() {
        let submitBets = [];

        for (let range = 0; range <= this.state.rangeMax; range++) {
            const betStorage = await betStorege(range);
            const betPool = await betStorage.methods.betPool().call();
            const bettorsLen = await betStorage.methods.getBettorsLen().call();
            trendsMap[range].betPool = betPool;
            trendsMap[range].bettorsLen = bettorsLen;
            trendsMap[range].poolSum = this.state.poolSum;

            submitBets.push(<SubmitBet key={range} range={range} predict={trendsMap[range]}/>);
        }

        this.setState({submitBets});
    }

    render() {
        const {isLoading} = this.state;
        const {classes} = this.props;

        if (isLoading)
            return (
                <ReactLoading className={classes.center} type="spinningBubbles" color="#cbcbcd" height={'30%'}
                              width={'30%'}/>
            );

        return (
            <div className={classes.root}>
                <Layout>
                    <h3>Coinbase ETH Price Prdiction, with Contract on ETH Rinkeby Network</h3>
                    <h3>ETH price on {this.state.priceDate.old.DATE} is {this.state.priceDate.old.USD} USD</h3>
                    <h3>Comapred with {this.state.priceDate.old.USD} USD, predict the ETH price on {this.state.priceDate.new.DATE}. Is it up, down or unchanged?</h3>
                    <Grid container spacing={24}>
                        <Grid item xs={9} sm={5}>
                            <PredictInit initState={this.state}/>
                        </Grid>
                        <Grid item xs={12} sm={7}>
                            {this.state.submitBets}
                        </Grid>
                    </Grid>
                </Layout>
            </div>
        );
    }
}

Index.propTypes = {
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(Index);
