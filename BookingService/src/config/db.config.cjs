const dotenv=require('dotenv');

dotenv.config();

const Config={
    development:{
        username:process.env.DB_USER,
        password:process.env.DB_PWD,
        database:process.env.DB_NAME,
        host:process.env.DB_HOST,
        dialect:'mysql'
    }
}
module.exports=Config