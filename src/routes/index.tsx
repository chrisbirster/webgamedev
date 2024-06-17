import { type VoidComponent, Suspense, Show } from "solid-js";
import { createSession, signOut, signIn } from "@solid-mediakit/auth/client";
import { helloQuery } from "~/server/hello/hello.queries";

const Home: VoidComponent = () => {
  const hello = helloQuery(() => ({
    hello: "JDev",
  }));

  return (
    <main class="flex min-h-screen flex-col items-center justify-center bg-gradient-to-b from-[#026d56] to-[#152a2c]">
      <div class="flex mx-auto w-1/2 bg-red-300">
        <ul class="flex gap-4">
          <li>
            <a href="/" class="">
              index
            </a>
          </li>
          <li>
            <a href="/protected" class="">
              Protected
            </a>
          </li>
        </ul>
        <p class="text-2xl bg-green-500">{hello.data}</p>
      </div>
      <div class="container flex flex-col items-center justify-center gap-12 px-4 py-16 ">
        <h1 class="text-4xl text-white">Welcome to WebGameDev</h1>
        <p class="text-xl text-white">Learn about the web</p>
        <Suspense>
          <AuthShowcase />
        </Suspense>
      </div>
    </main>
  );
};

export default Home;

const AuthShowcase: VoidComponent = () => {
  const session = createSession();
  return (
    <div class="flex flex-col items-center justify-center gap-4">
      <Show
        when={session()}
        fallback={
          <button
            onClick={() => signIn("discord", { redirectTo: "/" })}
            class="rounded-full bg-white/10 px-10 py-3 font-semibold text-white no-underline transition hover:bg-white/20"
          >
            Sign in
          </button>
        }
      >
        <span class="text-xl text-white">Welcome {session()?.user?.name}</span>
        <button
          onClick={() => signOut({ redirectTo: "/" })}
          class="rounded-full bg-white/10 px-10 py-3 font-semibold text-white no-underline transition hover:bg-white/20"
        >
          Sign out
        </button>
      </Show>
    </div>
  );
};
