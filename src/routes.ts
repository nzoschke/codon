import Home from "~/components/Home.svelte";
import Doc from "~/components/Doc.svelte";
import { type Route } from "~/components/Router.svelte";

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
