<script lang="ts">
  import { onMount } from "svelte";
  import type { components, paths } from "~/src/schema";
  import Layout from "../Layout.svelte";
  import { Icon } from "@steeze-ui/svelte-icon";
  import { ChatBubbleLeftRight, Envelope } from "@steeze-ui/heroicons";
  import createClient from "openapi-fetch";

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

  let topic = "apps";

  let subject = $derived(`Chat about ${topic}`);
  let body = $derived(`Hey ${contact.name}, lets chat more about ${topic}`);

  const del = async () => {
    const res = await client.DELETE("/api/contacts/{id}", {
      params: {
        path: {
          id: parseInt(
            new URLSearchParams(location.search).get("id") ?? "0",
          ),
        },
      },
    });
    if (!res.error) window.location.href = "#/contacts";
  };

  onMount(async () => {
    const res = await client.GET("/api/contacts/{id}", {
      params: {
        path: {
          id: parseInt(
            new URLSearchParams(location.search).get("id") ?? "0",
          ),
        },
      },
    });
    if (res.data) contact = res.data;
  });
</script>

<Layout>
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
