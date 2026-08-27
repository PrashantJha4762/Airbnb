import { Bookings, BookingStatus, IdempotencyKey } from "../db/models";
import type { Bookingdto } from "../dto/booking.dto";
import { GenerateIdempotencyKey } from "../utils/idemkey";

export async function createBooking(bookingdata:Bookingdto){
    const booking=await Bookings.create({
        ...bookingdata,
        status:BookingStatus.PENDING
    })
    return booking
}
export async function CreateIdempotencyKey(key:string,bookingid:number){
    const idempotencykey=await IdempotencyKey.create({
        idemkey:key,
        bookingId:bookingid
    })
    return idempotencykey
}
export async function GetIdempotencyKey(bookingid:number){
    const idempotencykey=await IdempotencyKey.findOne({
        where:{
            bookingId:bookingid 
        }
    })     
    return idempotencykey       
}
export async function GetBookinByid(bookingid:number){
    const booking=await Bookings.findByPk(bookingid)
    return booking
}
export async function CancelBooking(bookingid:number){
    const booking=await Bookings.findByPk(bookingid);
    if(!booking){
        throw new Error("Booking not found")
    }
    booking.status=BookingStatus.CANCELLED;
    await booking.save();
    return booking
}
export async function ConfirmBooking(bookingid:number){
    const booking=await Bookings.findByPk(bookingid);
    if(!booking){
        throw new Error("Booking not found")
    }
    booking.status=BookingStatus.CONFIRMED;
    await booking.save();
    return booking
}