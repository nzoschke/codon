import { serve } from "bun";
import homepage from "~/index.html";

const server = serve({
  routes: {
    "/": homepage,
  },

  development: true,
});

console.log(`Listening on ${server.url}`);
