import { Model, type CreationOptional, type InferAttributes, type InferCreationAttributes } from "sequelize";
import { sequelize } from "./sequelize";
enum BookingStatus{
  PENDING='PENDING',
  CONFIRMED='CONFIRMED',
  CANCELLED='CANCELLED'
}

class Bookings extends Model<InferAttributes<Bookings>,InferCreationAttributes<Bookings>>{
  declare id:CreationOptional<number>
  declare userid:number
  declare hotelid:number
  declare CreatedAt:CreationOptional<Date>
  declare UpdatedAt:CreationOptional<Date>
  declare bookingAmount:number
  declare status:BookingStatus
  declare totalguests:number
}

class IdempotencyKey extends Model<InferAttributes<IdempotencyKey>,InferCreationAttributes<IdempotencyKey>>{
  declare id:CreationOptional<number>
  declare idemkey:string
  declare Created_At:CreationOptional<Date>
  declare updated_At:CreationOptional<Date>
  declare finalized:CreationOptional<boolean>
  declare bookingId:number
}

Bookings.init({
    id:{
        type:"INTEGER",
        primaryKey:true,
        autoIncrement:true
    },
    userid:{
        type:"INTEGER",
        allowNull:false
    },
    hotelid:{
        type:"INTEGER",
        allowNull:false
    },
    CreatedAt:{
        type:"DATETIME",
        allowNull:false,
        defaultValue:()=>new Date()
    },
    UpdatedAt:{
        type:"DATETIME",
        allowNull:false,
        defaultValue:()=>new Date()
    },
    bookingAmount:{
        type:"INTEGER",
        allowNull:false
    },
    status:{
        type:"ENUM('PENDING', 'CONFIRMED', 'CANCELLED')",
        allowNull:false,
        defaultValue:'PENDING'
    },
    totalguests:{
      type:"INTEGER",
      allowNull:false
    }
},{
  sequelize,
  modelName:"Bookings"
})

IdempotencyKey.init({
  id:{
    type:"INTEGER",
    primaryKey:true,
    autoIncrement:true
  },
  idemkey:{
    type:"STRING",
    allowNull:false
  },
  Created_At:{
    type:"DATETIME",
    allowNull:false,
    defaultValue:()=>new Date()
  },
  updated_At:{
    type:"DATETIME",
    allowNull:false,
    defaultValue:()=>new Date()
  },
  finalized:{
    type:"BOOLEAN",
    allowNull:false,
    defaultValue:false
  },
  bookingId:{
    type:"INTEGER",
    allowNull:false,
    unique:true
  }
},{
  sequelize,
  modelName:"IdempotencyKey"
})

Bookings.hasOne(IdempotencyKey,{
  foreignKey:"bookingId",
  onDelete:"CASCADE"
})
IdempotencyKey.belongsTo(Bookings,{
  foreignKey:"bookingId"
})

export {Bookings,IdempotencyKey}
