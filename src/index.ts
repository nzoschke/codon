import { mount } from "svelte";
import Router from "~/components/Router.svelte";
import { routes } from "~/routes";

const target = document.createElement("span");
document.body.appendChild(target);

mount(Router, {
  props: { routes },
  target,
});
