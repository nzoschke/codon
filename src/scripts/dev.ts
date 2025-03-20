import { serve } from "bun";
import Index from "~/src/index.html";

const server = serve({
  routes: {
    "/": Index,
  },

  development: true,
});

console.log(`Listening on ${server.url}`);
