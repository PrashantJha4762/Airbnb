import z from "zod/v3";
export const hotelschema=z.object({
    name:z.string().min(1),
    address:z.string().min(1),
    location:z.string().min(1),
    rating:z.number().min(0).max(5)
})