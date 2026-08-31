import dotenv from "dotenv";

dotenv.config();

type ServerConfig = {
  PORT: number;
  REDIS_PORT:number;
  REDIS_HOST:string;
  MAILER_USER:string
  MAILER_PWD:string
};

export const serverconfig: ServerConfig = {
  PORT: Number(process.env.PORT) || 3001,
  REDIS_PORT: Number(process.env.REDIS_PORT) || 6379,
  REDIS_HOST: process.env.REDIS_HOST || "localhost",
  MAILER_USER: process.env.MAILER_USER || "",
  MAILER_PWD: process.env.MAILER_PWD || ""
};