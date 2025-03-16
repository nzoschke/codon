import Counter from "~/components/Counter.svelte";
import { type Route } from "~/components/Router.svelte";

export const routes: Route[] = [
  {
    component: Counter,
    hash: "#/",
  },
];
