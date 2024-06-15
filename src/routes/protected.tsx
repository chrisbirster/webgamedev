import { createSession, signOut } from "@solid-mediakit/auth/client";
import { Show, For, type VoidComponent } from "solid-js";

import { createAsync, type RouteDefinition } from "@solidjs/router";
import { cache } from "@solidjs/router";

import { db } from "~/server/db";
import { fooTable } from "../../drizzle/schema";

const getFoo = cache(async function getFoo() {
  "use server";
  const response = await db.select().from(fooTable).execute();
  console.log("response");
  console.log(response);
  return response;
}, "foo-data");

export const route = {
  load: () => getFoo(),
} satisfies RouteDefinition;

const Protected: VoidComponent = () => {
  const session = createSession();
  const foos = createAsync(() => getFoo());

  return (
    <Show
      when={session()}
      fallback={<p>Only shown for logged in users</p>}
      keyed
    >
      {(us) => (
        <main>
          <h1>Protected</h1>
          {JSON.stringify(us)}
          <For each={foos()} fallback={<p>Loading...</p>}>
            {(foo) => <div>{foo.bar}</div>}
          </For>
          {us.user?.image ? <img src={us.user?.image} /> : null}
          <span>Hey there {us.user?.name}! You are signed in!</span>
          <button
            onClick={() =>
              void signOut({
                redirectTo: "/",
                redirect: true,
              })
            }
          >
            Sign Out
          </button>
        </main>
      )}
    </Show>
  );
};

export default Protected;
