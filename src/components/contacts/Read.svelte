<script lang="ts">
  import { onMount } from "svelte";
  import type { ContactCreateRes as Contact } from "~/pkg/sql/q";
  import Layout from "../Layout.svelte";
  import Header from "./Header.svelte";
  import { Icon } from "@steeze-ui/svelte-icon";
  import { ChatBubbleLeftRight, Envelope } from "@steeze-ui/heroicons";

  let contact = $state<Contact>({
    id: 0,
    name: "",
  });

  let topic = "apps";

  let subject = $derived(`Noah <> ${contact.name} ${topic}`);
  let body = $derived(`Hey ${contact.name}, lets chat more about ${topic}`);

  onMount(async () => {
    const p = new URLSearchParams(window.location.search);
    const res = await fetch(`/api/contacts/${p.get("id")}`);
    contact = await res.json();
  });
</script>

<Layout>
  {#snippet header()}
    <Header />
  {/snippet}

  <div class="card w-96 bg-base-100 card-xl shadow-sm">
    <div class="card-body space-y-6">
      <h2 class="card-title">{contact.name}</h2>
      <a
        class="btn"
        href="mailto:{contact.email}?body={body}&subject={subject}"
      >
        <Icon src={Envelope} class="size-5" /> {contact.email}
      </a>

      <a class="btn" href="sms:{contact.phone}?body={body}">
        <Icon src={ChatBubbleLeftRight} class="size-5" /> {contact.phone}
      </a>

      <div class="justify-end card-actions">
        <a class="btn btn-warning" href="#/contacts">Go Back</a>
        <a class="btn btn-primary" href="?id={contact.id}#/contacts/update"
        >Update</a>
      </div>
    </div>
  </div>
</Layout>
