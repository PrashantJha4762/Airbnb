import express from 'express';
import { serverconfig } from './config';
import v1Router from './router/v1Router/index.router';
import { GenericErrorHandler } from './middlewares/error.middleware';
import { logger } from './config/logger.config';
import { attachCorrelationId } from './middlewares/correlationId.middleware';
import { sequelize } from './db/models/sequelize.js';
import hotel from './db/models';
const app=express();
app.use(express.json());
app.use(attachCorrelationId);
app.use('/api/v1',v1Router);
app.use(GenericErrorHandler)
app.listen(serverconfig.PORT, async() => {
  await sequelize.authenticate()
  logger.info(`Database connected succesfully `)
  try{
    await hotel.create({name:"Hotel1",address:"Address1",rating:4.5,location:"Delhi"})
  } catch (error) {
    logger.error(`Error creating hotel: ${error}`);
  }
  logger.info(`Server is running at http://localhost:${serverconfig.PORT}`);
});