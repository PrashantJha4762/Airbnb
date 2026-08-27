import { Sequelize } from "sequelize";
import { dbconfig } from "../../config";
export const sequelize=new Sequelize({
    username:dbconfig.username,
    password:dbconfig.password,
    host:dbconfig.host,
    dialect:"mysql",
    database:dbconfig.database
})