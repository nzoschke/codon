<script lang="ts">
  import { onMount } from "svelte";
  import type { Contact } from "~/pkg/sql";
  import Layout from "../Layout.svelte";
  import Time from "svelte-time";

  let contacts = $state<Contact[]>();

  onMount(async () => {
    const res = await fetch("/api/contacts");
    contacts = await res.json();
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
