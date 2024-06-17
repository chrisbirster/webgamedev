import { defineConfig } from "@solidjs/start/config";
import { prpcVite } from "@solid-mediakit/prpc-plugin";
import { authVite } from "@solid-mediakit/auth-plugin";
// @ts-ignore
import tailwindcss from "@tailwindcss/vite";

export default defineConfig({
  ssr: true,
  vite: {
    plugins: [
      tailwindcss(),
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
