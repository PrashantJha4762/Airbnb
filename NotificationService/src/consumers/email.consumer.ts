import { Job, Worker } from "bullmq"
import type { NotificationDto } from "../dto/notification.dto"
import { MAILER_QUEUE } from "../queue/mailer.queue"
import { GetRedisConnection } from "../config/redis.config"
import { renderEmailTemplate } from "../templates/template.handler"
import { sendEMail } from "../services/mailer.service"
import { logger } from "../config/logger.config"

export const SetUpMailWorker=()=>{
        const emailConsumer= new Worker<NotificationDto>(
            MAILER_QUEUE,
            async (job:Job<NotificationDto>)=>{
                const payload=job.data;
                console.log(`Processing email for: ${JSON.stringify(payload)}`);
                const emailcontent=await renderEmailTemplate(payload.template_id,payload.params)
                await sendEMail(payload.to,payload.subject,emailcontent)
                logger.info(`Email sent to ${payload.to} with subject "${payload.subject}"`);
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
