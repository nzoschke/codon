import Home from "~/components/Home.svelte";
import { type Route } from "~/components/Router.svelte";

export const routes: Route[] = [
  {
    component: Home,
    hash: "#/",
  },
];
