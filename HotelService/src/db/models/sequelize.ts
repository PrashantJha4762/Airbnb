import { Sequelize } from "sequelize";
import { dbconfig } from "../../config/index.js";

export const sequelize=new Sequelize({
    username:dbconfig.username,
    password:dbconfig.password,
    database:dbconfig.database,
    dialect:"mysql",
    host:dbconfig.host
})