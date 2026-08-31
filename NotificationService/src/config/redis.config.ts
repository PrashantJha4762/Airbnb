import Redis from "ioredis";
import { serverconfig } from ".";


//iss pattern ko singleton pattern bolte h mtlb ki agr coonection pehli baar create ho rha h tb to
//wo if block m jaake create hoga. agr nhi ho rha h to iska mtlb h ki pehle wala connection object
//hi return hoga aur hm whi chahte h ki both consumer aur producer same redis connection se communicate kre
 function ConnectToRedis(){
    try{
        const redisconfig={
            host:serverconfig.REDIS_HOST,
            port:serverconfig.REDIS_PORT,
            maxRetriesPerRequest:null
        }
        let connection:Redis;
        return ()=>{
            if(!connection){
                connection=new Redis(redisconfig);
                return connection
            }
            return connection
        }
    }
    catch(err){
        console.log("Not able to connect to redis");
        throw err;
    }
}
export const GetRedisConnection=ConnectToRedis();