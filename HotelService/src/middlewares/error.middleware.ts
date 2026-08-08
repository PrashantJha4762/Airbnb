import type { NextFunction, Request, Response } from "express";
import type { AppError } from "../utils/error/app.error";

export const GenericErrorHandler=(err:AppError,req:Request,res:Response,next:NextFunction)=>{
    res.status(err.StatusCode).json({
        success:false,
        message:err.message
    })
}