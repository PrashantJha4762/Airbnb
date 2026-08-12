import dotenv from 'dotenv';
dotenv.config();
const config={
  "development":{
    "username":process.env.DB_USERNAME || 'root',
    "password":process.env.DB_PASSWORD || 'abc',
    "database":process.env.DB_NAME || 'database_development',
    "host":process.env.DB_HOST || '127.0.0.1',
    "dialect":"mysql"
  }  
}
export default config;