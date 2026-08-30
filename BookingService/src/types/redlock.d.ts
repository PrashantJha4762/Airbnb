/**
 * Temporary typing shim for redlock@5.0.0-beta.2.
 * The package ships declarations, but omits the `types` condition from its
 * package `exports`, so TypeScript cannot resolve them in ESM/bundler mode.
 */
declare module "redlock" {
  type Settings = {
    driftFactor?: number;
    retryCount?: number;
    retryDelay?: number;
    retryJitter?: number;
    automaticExtensionThreshold?: number;
  };

  export class ExecutionError extends Error {}

  export interface Lock {
    release(): Promise<unknown>;
  }

  export type RedlockAbortSignal = AbortSignal & {
    error?: Error;
  };

  export default class Redlock {
    constructor(clients: Iterable<unknown>, settings?: Settings);
    on(event: "error", listener: (error: Error) => void): this;
    acquire(resources: string[], duration: number): Promise<Lock>;
    using<T>(
      resources: string[],
      duration: number,
      routine: (signal: RedlockAbortSignal) => Promise<T>,
    ): Promise<T>;
  }
}
