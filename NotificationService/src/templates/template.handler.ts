import fs from 'fs/promises';
import Handlebars from 'handlebars';
import path from "path";
import { fileURLToPath } from 'node:url';
import { InternalServerError } from '../utils/error/app.error';

export async function renderEmailTemplate(templateId:string,params:Record<string,any>){
    const templateDirectory=path.join(path.dirname(fileURLToPath(import.meta.url)),'mailer');
    const templatepath=path.join(templateDirectory,`${templateId}.hbs`);
    try{
        const content=await fs.readFile(templatepath,'utf-8');
        const finalcontent=Handlebars.compile(content);
        return finalcontent(params);
    }
    catch(err){
        throw new InternalServerError(`Template not found ${templateId}`)
    }
}
