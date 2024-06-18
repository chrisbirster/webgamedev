import { onCleanup } from "solid-js";
import createRAF, { targetFPS } from "@solid-primitives/raf";

const canvas = document.getElementById("canvas");
const ctx = canvas.getContext("2d");
const location = window.document.location;

class Particle {
  constructor() {
    this.radius = 2;
    this.x = Math.random() * canvas.width;
    this.y = Math.random() * canvas.height;
    this.speedX = Math.random() * 0.4 - 0.2;
    this.speedY = Math.random() * 0.4 - 0.2;
    this.color = `hsl(240, 20%, 50%)`;
  }

  draw() {
    ctx.beginPath();
    ctx.arc(this.x, this.y, this.radius, 0, Math.PI * 2);
    ctx.fillStyle = this.color;
    ctx.fill();
  }

  update() {
    this.x += this.speedX;
    this.y += this.speedY;
    if (this.x > canvas.width || this.x < 0) {
      this.speedX *= -1;
    }
    if (this.y > canvas.height || this.y < 0) {
      this.speedY *= -1;
    }
    this.draw();
  }
}

ctx.canvas.width = window.innerWidth;
ctx.canvas.height = window.innerHeight;

const particles = [];

for (let i = 0; i < 100; i++) {
  particles.push(new Particle(canvas));
}

function animate() {
  ctx.clearRect(0, 0, canvas.width, canvas.height);
  particles.forEach((particle) => {
    particle.update();
  });
}

export default function ClientOnlyComponent() {
  const [, start, stop] = createRAF(targetFPS(() => animate(), 60));
  start();

  onCleanup(() => {
    stop();
  });

  console.log("ctx: ", ctx);
  return <div>{location.href}</div>;
}
