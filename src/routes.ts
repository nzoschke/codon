import Home from "~/src/components/Home.svelte";
import Doc from "~/src/components/Doc.svelte";
import { type Route } from "~/src/components/Router.svelte";

export const routes: Route[] = [
  {
    component: Home,
    hash: "#/",
  },
  {
    component: Doc,
    hash: "#/doc",
  },
];
