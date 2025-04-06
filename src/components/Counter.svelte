<script lang="ts">
  import { Counter } from "./counter.svelte.ts";
  let c = Counter();
</script>

<button
  class="btn btn-sm"
  onclick={c.inc}
>
  clicks: {c.count}
</button>
