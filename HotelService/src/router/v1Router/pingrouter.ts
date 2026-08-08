import express from 'express'
import { pinghandler } from '../../controllers/pinghandler';
import { validateRequestBody } from '../../validators';
import { pingSchema } from '../../validators/ping.validator';

const pingrouter=express.Router();

pingrouter.get('/',validateRequestBody(pingSchema),pinghandler)
export default pingrouter