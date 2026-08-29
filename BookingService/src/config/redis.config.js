import Redis from "ioredis";
import Redlock from 'redlock';
import { serverconfig } from ".";
export const redisclient=new Redis(serverconfig.RedisUrl);

export const redlock=new Redlock([redisclient],{
    driftFactor:0.01,
    retryCount:10,
    retryDelay:200,
    retryJitter:200
})  
