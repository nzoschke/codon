<script lang="ts">
  let count = $state(0);
</script>

<button
  class="btn btn-sm"
  onclick={() => {
    count += 1;
  }}
>
  clicks: {count}
</button>
