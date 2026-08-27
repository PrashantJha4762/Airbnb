import type { NextFunction, Request, Response } from "express";
import type { AnyZodObject } from "zod/v3";
import { logger } from "../config/logger.config";

export const validateRequestBody=(schema:AnyZodObject)=>{
    return async (req:Request,res:Response,next:NextFunction)=>{
        try{
            logger.info("Validating request body",);
            await schema.parseAsync(req.body)
            logger.info("Request is valid");
            next();
        }
        catch(err){
             res.status(400).json({
                message:"Invalid Request Body",
                success:false,
                error:err
            })
        }
    }
}

export const validateRequestQuery=(schema:AnyZodObject)=>{
    return async (req:Request,res:Response,next:NextFunction)=>{
        try{
            logger.info("Validating request query");
            await schema.parseAsync(req.query)
            logger.info("Query is valid");
            next();
        }
        catch(err){
             res.status(400).json({
                message:"Invalid Request Query",
                success:false,
                error:err
            })
        }
    }
}