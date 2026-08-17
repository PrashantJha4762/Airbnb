import hotel from "../db/models";
import { createhoteldto } from "../dto/hotel.dto";

export async function CreateHotel(hotelData:createhoteldto){
    const Hotel=await hotel.create(hotelData);
    return Hotel;
}
export async function GetAllHotels(){
    const Hotels=await hotel.findAll()
    return Hotels;
}
export async function GetHotelById(id:number){
    const Hotel=await hotel.findByPk(id);
    return Hotel;
}