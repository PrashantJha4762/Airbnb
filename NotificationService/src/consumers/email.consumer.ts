import { Job, Worker } from "bullmq"
import type { NotificationDto } from "../dto/notification.dto"
import { MAILER_QUEUE } from "../queue/mailer.queue"
import { GetRedisConnection } from "../config/redis.config"

export const SetUpMailWorker=()=>{
        const emailConsumer= new Worker<NotificationDto>(
            MAILER_QUEUE,
            async (job:Job)=>{
                const payload=job.data;
                console.log(`Processing email for: ${JSON.stringify(payload)}`);
            },
            {
                connection:GetRedisConnection()
            }  
        )
            emailConsumer.on("failed", () => {
            console.error("Email processing failed");
        });

        emailConsumer.on("completed", () => {
            console.log("Email processing completed successfully");
        });
}