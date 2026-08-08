import winston from "winston";
import { getcorrelationId } from "../utils/helpers/request.helpers";
import DailyRotateFile from "winston-daily-rotate-file";

export const logger=winston.createLogger({
    format:winston.format.combine(
        winston.format.timestamp({format:"YYYY-MM-DD HH:mm:ss"}), //it adds timestamp to the log

        winston.format.json(), //it converts the log into json format

        winston.format.printf(({level,message,timestamp,...data})=>{  //this printf thing ensures how the logs should be printed 
            const output={level,message,timestamp,correlationId:getcorrelationId(),data};
            return JSON.stringify(output)
        })

    ),//it combines more than one thing which we require for logging

    transports:[
        new winston.transports.Console(),
        new DailyRotateFile({
            filename:"logs/%DATE%-app.log", //it will create a new log file for each day and store it in logs folder
            datePattern:"YYYY-MM-DD", //it will create a new log file for each day
            maxSize:"20m", //it will create a new log file if the size of the log file exceeds 20mb
            maxFiles:"14d" //it will keep the log files for 14 days
        })
    ]

}) //this fn is used to create the logger obj



//STEPS OF SETTING UP PRODUCTION GRADE LOGGER
// 1. Create a logger object using winston.createLogger() and configure it with desired formats and transports.
// 2. Use the logger object to log messages at different levels (info, error, etc.) throughout your application.
// 3. Configure the logger to output logs in a structured format (e.g., JSON) for better analysis and monitoring.
// 4. Set up transports to determine where the logs should be sent (e.g., console, file, external logging service).
// 5. Ensure that sensitive information is not logged and that logs are properly managed and rotated in production environments.   


//There is a problem that if 2 requests come at the same time then the logs will be mixed up and it will be hard to debug. To solve this problem we will first generate a unique request id for each request and then we will use that request id in the logs so that we can easily identify which log belongs to which request. This will be done in the middleware.
//But there is a problem that if we generate a unique request id in the middleware and then use that request id in the logger then it will not work for asynchronous operations like database calls, network calls etc. because the request id will be lost in the asynchronous context. To solve this problem we will use async local storage which is a concept of thread local storage in nodejs. It allows us to store data that is specific to a particular asynchronous context, such as a request or a background task. This is useful for propagating context information, such as correlation IDs, across different parts of the application, even when they are executed in different asynchronous contexts.
//We basically create a new instance of the AsyncLocalStorage class, which is used to store and retrieve data that is specific to a particular asynchronous context. In this case, we are using it to store the correlation ID for each request. The correlation ID is generated in the middleware and then stored in the async local storage. The logger then retrieves the correlation ID from the async local storage and uses it in the logs. This way, the correlation ID is available in the logger middleware for all asynchronous contexts, ensuring that the logs are correctly correlated with the corresponding requests.