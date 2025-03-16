<script lang="ts">
  import Layout from "~/components/Layout.svelte";
</script>

<Layout>
  {#snippet header()}
    header
  {/snippet}

  {#snippet asideL()}
    aside left
  {/snippet}

  {#snippet main()}
    <p class="h-[512px] bg-primary p-4">Paragraph 1</p>
    <p class="h-[512px] bg-primary p-4">Paragraph 2</p>
    <p class="h-[512px] bg-primary p-4">Paragraph 3</p>
  {/snippet}

  {#snippet asideR()}
    aside right
  {/snippet}

  {#snippet footer()}
    footer
  {/snippet}
</Layout>
