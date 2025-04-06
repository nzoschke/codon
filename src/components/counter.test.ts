import { Counter } from "./counter.svelte.ts";
import { expect, test } from "bun:test";

// https://svelte.dev/docs/svelte/testing
test("Inc", () => {
  const c = Counter();
  expect(c.count).toEqual(0);

  c.inc();
  expect(c.count).toEqual(1);
});
