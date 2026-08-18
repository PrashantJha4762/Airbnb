import { NextFunction, Request, Response } from "express";
import { CreateHotelService, GetALlHotelService, GetByIDService } from "../services/hotel.services";
import { StatusCodes } from "http-status-codes";

export async function CreateHotelHandler(req:Request,res:Response,next:NextFunction){
    const hotelresponse= await CreateHotelService(req.body);
    res.status(StatusCodes.CREATED).json({
        message:"Hotel created succesfully",
        data:hotelresponse,
        success:true
    })
}
export async function GetHoetlByidhandler(req:Request,res:Response,next:NextFunction){
    const hotelresponse= await GetByIDService(Number(req.params.id));
    res.status(StatusCodes.OK).json({
        message:"Hotel found succesfully",
        data:hotelresponse,
        success:true
    })
}
export async function GetAllHotelsHandler(req:Request,res:Response,next:NextFunction){
    const hotelresponse= await GetALlHotelService();
    res.status(StatusCodes.OK).json({
        message:"Hotels found succesfully",
        data:hotelresponse,
        success:true
    })
}