import { Model, type CreationOptional, type InferAttributes, type InferCreationAttributes } from "sequelize";
import { sequelize } from "./sequelize";
export enum BookingStatus{
  PENDING='PENDING',
  CONFIRMED='CONFIRMED',
  CANCELLED='CANCELLED'
}

class Bookings extends Model<InferAttributes<Bookings>,InferCreationAttributes<Bookings>>{
  declare id:CreationOptional<number>
  declare userId:number
  declare hotelId:number
  declare createdAt:CreationOptional<Date>
  declare updatedAt:CreationOptional<Date>
  declare bookingAmount:number
  declare status:BookingStatus
  declare totalGuests:number
}

class IdempotencyKey extends Model<InferAttributes<IdempotencyKey>,InferCreationAttributes<IdempotencyKey>>{
  declare id:CreationOptional<number>
  declare idemkey:string
  declare createdAt:CreationOptional<Date>
  declare updatedAt:CreationOptional<Date>
  declare finalized:CreationOptional<boolean>
  declare bookingId:number
}

Bookings.init({
    id:{
        type:"INTEGER",
        primaryKey:true,
        autoIncrement:true
    },
    userId:{
        type:"INTEGER",
        allowNull:false
    },
    hotelId:{
        type:"INTEGER",
        allowNull:false
    },
    createdAt:{
        type:"DATETIME",
        field:"CreatedAt",
        allowNull:false,
        defaultValue:()=>new Date()
    },
    updatedAt:{
        type:"DATETIME",
        field:"UpdatedAt",
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
    totalGuests:{
      type:"INTEGER",
      field:"totalguest",
      allowNull:false
    }
},{
  sequelize,
  modelName:"Bookings",
  tableName:"Bookings",
  timestamps:false
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
  createdAt:{
    type:"DATETIME",
    field:"Created_At",
    allowNull:false,
    defaultValue:()=>new Date()
  },
  updatedAt:{
    type:"DATETIME",
    field:"Updated_At",
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
  modelName:"IdempotencyKey",
  tableName:"Idempotency",
  timestamps:false
})

Bookings.hasOne(IdempotencyKey,{
  foreignKey:"bookingId",
  onDelete:"CASCADE"
})
IdempotencyKey.belongsTo(Bookings,{
  foreignKey:"bookingId"
})

export {Bookings,IdempotencyKey}
