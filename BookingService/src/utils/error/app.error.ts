export interface AppError extends Error{
    StatusCode:number
}
export class InternalServerError implements AppError{
    name: string;
    message: string;
    StatusCode: number;
    constructor(message:string){
        this.name="InternalServerError";
        this.message=message;
        this.StatusCode=500;
    }
}
export class BadRequestError implements AppError{
    name: string;
    message: string;
    StatusCode: number;
    constructor(message:string){
        this.name="BadRequestError";
        this.message=message;
        this.StatusCode=400;
    }
}
export class NotFoundError implements AppError{
    name: string;
    message: string;
    StatusCode: number;
    constructor(message:string){
        this.name="NotFoundError";
        this.message=message;
        this.StatusCode=404;
    }
}
export class UnauthorizedError implements AppError{
    name: string;
    message: string;
    StatusCode: number;
    constructor(message:string){
        this.name="UnauthorizedError";
        this.message=message;
        this.StatusCode=401;
    }
}
export class ForbiddenError implements AppError{
    name: string;
    message: string;
    StatusCode: number;
    constructor(message:string){
        this.name="ForbiddenError";
        this.message=message;
        this.StatusCode=403;
    }       
}
export class ConflictError implements AppError{
    name: string;
    message: string;
    StatusCode: number;
    constructor(message:string){
        this.name="ConflictError";
        this.message=message;
        this.StatusCode=409;
    }   
}    