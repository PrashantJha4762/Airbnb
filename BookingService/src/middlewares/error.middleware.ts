import type { NextFunction, Request, Response } from "express";
import { logger } from "../config/logger.config";

type HttpError = Error & {
    StatusCode?: number;
};

export const GenericErrorHandler=(err:HttpError,req:Request,res:Response,next:NextFunction)=>{
    const statusCode = err.StatusCode ?? 500;

    logger.error(err.message, { stack: err.stack });

    res.status(statusCode).json({
        success:false,
        message: statusCode === 500 ? "Internal server error" : err.message
    })
}
