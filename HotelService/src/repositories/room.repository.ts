import Rooms from "../db/models/roomodel";
import BaseRepsoitory from "./base.repository";
import type { CreationAttributes } from "sequelize";

export class Room extends BaseRepsoitory<Rooms>{
    constructor(){
        super(Rooms)
    }   
    async findByRoomCategoryIdAndDate(
        roomCategoryId: number,
        currentDate: Date
    ) {
        return await this.model.findOne({
            where: {
                room_category_id: roomCategoryId,
                DateOFAVAIlibility: currentDate,
                deleted_at: null
            }
        })
    }

    async bulkCreate(rooms: CreationAttributes<Rooms>[]) {
        return await this.model.bulkCreate(rooms);
    }

    async findLatestDateByRoomCategoryId(roomCategoryId: number): Promise<Date | null> {
        const result = await this.model.findOne({
            where: {
                room_category_id: roomCategoryId,
                deleted_at: null
            },
            attributes: ['DateOFAVAIlibility'],
            order: [['DateOFAVAIlibility', 'DESC']]
        });
        
        return result ? result.DateOFAVAIlibility : null;
    }

    async findLatestDatesForAllCategories(): Promise<Array<{roomCategoryId: number, latestDate: Date}>> {
        const results = await this.model.findAll({
            where: {
                deleted_at: null
            },
            attributes: [
                'room_category_id',
                [this.model.sequelize!.fn('MAX', this.model.sequelize!.col('DateOFAVAIlibility')), 'latestDate']
            ],
            group: ['room_category_id'],
            raw: true
        });
        
        return results.map((result: any) => ({
            roomCategoryId: result.room_category_id,
            latestDate: new Date(result.latestDate)
        }));
    }    
}    
