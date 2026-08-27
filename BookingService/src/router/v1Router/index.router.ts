import express from 'express'
import bookingRouter from '../booking.router';
import pingrouter from './pingrouter';

const v1Router=express.Router();

v1Router.use('/ping',pingrouter);
v1Router.use('/bookings', bookingRouter);

export default v1Router;
