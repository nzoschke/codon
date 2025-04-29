<script lang="ts">
  import createClient from "openapi-fetch";
  import type { components, paths } from "~/src/schema";
  import Layout from "../Layout.svelte";
  import Form from "./Form.svelte";
  import { onMount } from "svelte";

  type Contact = components["schemas"]["Contact"];
  const client = createClient<paths>({});

  const id = new URLSearchParams(location.search).get("id") ?? "0";

  let contact = $state<Contact>({});

  onMount(async () => {
    const res = await client.GET("/api/contacts/{id}", {
      params: {
        path: {
          id,
        },
      },
    });
    if (res.data) contact = res.data;
  });

  const onsubmit = async (e: SubmitEvent) => {
    e.preventDefault();
    const res = await client.PUT("/api/contacts/{id}", {
      body: {
        email: contact.email,
        info: contact.info,
        name: contact.name,
        phone: contact.phone,
      },
      params: {
        path: {
          id,
        },
      },
    });
    if (!res.error) {
      window.location.hash = "#/contacts/read";
    }
  };
</script>

<Layout>
  <div class="card w-96 bg-base-100 card-xl shadow-sm">
    <Form cancel="#/contacts/read" bind:contact {onsubmit} />
  </div>
</Layout>
