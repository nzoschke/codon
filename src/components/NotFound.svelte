<script lang="ts">
  import Layout from "~/components/Layout.svelte";

  let { hash }: { hash: string } = $props();
</script>

<Layout>
  <div class="grid place-items-center">
    <div class="mockup-code w-96">
      <pre
        data-prefix="1"
        class="bg-error text-error-content"
      >
        <code>Not Found: /{hash}</code></pre>
    </div>
    <a class="underline italic" href="#/">Go home</a>
  </div>
</Layout>
