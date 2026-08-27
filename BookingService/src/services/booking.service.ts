import type { Bookingdto } from "../dto/booking.dto";
import { createBooking, CreateIdempotencyKey } from "../repositories/booking.repo";
import { GenerateIdempotencyKey } from "../utils/idemkey";

export async function CreateBookingService(bookingdata:Bookingdto){
    const booking=await createBooking(bookingdata);
    const idempotencykey=await GenerateIdempotencyKey();
    await CreateIdempotencyKey(idempotencykey,booking.id);
    return {booking,idempotencykey}
}
// export async function FinalizeBookingService(bookingid:number){
//     const booking=await createBooking(bookingid);
// }