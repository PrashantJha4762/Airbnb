import z from "zod/v3"

export const pingSchema=z.object({
    message:z.string().min(1)
})

//Setting up zod schema for validating the request body and query parameters for the ping endpoint. The schema expects an object with a single property "message" which should be a non-empty string.

// 1.creating a zod schema for validating the request body and query parameters for the ping endpoint. The schema expects an object with a single property "message" which should be a non-empty string.
// 2. The schema is exported as pingSchema so that it can be used in the ping handler to validate the request body and query parameters before processing the request.
// 3.in other files we will use this schema to validate the request body and query parameters for the ping endpoint. If the validation fails, we will return a 400 Bad Request error with a message indicating that the request is invalid.
// 4.for validating we use parseasync method of zod schema which returns a promise that resolves if the validation is successful and rejects with an error if the validation fails and it happens on the go