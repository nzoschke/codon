import List from "~/src/components/contacts/List.svelte";
import Create from "~/src/components/contacts/Create.svelte";
import Read from "~/src/components/contacts/Read.svelte";
import Update from "~/src/components/contacts/Update.svelte";
import { type Route } from "~/src/components/Router.svelte";

export default [
  {
    component: Create,
    hash: "#/contacts/create",
  },
  {
    component: List,
    hash: "#/contacts",
  },
  {
    component: Read,
    hash: "#/contacts/read",
  },
  {
    component: Update,
    hash: "#/contacts/update",
  },
] satisfies Route[];
