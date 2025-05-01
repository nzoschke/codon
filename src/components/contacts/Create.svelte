<script lang="ts">
  import createClient from "openapi-fetch";
  import type { components, paths } from "~/src/schema";
  import Layout from "../Layout.svelte";
  import Form from "./Form.svelte";

  type Contact = components["schemas"]["Contact"];
  const client = createClient<paths>({});

  let contact = $state<Contact>({
    created_at: "",
    email: "",
    id: 0,
    info: {
      age: 0,
    },
    name: "",
    phone: "",
    updated_at: "",
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
