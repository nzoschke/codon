import { mount } from "svelte";
import App from "~/App.svelte";

const target = document.createElement("div");
document.body.appendChild(target);

mount(App, {
  props: {},
  target,
});
