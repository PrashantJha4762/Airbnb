import nodemailer from 'nodemailer';
import { serverconfig } from '.';

const transporter=nodemailer.createTransport({
    service:'gmail',
    auth:{
        user:serverconfig.MAILER_USER,
        pass:serverconfig.MAILER_PWD
    }
})
export default transporter