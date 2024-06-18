import { type VoidComponent } from "solid-js";
import { clientOnly } from "@solidjs/start";
import { Button } from "~/components/ui/button";

const PulseC = clientOnly(() => import("~/components/pulse"));

const Pulse: VoidComponent = () => {
  return (
    <main class="flex min-h-screen flex-col items-center justify-center bg-gradient-to-b from-[#026d56] to-[#152a2c]">
      <PulseC />
      <Button />
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
      </div>
      <div class="container flex flex-col items-center justify-center gap-12 px-4 py-16 ">
        <h1 class="text-4xl text-white">Welcome to WebGameDev</h1>
        <p class="text-xl text-white">Learn about the web</p>
        <canvas id="canvas" class="w-96 h-96 bg-white"></canvas>
      </div>
    </main>
  );
};

export default Pulse;
