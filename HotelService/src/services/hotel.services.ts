import { createhoteldto } from "../dto/hotel.dto";
import { CreateHotel, GetAllHotels, GetHotelById, SoftDelete } from "../repositories/hotel.repository"

export async function CreateHotelService(hotelData:createhoteldto){
    const Hotel=await CreateHotel(hotelData);
    return Hotel;
}  
export async function GetByIDService(id:number){
    const Hotel=await GetHotelById(id)
    return Hotel;
} 
export async function GetALlHotelService(){
    const Hotels=await GetAllHotels();
    return Hotels;
}
export async function SoftDeleteServices(id:number){
    const HOtel=await SoftDelete(id);
    return HOtel;
}