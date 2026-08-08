import {v4 as uuidv4} from "uuid";
import type { NextFunction, Request, Response } from "express";
import { asynclocalstorage } from "../utils/helpers/request.helpers";

//The below fn is attaching the correlation id to the request headers so that it can be used in the logger middleware to log the correlation id in the logs.
export const attachCorrelationId=(req:Request,res:Response,next:NextFunction)=>{
    const correlationId=uuidv4();
    asynclocalstorage.run({correlationId:correlationId},()=>{
        next();
    })
}