/* eslint-disable jsx-a11y/anchor-is-valid */

import React from 'react';
import PropTypes from 'prop-types';
import {withStyles} from '@material-ui/core/styles';
import Grid from "@material-ui/core/Grid";

import predict from '../eth/predict';
import betStorege from '../eth/betStorage';
import Layout from '../components/Layout';
import PredictInit from '../components/PredictInit';
import SubmitBet from '../components/SubmitBet';

const styles = () => ({
    root: {
        flexGrow: 1
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
        submitBets: []
    };

    async componentDidMount() {
        const manager = await predict.methods.manager().call();
        const minBet = await predict.methods.minBet().call();
        const rangeMax = await predict.methods.rangeMax().call();
        const poolSum = await predict.methods.poolSum().call();

        this.setState({manager, minBet, rangeMax, poolSum});
        this.getSubmitBets();
    }

    async getSubmitBets() {
        let submitBets = [];

        for(let range = 0; range <= this.state.rangeMax; range++) {
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
        const {classes} = this.props;

        return (
            <div className={classes.root}>
                <Layout>
                    <h3>ETH Price Prdiction (coundown - 00:00:00)</h3>
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
