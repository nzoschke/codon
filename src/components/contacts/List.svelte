<script lang="ts">
  import { onMount } from "svelte";
  import type { ContactCreateRes as Contact } from "~/pkg/sql/q";
  import Layout from "../Layout.svelte";
  import Header from "./Header.svelte";
  import Time from "svelte-time";

  let contacts = $state<Contact[]>();

  onMount(async () => {
    const res = await fetch("/api/contacts");
    contacts = await res.json();
  });
</script>

<Layout>
  {#snippet header()}
    <Header />
  {/snippet}

  <div class="toast">
    <a class="btn btn-accent" href="#/contacts/create">New Contact</a>
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
