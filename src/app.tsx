// @refresh reload
import { MetaProvider, Title } from "@solidjs/meta";
import { Router } from "@solidjs/router";
import { FileRoutes } from "@solidjs/start/router";
import { Suspense } from "solid-js";
import { QueryClient } from "@tanstack/solid-query";
import { SessionProvider } from "@solid-mediakit/auth/client";
import { PRPCProvider } from "@solid-mediakit/prpc/provider";

import "./app.css";

export default function App() {
  const queryClient = new QueryClient();
  return (
    <Router
      root={(props) => (
        <MetaProvider>
          <Title>WebGameDev</Title>
          <Suspense>
            <SessionProvider>
              <PRPCProvider queryClient={queryClient}>
                {props.children}
              </PRPCProvider>
            </SessionProvider>
          </Suspense>
        </MetaProvider>
      )}
    >
      <FileRoutes />
    </Router>
  );
}
