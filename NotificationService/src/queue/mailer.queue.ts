import { Queue } from "bullmq"
import { GetRedisConnection } from "../config/redis.config"

export const MAILER_QUEUE="queue-mailer"

export const mailerqueue=new Queue(MAILER_QUEUE,{
    connection:GetRedisConnection()
})