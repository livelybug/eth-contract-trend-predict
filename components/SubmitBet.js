import './utils/helper.css';
import {DisplayFormikState} from './utils/helper';
import React from 'react';
import {Formik} from 'formik';
import * as Yup from 'yup';
import Typography from "@material-ui/core/Typography/Typography";
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';

import predict from '../eth/predict';
import web3 from '../eth/web3';
import PoolProportion from './PoolProportion';
import { globalConfig } from '../src/utils/utils';

const blockGasLimit = 2900000;

const SubmitBet = (_props) => (
    <Card className="card">
        <CardContent>
            <Formik
                initialValues={{
                    wager: '',
                    _props
                }}

                onSubmit={async (values, {setSubmitting}) => {
                    const accounts = await web3.eth.getAccounts();

                    await predict.methods.bet(_props.range).send({
                        from: accounts[0],
                        value: web3.utils.toWei(values.wager.toString(), "ether"),
                        gas: blockGasLimit
                    });
                    setSubmitting(false);

                    if(globalConfig.debug)
                        setTimeout(() => {
                            alert(JSON.stringify(values, null, 2));
                        }, 500);
                }}

                validationSchema={Yup.object().shape({
                    wager: Yup.number()
                        .positive(),
                })}
            >
                {props => {
                    const {
                        values,
                        touched,
                        errors,
                        // dirty,
                        isSubmitting,
                        handleChange,
                        handleBlur,
                        handleSubmit
                        // handleReset
                    } = props;

                    const _predict = values._props.predict;

                    return (
                        <form onSubmit={handleSubmit}>
                            <label style={{display: 'block'}}>
                                {_predict.predict}
                            </label>

                            <PoolProportion percent={_predict.poolSum !== '0' ? (_predict.betPool / _predict.poolSum * 100) : 0 }
                                            betPool={_predict.betPool}
                                            bettorsLen={_predict.bettorsLen}
                            />

                            <Typography color="textSecondary" gutterBottom htmlFor="wager">
                                Bet Amount (ETH)
                            </Typography>

                            <input
                                id="wager"
                                placeholder="Place Your Bet (ETH)"
                                type="number"
                                value={values.wager}
                                onChange={handleChange}
                                onBlur={handleBlur}
                                className={
                                    errors.wager && touched.wager ? 'text-input error' : 'text-input'
                                }
                            />
                            {errors.wager &&
                            touched.wager && <div className="input-feedback">{errors.wager}</div>}

                            <button type="submit" disabled={isSubmitting}>
                                Bet !
                            </button>

                            { globalConfig.debug && <DisplayFormikState {...props} /> }

                        </form>
                    );
                }}
            </Formik>
        </CardContent>
    </Card>
);

export default SubmitBet;
