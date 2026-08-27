import { Bookings, BookingStatus } from "../db/models";
import type { Bookingdto } from "../dto/booking.dto";
import { GenerateIdempotencyKey } from "../utils/idemkey";

export async function createBooking(bookingdata:Bookingdto){
    const booking=await Bookings.create({
        ...bookingdata,
        status:BookingStatus.PENDING
    })
    return booking
}
export async function CreateIdempotencyKey(){
    const idemkey= await GenerateIdempotencyKey()

}
