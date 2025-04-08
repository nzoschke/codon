<script lang="ts">
  import Layout from "~/src/components/Layout.svelte";
  import Header from "~/src/components/Header.svelte";
  import Footer from "~/src/components/Footer.svelte";
</script>

<Layout>
  {#snippet header()}
    <Header />
  {/snippet}

  {#snippet asideL()}
    <div class="">aside left</div>
  {/snippet}

  <div class="text-primary-content">
    <p class="h-[512px] bg-primary p-4 border-2">Paragraph 1</p>
    <p class="h-[512px] bg-primary p-4">Paragraph 2</p>
    <p class="h-[512px] bg-primary p-4">Paragraph 3</p>
  </div>

  {#snippet asideR()}
    <div class="">aside right</div>
  {/snippet}

  {#snippet footer()}
    <Footer />
  {/snippet}
</Layout>
