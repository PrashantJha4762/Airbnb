import { createhoteldto } from "../dto/hotel.dto";
import { CreateHotel, GetAllHotels, GetHotelById, HotelRepository, SoftDelete } from "../repositories/hotel.repository"

export const hotelrepository=new HotelRepository();
export async function CreateHotelService(hotelData:createhoteldto){
    const Hotel=await hotelrepository.create(hotelData);
    return Hotel;
}  
export async function GetByIDService(id:number){
    const Hotel=await  hotelrepository.findById(id)
    return Hotel;
} 
export async function GetALlHotelService(){
    const Hotels=await hotelrepository.findall();
    return Hotels;
}
export async function SoftDeleteServices(id:number){
    const HOtel=await hotelrepository.softDelete(id);
    return HOtel;
}