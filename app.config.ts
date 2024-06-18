import { defineConfig } from "@solidjs/start/config";
import { prpcVite } from "@solid-mediakit/prpc-plugin";
import { authVite } from "@solid-mediakit/auth-plugin";

export default defineConfig({
  ssr: true,
  vite: {
    plugins: [
      authVite({
        authOpts: {
          name: "authOptions",
          dir: "~/server/auth",
        },
        redirectTo: "/",
      }),
      prpcVite({ log: false }),
    ],
  },
  server: {
    preset: "vercel",
  },
});
