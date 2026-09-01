import { InferAttributes, InferCreationAttributes, Model } from "sequelize";
import { sequelize } from "./sequelize.js";
class RoomCategory extends Model<InferAttributes<RoomCategory>,InferCreationAttributes<RoomCategory>>{
    declare id:number
    declare hotel_id:number
    declare price:number
    declare room_type:string
    declare room_count:number
    declare created_at:Date
    declare updated_at:Date
    declare deleted_at:Date|null
}
RoomCategory.init({
    id:{
        type:"INTEGER", 
        autoIncrement:true,
        primaryKey:true
    },  
    hotel_id:{
        type:"INTEGER",
        allowNull:false
    },
    price:{
        type:"DECIMAL(10,2)",
        allowNull:false
    },
    room_type:{
        type:"ENUM('single', 'double', 'suite')",
        allowNull:false
    },
    room_count:{
        type:"INTEGER",
        allowNull:false
    },  
    created_at:{
        type:"DATE",
        allowNull:false
    },
    updated_at:{
        type:"DATE",
        allowNull:false
    },
    deleted_at:{
        type:"DATE",
        allowNull:true
    }
}, {
    sequelize,
    modelName:'RoomCategory'
});
export default RoomCategory    