import type { Request, Response } from "express";
import {
  ConfirmBookingService,
  CreateBookingService,
} from "../services/booking.service.js";


export const createBookingHandler = async (req: Request, res: Response) => {
    const result = await CreateBookingService(req.body);

    return res.status(201).json({
        bookingId: result.booking.id,
        status: result.booking.status,
        idempotencyKey: result.idempotencykey,
    });
}

export const confirmBookingHandler = async (req: Request, res: Response) => {
    const bookingId = Number(req.params.bookingId);

    if (!Number.isSafeInteger(bookingId) || bookingId <= 0) {
        return res.status(400).json({
            message: "bookingId must be a positive integer",
        });
    }

    const booking = await ConfirmBookingService(bookingId);

    return res.status(200).json({
        bookingId: booking.id,
        status: booking.status,
    });
}
