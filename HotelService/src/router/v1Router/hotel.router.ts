import express from 'express'
import { CreateHotelHandler, GetAllHotelsHandler, GetHoetlByidhandler } from '../../controllers/hotel.controller';
import { validateRequestBody } from '../../validators';
import { hotelschema } from '../../validators/hotel.validator';
export const hotelrouter=express.Router();
hotelrouter.post('/',validateRequestBody(hotelschema),CreateHotelHandler)
hotelrouter.get('/all',GetAllHotelsHandler)
hotelrouter.get('/:id',GetHoetlByidhandler)
