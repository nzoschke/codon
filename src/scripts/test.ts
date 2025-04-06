import { plugin } from "bun";
import { SveltePlugin } from "bun-plugin-svelte";
import { afterEach, beforeEach } from "bun:test";
import { GlobalRegistrator } from "@happy-dom/global-registrator";

beforeEach(async () => {
  await GlobalRegistrator.register();
});

afterEach(async () => {
  await GlobalRegistrator.unregister();
});

plugin(SveltePlugin({ forceSide: "client" }));
