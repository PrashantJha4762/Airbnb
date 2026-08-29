import { sequelize } from "../db/models/sequelize";
import type { Bookingdto } from "../dto/booking.dto";
import { ConfirmBooking, createBooking, CreateIdempotencyKey, FinalizeBooking, GetIdempotencyKey } from "../repositories/booking.repo";
import { GenerateIdempotencyKey } from "../utils/idemkey";


//Hm chahte h ki concurrency as early as possible handle ho jaae mtlb ki agr ek user pehle se create booking request bhej chuka h to dusra request uske liye block ho jaae aur ye kaam hmne idempotency key ke through kiya h.
export async function CreateBookingService(bookingdata:Bookingdto){
    const booking=await createBooking(bookingdata);
    const idempotencykey=await GenerateIdempotencyKey();
    await CreateIdempotencyKey(idempotencykey,booking.id);
    return {booking,idempotencykey}
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