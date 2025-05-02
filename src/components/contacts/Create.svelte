<script lang="ts">
  import createClient from "openapi-fetch";
  import type { components, paths } from "~/src/schema";
  import Layout from "../Layout.svelte";
  import Form from "./Form.svelte";

  type Contact = components["schemas"]["ContactCreateIn"];
  const client = createClient<paths>({});

  let contact = $state<Contact>({
    email: "",
    info: {
      age: 0,
    },
    name: "",
    phone: "",
  });

  const onsubmit = async (e: SubmitEvent) => {
    e.preventDefault();
    const res = await client.POST("/api/contacts", {
      body: contact,
    });
    if (!res.error) {
      window.location.hash = "#/contacts";
    }
  };
</script>

<Layout>
  <div class="card w-96 bg-base-100 card-xl shadow-sm">
    <Form cancel="/#/contacts" bind:contact {onsubmit} />
  </div>
</Layout>
