<script lang="ts">
  import { onMount } from "svelte";
  import type { ContactCreateRes as Contact } from "~/pkg/sql/q";
  import Layout from "../Layout.svelte";
  import Header from "./Header.svelte";

  onMount(async () => {
  });
</script>

<Layout>
  {#snippet header()}
    <Header />
  {/snippet}

  <div class="card w-96 bg-base-100 card-xl shadow-sm">
    <form method="POST" action="/api/contacts">
      <div class="card-body">
        <fieldset class="fieldset">
          <legend class="fieldset-legend">What is your name?</legend>
          <input
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
          <a class="btn btn-warning" href="#/contacts">Cancel</a>
          <button class="btn btn-primary" type="submit">Create</button>
        </div>
      </div>
    </form>
  </div>
</Layout>
