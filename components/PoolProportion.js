import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import CircularProgress from '@material-ui/core/CircularProgress';
import Button from '@material-ui/core/Button';
import web3 from '../eth/web3';
import {precise} from '../src/utils/utils';

const styles = theme => ({
    button: {
        margin: theme.spacing.unit * 2
    }
});

class CircularStatic extends React.Component {

    constructor(props) {
        super(props);
    }

    state = {
        completed: 0,
    };

    render() {
        const { classes } = this.props;
        return (
            <div>
                <CircularProgress size={60} variant="static" value={this.props.percent} />

                <Button className={classes.button} variant="text" color="primary">
                    {precise(this.props.percent)}% / {web3.utils.fromWei(this.props.betPool)} ETH
                </Button>
                -
                <Button className={classes.button}  variant="text" color="primary">
                    {this.props.bettorsLen} bettor(s)
                </Button>

            </div>
        );
    }
}

CircularStatic.propTypes = {
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(CircularStatic);
