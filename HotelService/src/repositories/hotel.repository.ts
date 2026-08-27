import { logger } from "../config/logger.config";
import hotel from "../db/models";
import { createhoteldto } from "../dto/hotel.dto";
import { NotFoundError } from "../utils/error/app.error";

export async function CreateHotel(hotelData:createhoteldto){
    const Hotel=await hotel.create(hotelData);
    return Hotel;
}
export async function GetAllHotels(){
    const Hotels=await hotel.findAll({
        where:{
            deleted_at:null
        }
    })
    return Hotels;
}
export async function GetHotelById(id:number){
    const Hotel=await hotel.findByPk(id);
    return Hotel;
}
export async function SoftDelete(id:number){
    const Hotel=await hotel.findByPk(id);
    if(!Hotel){
        logger.error("Hotel not found")
        throw new NotFoundError(`Hotel with ${id} not found`);
    }
    Hotel.deleted_at= new Date();
    Hotel.save();
    return true;
}