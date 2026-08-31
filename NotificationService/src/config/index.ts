import dotenv from "dotenv";

dotenv.config();

type ServerConfig = {
  PORT: number;
  REDIS_PORT:number;
  REDIS_HOST:string;
};

export const serverconfig: ServerConfig = {
  PORT: Number(process.env.PORT) || 3001,
  REDIS_PORT: Number(process.env.REDIS_PORT) || 6379,
  REDIS_HOST: process.env.REDIS_HOST || "localhost",
};