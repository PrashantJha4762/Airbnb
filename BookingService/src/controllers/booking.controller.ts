import type { Request, Response } from "express";
import { CreateBookingService } from "../services/booking.service";


export const createBookingHandler = async (req: Request, res: Response) => {

    const booking = await CreateBookingService(req.body);

    res.status(200).json({
        bookingId: booking.booking.id,
        status: booking.booking.status,
        idempotencyKey: booking.idempotencykey
    });
}

// export const confirmBookingHandler = async (req: Request, res: Response) => {
//     const booking = await confirm(req.params.idempotencyKey);

//     res.status(200).json({
//         bookingId: booking.id,
//         status: booking.status,
//     });
// }