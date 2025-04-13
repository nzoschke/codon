<script lang="ts">
  import { onMount } from "svelte";
  import type { Contact } from "~/pkg/sql";
  import Layout from "../Layout.svelte";

  let contact = $state<Contact>({
    id: 0,
    name: "",
    created_at: "",
    email: "",
    meta: {},
    phone: "",
    updated_at: "",
  });

  onMount(async () => {
    const p = new URLSearchParams(window.location.search);
    const res = await fetch(`/api/contacts/${p.get("id")}`);
    contact = await res.json();
  });
</script>

<Layout>
  <div class="card w-96 bg-base-100 card-xl shadow-sm">
    <form method="POST" action="/api/contacts/{contact.id}">
      <div class="card-body">
        <fieldset class="fieldset">
          <legend class="fieldset-legend">What is your name?</legend>
          <input
            bind:value={contact.name}
            class="input validator"
            name="name"
            placeholder="Joe Cool"
            required
            type="text"
          />
          <p class="fieldset-label validator-hint">Name is required</p>
        </fieldset>

        <fieldset class="fieldset">
          <legend class="fieldset-legend">What is your email?</legend>
          <input
            bind:value={contact.email}
            class="input validator"
            name="email"
            placeholder="mail@site.com"
            required
            type="email"
          />
          <p class="fieldset-label validator-hint">Email is invalid</p>
        </fieldset>

        <fieldset class="fieldset">
          <legend class="fieldset-legend">What is your phone?</legend>
          <input
            bind:value={contact.phone}
            class="tabular-nums input validator"
            minlength="10"
            maxlength="10"
            name="phone"
            placeholder="4445551212"
            pattern="[0-9]*"
            required
            title="Must be 10 digits"
            type="tel"
          />
          <p class="fieldset-label validator-hint">Phone must be 10 digits</p>
        </fieldset>

        <div class="justify-end card-actions">
          <a
            class="btn btn-warning btn-soft"
            href="?id={contact.id}#/contacts/read"
          >Cancel</a>
          <button class="btn btn-success" type="submit">Update</button>
        </div>
      </div>
    </form>
  </div>
</Layout>
