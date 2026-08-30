import { serverconfig } from "../config";
import { sequelize } from "../db/models/sequelize";
import type { Bookingdto } from "../dto/booking.dto";
import { ConfirmBooking, createBooking, CreateIdempotencyKey, FinalizeBooking } from "../repositories/booking.repo";
import { GenerateIdempotencyKey } from "../utils/idemkey";
import { redlock } from "../config/redis.config";
import { InternalServerError } from "../utils/error/app.error";


//Hm chahte h ki concurrency as early as possible handle ho jaae mtlb ki agr ek user pehle se create booking request bhej chuka h to dusra request uske liye block ho jaae aur ye kaam hmne idempotency key ke through kiya h.
export async function CreateBookingService(bookingdata:Bookingdto){
    const bookingResource=`hotel:${bookingdata.hotelId}`;

    try {
        // Intentionally do not release this lock: it acts as a short-lived
        // hotel reservation and Redis releases it when the TTL expires.
        await redlock.acquire([bookingResource], serverconfig.TTL);

        const booking = await createBooking(bookingdata);
        const idempotencykey = await GenerateIdempotencyKey();

        await CreateIdempotencyKey(idempotencykey, booking.id);

        return { booking, idempotencykey };
    } catch {
        throw new InternalServerError("Failed to acquire lock for booking resource");
    }
}

//ye isme tx jo h wo ek transaction object h jo ki hme concurrent bookings request from same user ko handle krne me help krega.
//isme hmne pessimistic locking use kia h aur ye usi ka implementation h
export async function ConfirmBookingService(bookingid:number){
    return await sequelize.transaction(async (tx)=>{
        // const idempotencykeydata=GetIdempotencyKey(bookingid)
        const booking=await ConfirmBooking(bookingid,tx)
        await FinalizeBooking(bookingid,tx)
        return booking
    })
}
