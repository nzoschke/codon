<script lang="ts">
  import { Icon } from "@steeze-ui/svelte-icon";
  import { BookOpen, Home } from "@steeze-ui/heroicons";
</script>

<a class="btn btn-sm" href="#/">
  <Icon src={Home} theme="solid" class="color-gray-900 h-5" />
  Home
</a>

<a class="btn btn-sm" href="#/doc">
  <Icon src={BookOpen} theme="solid" class="color-gray-900 h-5" />
  Docs
</a>
