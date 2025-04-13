import { mount } from "svelte";
import Router from "~/src/components/Router.svelte";
import routes from "~/src/routes";

const target = document.createElement("span");
document.body.appendChild(target);

mount(Router, {
  props: { routes },
  target,
});
