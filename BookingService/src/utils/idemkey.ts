import { v4 as uuidv4 } from 'uuid';
export async function GenerateIdempotencyKey(){
    return uuidv4();
}