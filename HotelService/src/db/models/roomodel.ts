import { InferAttributes, InferCreationAttributes, Model } from "sequelize";
import { sequelize } from "./sequelize.js";
class Rooms extends Model<InferAttributes<Rooms>,InferCreationAttributes<Rooms>>{
    declare id:number
    declare room_category_id:number
    declare hotel_id:number
    declare room_no:string
    declare DateOFAVAIlibility:Date
    declare booking_id:number
    declare created_at:Date
    declare updated_at:Date
    declare deleted_at:Date|null
}
Rooms.init({
    id:{
        type:"INTEGER", 
        autoIncrement:true,
        primaryKey:true
    },  
    room_category_id:{
        type:"INTEGER",
        allowNull:false
    },
    hotel_id:{
        type:"INTEGER",
        allowNull:false
    },
    room_no:{
        type:"STRING",
        allowNull:false
    },
    DateOFAVAIlibility:{
        type:"DATE",
        allowNull:false
    },
    booking_id:{
        type:"INTEGER",
        allowNull:true
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
    modelName:'Rooms'
});
export default Rooms