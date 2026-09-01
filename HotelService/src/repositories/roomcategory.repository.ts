import RoomCategory from "../db/models/roomCategory.model";
import BaseRepsoitory from "./base.repository";
export class RoomcategoryRepository extends BaseRepsoitory<RoomCategory>{
    constructor(){
        super(RoomCategory)
    }
}    