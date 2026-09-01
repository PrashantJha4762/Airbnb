import { Model, ModelStatic, WhereOptions,CreationAttributes } from "sequelize";
abstract class BaseRepsoitory<T extends Model>{
    protected model:ModelStatic<T>
    constructor(model:ModelStatic<T>){
        this.model = model
    }
    async create(data:CreationAttributes<T>):Promise<T>{
        const created=await this.model.create(data)
        return created
    }
    async findall():Promise<T[]>{
        const data=await this.model.findAll()
        return data
    }
    async findById(id:number):Promise<T|null>{
        const data=await this.model.findByPk(id)
        return data
    }
    async delete(whereoptions:WhereOptions<T>):Promise<void>{
        const data=await this.model.destroy(
            {
                where:{
                    ...whereoptions
                }

            }
        )
    }
    async update(data:Partial<CreationAttributes<T>>,id:number):Promise<T|null>{
        const record=await this.model.findByPk(id)
        if(!record){
            return null
        }
        await record.update(data)
        await record.save()
        return record
    }
}
export default BaseRepsoitory