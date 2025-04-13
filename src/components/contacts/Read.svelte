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

  const del = async () => {
    const res = await fetch(`/api/contacts/${contact.id}`, {
      method: "DELETE",
    });
    if (res.status == 200) {
      window.location.href = "#/contacts";
    }
  };

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
        class="btn btn-primary"
        href="mailto:{contact.email}?body={body}&subject={subject}"
      >
        <Icon src={Envelope} class="size-5" /> {contact.email}
      </a>

      <a class="btn btn-primary" href="sms:{contact.phone}?body={body}">
        <Icon src={ChatBubbleLeftRight} class="size-5" /> {contact.phone}
      </a>

      <div class="justify-end card-actions">
        <a class="btn btn-warning btn-soft" href="#/contacts">Go Back</a>
        <a
          class="btn btn-success btn-soft"
          href="?id={contact.id}#/contacts/update"
        >Update</a>
      </div>
    </div>
  </div>

  <button
    class="btn btn-ghost btn-error btn-xs btn-soft italic"
    onclick={() => {
      const el = document.getElementById("modal") as HTMLDialogElement;
      el.showModal();
    }}
  >
    Delete
  </button>

  <dialog id="modal" class="modal modal-bottom sm:modal-middle">
    <div class="modal-box">
      <h3 class="text-lg font-bold">Are you sure?</h3>
      <p class="py-4">This will delete everything for <b>{contact.name}</b>.</p>
      <div class="modal-action">
        <form method="dialog">
          <button class="btn btn-warning btn-soft">Cancel</button>
          <button class="btn btn-error" onclick={del}>Delete</button>
        </form>
      </div>
    </div>
  </dialog>
</Layout>
