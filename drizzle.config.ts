import type { Config } from "drizzle-kit";
import { serverEnv } from "~/env/server";

export default {
  schema: "./drizzle/schema.ts",
  out: "./drizzle/migrations",
  dialect: "sqlite",
  driver: "turso",
  dbCredentials: {
    url: serverEnv.TURSO_DATABASE_URL!,
    authToken: serverEnv.TURSO_AUTH_TOKEN,
  },
} satisfies Config;
