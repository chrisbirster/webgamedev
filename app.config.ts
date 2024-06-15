import { defineConfig } from "@solidjs/start/config";
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
    ],
  },
  server: {
    preset: "vercel",
  },
});
