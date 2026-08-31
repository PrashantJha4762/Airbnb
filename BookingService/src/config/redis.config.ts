import Redis from "ioredis";
import Redlock from "redlock";
import { serverconfig } from "./index.js";

function connectToRedis() {
    try {

        let connection: Redis;

        return () => {
            if (!connection) {
                connection = new Redis(serverconfig.RedisUrl);
                return connection;
            }

            return connection;
        }
        

    } catch (error) {
        console.error('Error connecting to Redis:', error);
        throw error;
    }
}

export const GetRedisConnection=connectToRedis();
export const redisClient = new Redis(serverconfig.RedisUrl, {
  // Redlock performs the retry logic itself. ioredis must not fail commands
  // after its own retry limit while a lock is being acquired or released.
  maxRetriesPerRequest: null,
});

redisClient.on("error", (error: Error) => {
  // ioredis emits connection failures as events. Handling them prevents the
  // noisy "Unhandled error event" message and preserves the original cause.
  console.error(`Redis connection error (${serverconfig.RedisUrl}):`, error.message);
});

export const redlock = new Redlock([redisClient], {
  driftFactor: 0.01,
  // A competing request for the same hotel should fail immediately instead
  // of waiting in the API request queue.
  retryCount: 0,
  retryDelay: 200,
  retryJitter: 200,
});

redlock.on("error", (error: Error) => {
  // A minority Redis-node failure can be expected in a distributed lock.
  // Log it for observability; lock operations still reject if quorum is lost.
  console.error("Redlock error:", error);
});
