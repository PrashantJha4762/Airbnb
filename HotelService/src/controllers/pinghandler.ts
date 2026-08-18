import type { NextFunction, Request, Response } from "express"
import { logger } from "../config/logger.config";
import { StatusCodes } from "http-status-codes";
export const pinghandler=async(req:Request,res:Response,next:NextFunction)=>{
        logger.info("Ping request received");
        res.status(StatusCodes.OK).json({
                message:"pong",
        })
}