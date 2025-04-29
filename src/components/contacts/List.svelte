<script lang="ts">
  import createClient from "openapi-fetch";
  import { onMount } from "svelte";
  import Time from "svelte-time";
  import type { components, paths } from "~/src/schema";
  import Layout from "../Layout.svelte";

  type Contact = components["schemas"]["Contact"];

  const client = createClient<paths>({});
  let contacts = $state<Contact[]>();

  onMount(async () => {
    const res = await client.GET("/api/contacts");
    contacts = res.data;
  });
</script>

<Layout>
  <div class="toast">
    <a class="btn btn-info" href="#/contacts/create">Create Contact</a>
  </div>

  <div class="overflow-x-auto">
    <table class="table table-zebra">
      <thead>
        <tr>
          <th>Name</th>
          <th>Email</th>
          <th>Phone</th>
          <th>Created</th>
          <th></th>
        </tr>
      </thead>
      <tbody>
        {#if !contacts}
          <tr><td class="text-center italic">None</td></tr>
        {:else}
          {#each contacts as c}
            <tr>
              <th>{c.name}</th>
              <td>{c.email}</td>
              <td>{c.phone}</td>
              <td><Time live relative timestamp={c.created_at} /></td>
              <td>
                <a class="btn btn-primary" href="?id={c.id}#/contacts/read"
                >Say Hi</a>
              </td>
            </tr>
          {/each}
        {/if}
      </tbody>
    </table>
  </div>
</Layout>
