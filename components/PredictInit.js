import React from 'react';
import PropTypes from 'prop-types';
import {withStyles} from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import Typography from '@material-ui/core/Typography';
import web3 from '../eth/web3';

const styles = {
    card: {
        minWidth: 275,
    },
    title: {
        fontSize: 14,
    }
};

function SimpleCard(props) {
    const {classes, initState} = props;

    return (
        <Card className={classes.card}>
            <CardContent>
                <Typography className={classes.title} color="textSecondary" gutterBottom>
                    Manager Address
                </Typography>
                <Typography component="p">
                    {initState.manager}
                </Typography>
            </CardContent>

            <CardContent>
                <Typography className={classes.title} color="textSecondary" gutterBottom>
                    Pool Amount(ETH)
                </Typography>
                <Typography component="p">
                    {web3.utils.fromWei(initState.poolSum.toString())}
                </Typography>
            </CardContent>

            <CardContent>
                <Typography className={classes.title} color="textSecondary" gutterBottom>
                    Minimum Bet(ETH)
                </Typography>
                <Typography component="p">
                    {web3.utils.fromWei(initState.minBet.toString())}
                </Typography>
            </CardContent>

        </Card>
    );
}

SimpleCard.propTypes = {
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(SimpleCard);
