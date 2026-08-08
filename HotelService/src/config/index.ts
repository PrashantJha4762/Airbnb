import dotenv from "dotenv";

dotenv.config();

type ServerConfig = {
  PORT: number;
};

export const serverconfig: ServerConfig = {
  PORT: Number(process.env.PORT) || 3001,
};