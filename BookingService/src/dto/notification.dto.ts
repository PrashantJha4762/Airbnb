export interface NotificationDto{
    to:string,
    subject:string,
    template_id:string,
    params:Record<string,any>
}