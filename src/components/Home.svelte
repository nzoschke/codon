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

  <div class="text-base-content space-y-4">
    <div class="hero min-h-[60vh]">
      <div class="hero-content flex-col lg:flex-row">
        <img
          src="https://img.daisyui.com/images/stock/photo-1635805737707-575885ab0820.webp"
          class="max-w-sm rounded-lg shadow-2xl"
          alt="spiderman"
        />
        <div>
          <h1 class="text-5xl font-bold">Hero 1!</h1>
          <p class="py-6">
            Provident cupiditate voluptatem et in. Quaerat fugiat ut assumenda
            excepturi exercitationem quasi. In deleniti eaque aut repudiandae et
            a id nisi.
          </p>
          <button class="btn btn-primary">Get Started</button>
        </div>
      </div>
    </div>

    <div class="flex flex-col items-center">
      <div class="max-w-lg mx-auto text-center">
        <div
          class="text-3xl md:text-5xl font-bold bg-clip-text text-transparent bg-linear-to-r pb-2 from-gray-500 to-blue-600"
        >
          Explore the Features
        </div>
        <div class="mt-4 text-xl font-bold">
          And try them on this <span
            class="underline decoration-secondary decoration-[3px] md:decoration-[4px]"
          >fully functional demo</span>
        </div>
      </div>

      <div class="p-4 flex flex-wrap justify-center gap-4 max-w-4xl">
        {#each [...Array(6)] as _}
          <div class="card bg-base-100 w-64 shadow-sm">
            <figure>
              <img
                src="https://img.daisyui.com/images/stock/photo-1606107557195-0e29a4b5b4aa.webp"
                alt="Shoes"
              />
            </figure>
            <div class="card-body">
              <h2 class="card-title">Card Title</h2>
              <p>
                A card component has a figure, a body part, and inside body
                there are title and actions parts
              </p>
              <div class="card-actions justify-center">
                <button class="btn btn-primary">Buy Now</button>
              </div>
            </div>
          </div>
        {/each}
      </div>
    </div>

    <p class="h-[512px] bg-primary p-4">Paragraph 1</p>
  </div>

  {#snippet asideR()}
    <div class="">aside right</div>
  {/snippet}

  {#snippet footer()}
    <Footer />
  {/snippet}
</Layout>
