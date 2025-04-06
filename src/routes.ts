import Counter from "~/src/components/Counter.svelte";
import Doc from "~/src/components/Doc.svelte";
import Home from "~/src/components/Home.svelte";
import { type Route } from "~/src/components/Router.svelte";

export const routes: Route[] = [
  {
    component: Home,
    hash: "#/",
  },
  {
    component: Counter,
    hash: "#/counter",
  },
  {
    component: Doc,
    hash: "#/doc",
  },
];
