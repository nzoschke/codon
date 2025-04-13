import Counter from "~/src/components/Counter.svelte";
import Doc from "~/src/components/Doc.svelte";
import Gallery from "~/src/components/Gallery.svelte";
import Home from "~/src/components/Home.svelte";
import { type Route } from "~/src/components/Router.svelte";
import contacts from "~/src/components/contacts/routes";

export default [
  ...contacts,
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
  {
    component: Gallery,
    hash: "#/gallery",
  },
] satisfies Route[];
