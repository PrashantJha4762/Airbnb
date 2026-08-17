import { NextFunction, Request, Response } from "express";
import { CreateHotelService, GetALlHotelService, GetByIDService } from "../services/hotel.services";

export async function CreateHotelHandler(req:Request,res:Response,next:NextFunction){
    const hotelresponse= await CreateHotelService(req.body);
    res.status(201).json({
        message:"Hotel created succesfully",
        data:hotelresponse,
        success:true
    })
}
export async function GetHoetlByidhandler(req:Request,res:Response,next:NextFunction){
    const hotelresponse= await GetByIDService(Number(req.params.id));
    res.status(200).json({
        message:"Hotel found succesfully",
        data:hotelresponse,
        success:true
    })
}
export async function GetAllHotelsHandler(req:Request,res:Response,next:NextFunction){
    const hotelresponse= await GetALlHotelService();
    res.status(200).json({
        message:"Hotels found succesfully",
        data:hotelresponse,
        success:true
    })
}