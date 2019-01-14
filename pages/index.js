/* eslint-disable jsx-a11y/anchor-is-valid */

import React from 'react';
import PropTypes from 'prop-types';
import {withStyles} from '@material-ui/core/styles';
import predict from '../eth/predict';
import Layout from '../components/Layout';
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import PredictInit from '../components/PredictInit';

const styles = theme => ({
    root: {
        flexGrow: 1
    }
});

class Index extends React.Component {
    state = {
        open: false,
        manager: '',
        minBet: 0,
        rangeMax: 0,
        poolSum: 0,
        storeAdds: []
    };

    async componentDidMount() {
        const manager = await predict.methods.manager().call();
        const minBet = await predict.methods.minBet().call();
        const rangeMax = await predict.methods.rangeMax().call();

        this.setState({manager, minBet, rangeMax});
    }


    render() {
        const {classes} = this.props;

        return (
            <div className={classes.root}>
                <Layout>
                    <h3>ETH Price Prdiction</h3>
                    <Grid container spacing={24}>
                        <Grid item xs={9} sm={5}>
                            <PredictInit initState={this.state}/>
                        </Grid>
                        <Grid item xs={12} sm={7}>
                            <Paper className={classes.paper}>xs=6 sm=3</Paper>
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
