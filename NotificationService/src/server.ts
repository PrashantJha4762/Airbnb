import express from 'express';
import { serverconfig } from './config';
import v1Router from './router/v1Router/index.router';
import { GenericErrorHandler } from './middlewares/error.middleware';
import { logger } from './config/logger.config';
import { attachCorrelationId } from './middlewares/correlationId.middleware';
import { SetUpMailWorker } from './consumers/email.consumer';
import { AddEmailToQueue } from './producers/email.producer';
const app=express();
app.use(express.json());
app.use(attachCorrelationId);
app.use('/api/v1',v1Router);
app.use(GenericErrorHandler)
app.listen(serverconfig.PORT, () => {
  logger.info(`Server is running at http://localhost:${serverconfig.PORT}` );
  SetUpMailWorker()
  const samplenotification={
    to:"user@example.com",
    subject:"Welcome to our service",
    template_id:"welcome_template",
    params:{
        name:"John Doe",
        signupDate:"2023-10-01"
  }
  };
  AddEmailToQueue(samplenotification)
});