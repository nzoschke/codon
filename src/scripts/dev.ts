import { serve } from "bun";
import { parseArgs } from "util";
import Index from "~/src/index.html";

const { values } = parseArgs({
  args: Bun.argv,
  options: {
    silent: {
      type: "boolean",
    },
  },
  strict: true,
  allowPositionals: true,
});

const server = serve({
  routes: {
    "/": Index,
  },

  development: true,
});

if (!values.silent) {
  console.log(`Listening on ${server.url}`);
}
