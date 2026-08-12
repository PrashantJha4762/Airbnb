import dotenv from "dotenv";

dotenv.config();

type ServerConfig = {
  PORT: number;
};

type DBCONFIG={
  username: string;
  password: string;
  database: string;
  host: string;
};

export const serverconfig: ServerConfig = {
  PORT: Number(process.env.PORT) || 3001,
};

export const dbconfig: DBCONFIG = {
  username: process.env.DB_USERNAME || 'root',
  password: process.env.DB_PASSWORD || 'abc',
  database: process.env.DB_NAME || 'database_development',
  host: process.env.DB_HOST || '127.0.0.1',
};