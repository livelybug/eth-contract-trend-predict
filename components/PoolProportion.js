import React from 'react';
import PropTypes from 'prop-types';
import {withStyles} from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import web3 from '../eth/web3';
import {precise} from '../src/utils/utils';
import CircularProgressbar from 'react-circular-progressbar';
import 'react-circular-progressbar/dist/styles.css';

const styles = theme => ({
    button: {
        margin: theme.spacing.unit * 2
    },
    progressBar: {
        width: '100px'
    }
});

class CircularStatic extends React.Component {

    constructor(props) {
        super(props);
    }

    render() {
        const { classes } = this.props;
        return (
            <div>
                <CircularProgressbar percentage={precise(this.props.percent)}
                                     text={`${precise(this.props.percent)}%`} className={classes.progressBar}/>

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
