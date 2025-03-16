<script lang="ts">
  import type { Snippet } from "svelte";

  interface Props {
    asideL?: Snippet;
    asideR?: Snippet;
    header?: Snippet;
    main: Snippet;
    footer?: Snippet;
  }
  let { asideL, asideR, header, main, footer }: Props = $props();
</script>

<!-- https://www.skeleton.dev/docs/guides/layouts -->
<div class="grid min-h-screen grid-rows-[auto_1fr_auto]">
  {#if header}
    <header
      class="sticky top-0 z-10 bg-base-300/80 text-base-content backdrop-blur-sm p-4"
    >
      {@render header()}
    </header>
  {/if}

  <div class="grid grid-cols-1 md:grid-cols-[auto_1fr_auto]">
    {#if asideL}
      <aside class="bg-base-200 text-base-content p-4">
        {@render asideL()}
      </aside>
    {/if}

    <main class="bg-base-100 text-primary-content p-4 space-y-4">
      {@render main()}
    </main>

    {#if asideR}
      <aside class="bg-base-200 text-base-content p-4">
        {@render asideR()}
      </aside>
    {/if}
  </div>

  {#if footer}
    <footer class="bg-base-300 text-base-content p-4">
      {@render footer()}
    </footer>
  {/if}
</div>
