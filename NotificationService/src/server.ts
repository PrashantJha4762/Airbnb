import express from 'express';
import { serverconfig } from './config';
import v1Router from './router/v1Router/index.router';
import { GenericErrorHandler } from './middlewares/error.middleware';
import { logger } from './config/logger.config';
import { attachCorrelationId } from './middlewares/correlationId.middleware';
const app=express();
app.use(express.json());
app.use(attachCorrelationId);
app.use('/api/v1',v1Router);
app.use(GenericErrorHandler)
app.listen(serverconfig.PORT, () => {
  logger.info(`Server is running at http://localhost:${serverconfig.PORT}` );
});