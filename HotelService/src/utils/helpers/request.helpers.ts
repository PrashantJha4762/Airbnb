import { AsyncLocalStorage } from "async_hooks"

type AsyncLocalStoragetype={
    correlationId:string
}
export const asynclocalstorage=new AsyncLocalStorage<AsyncLocalStoragetype>() //it creates a new instance of the AsyncLocalStorage class, which is used to store and retrieve data that is specific to a particular asynchronous context. In this case, we are using it to store the correlation ID for each request.

// The below fn is used to set the correlation id in the async local storage so that it can be used in the logger middleware to log the correlation id in the logs.

export const getcorrelationId=()=>{
    const asyncstore=asynclocalstorage.getStore()
    return asyncstore?.correlationId||"unknown error while fetching correlation id" //it retrieves the correlation ID from the async local storage. If the correlation ID is not found, it returns a default message indicating an error while fetching the correlation ID.
}

//The previous approach of just attaching the correlation id to the request headers was not working as expected because for the background tasks, asynchronous network calls, and other asynchronous operations, the correlation ID was not being propagated correctly. This is because the request headers are not automatically passed along to these asynchronous contexts. As a result, the correlation ID was not available in the logger middleware for these cases, leading to missing or incorrect correlation IDs in the logs.
//That's why we have had to use this approach of using async local storage to store the correlation ID for each request. This way, the correlation ID is available in the logger middleware for all asynchronous contexts, ensuring that the logs are correctly correlated with the corresponding requests.
//This asynclocalstorage is the concept of thread local storage in nodejs. It allows us to store data that is specific to a particular asynchronous context, such as a request or a background task. This is useful for propagating context information, such as correlation IDs, across different parts of the application, even when they are executed in different asynchronous contexts.