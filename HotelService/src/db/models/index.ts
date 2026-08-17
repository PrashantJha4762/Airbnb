import { DataTypes, Model, type CreationOptional, type InferAttributes, type InferCreationAttributes } from "sequelize";
import { sequelize } from "./sequelize.js";

class hotel extends Model<InferAttributes<hotel>,InferCreationAttributes<hotel>>{
        declare id:CreationOptional<number>
        declare name:string
        declare address:string
        declare location:CreationOptional<string>
        declare rating:CreationOptional<number>
        declare createdAt:CreationOptional<Date>
        declare updatedAt:CreationOptional<Date>
}
hotel.init({
    id:{
        type:"INTEGER",
        autoIncrement:true,
        primaryKey:true
    },
    name:{
        type:"STRING",
        allowNull:false 
    },
    address:{
        type:"STRING",
        allowNull:false
    },
    location:{
        type:"STRING",
        allowNull:true
    },
    rating:{
        type:"DECIMAL(3,2)",
        defaultValue:0
    },
    createdAt:{
        type:"DATE",
        defaultValue:DataTypes.NOW
    },
    updatedAt:{
        type:"DATE",
        defaultValue:DataTypes.NOW
    }    
},{
    sequelize,
    modelName:"hotel",
    tableName:"hotel"        
})
export default hotel