<script lang="ts">
  import Layout from "../Layout.svelte";
  import Form from "./Form.svelte";

  const id = +(new URLSearchParams(location.search).get("id") ?? "0");
</script>

<Layout>
  <div class="card w-96 bg-base-100 card-xl shadow-sm">
    <Form {id} />
  </div>
</Layout>
