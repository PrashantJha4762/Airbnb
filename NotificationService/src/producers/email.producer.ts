import type { NotificationDto } from "../dto/notification.dto"
import { MAILER_QUEUE, mailerqueue } from "../queue/mailer.queue"

const MAILER_PAYLOAD="payload-mail"
export const AddEmailToQueue=async(payload:NotificationDto)=>{
    await mailerqueue.add(MAILER_PAYLOAD,payload)
    console.log(`Email added to queue: ${JSON.stringify(payload)}`);
}