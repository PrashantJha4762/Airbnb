import express from 'express'
import { createBookingSchema } from '../validators/booking.validator';
import { createBookingHandler } from '../controllers/booking.controller';
import { validateRequestBody } from '../validators';

const bookingRouter = express.Router();

bookingRouter.post('/', validateRequestBody(createBookingSchema), createBookingHandler);
export default bookingRouter