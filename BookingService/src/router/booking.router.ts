import express from 'express'
import { createBookingSchema } from '../validators/booking.validator';
import { confirmBookingHandler, createBookingHandler } from '../controllers/booking.controller';
import { validateRequestBody } from '../validators';

const bookingRouter = express.Router();

bookingRouter.post('/', validateRequestBody(createBookingSchema), createBookingHandler);
bookingRouter.post('/:bookingId/confirm', confirmBookingHandler);
export default bookingRouter