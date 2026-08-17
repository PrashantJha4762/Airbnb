import express from 'express'
import pingrouter from './pingrouter';
import { hotelrouter } from './hotel.router';

const v1Router=express.Router();

v1Router.use('/ping',pingrouter);
v1Router.use('/hotels',hotelrouter)
export default v1Router;