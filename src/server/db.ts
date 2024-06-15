import { drizzle } from "drizzle-orm/libsql";
import { createClient } from "@libsql/client";
import { serverEnv } from "~/env/server";

const client = createClient({
  url: serverEnv.TURSO_DATABASE_URL,
  authToken: serverEnv.TURSO_AUTH_TOKEN,
});

export const db = drizzle(client);
