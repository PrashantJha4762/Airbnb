import dotenv from "dotenv";

dotenv.config();

type ServerConfig = {
  PORT: number;
  RedisUrl: string;
  TTL: number;
};
type DBconfig={
  username:string,
  password:string,
  database:string,
  host:string,
}

export const serverconfig: ServerConfig = {
  PORT: Number(process.env.PORT) || 3001,
  // Prefer the conventional uppercase name, while accepting the existing
  // RedisUrl entry in local .env files.
  RedisUrl: process.env.REDIS_URL || process.env.RedisUrl || "redis://localhost:6379",
  TTL: Number(process.env.TTL) || 6000
};
export const dbconfig:DBconfig={
  username:process.env.DB_USER||'root',
  password:process.env.DB_PWD||'areyoumad4762',
  database:process.env.DB_NAME||'airbnb_booking_service',
  host:process.env.DB_HOST||'127.0.0.1',
}
