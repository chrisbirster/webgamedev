import { type SolidAuthConfig } from "@solid-mediakit/auth";
import Github from "@auth/core/providers/github";
import { serverEnv } from "~/env/server";
import { DrizzleAdapter } from "@auth/drizzle-adapter";
import { db } from "~/server/db";

declare module "@auth/core/types" {
  export interface Session {
    user?: DefaultSession["user"];
  }
}

export const authOptions: SolidAuthConfig = {
  adapter: DrizzleAdapter(db),
  providers: [
    Github({
      clientId: serverEnv.TEST_GITHUB_CLIENT_ID,
      clientSecret: serverEnv.TEST_GITHUB_CLIENT_SECRET,
    }),
  ],
  debug: false,
};
