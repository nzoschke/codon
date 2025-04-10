<script>
  import Layout from "./Layout.svelte";
</script>

<Layout class="flex flex-col items-center justify-center w-full">
  <div class="collapse bg-base-100 border border-base-300">
    <input type="radio" name="my-accordion-1" checked="checked" />
    <div class="collapse-title font-semibold">How do I create an account?</div>
    <div class="collapse-content text-sm">
      Click the "Sign Up" button in the top right corner and follow the
      registration process.
    </div>
  </div>

  <div class="collapse bg-base-100 border border-base-300">
    <input type="radio" name="my-accordion-1" />
    <div class="collapse-title font-semibold">
      I forgot my password. What should I do?
    </div>
    <div class="collapse-content text-sm">
      Click on "Forgot Password" on the login page and follow the instructions
      sent to your email.
    </div>
  </div>

  <div class="collapse bg-base-100 border border-base-300">
    <input type="radio" name="my-accordion-1" />
    <div class="collapse-title font-semibold">
      How do I update my profile information?
    </div>
    <div class="collapse-content text-sm">
      Go to "My Account" settings and select "Edit Profile" to make changes.
    </div>
  </div>

  <div role="alert" class="alert">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      fill="none"
      viewBox="0 0 24 24"
      class="stroke-info h-6 w-6 shrink-0"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
      >
      </path>
    </svg>
    <span>12 unread messages. Tap to see.</span>
  </div>

  <div class="avatar">
    <div class="w-24 rounded">
      <img
        src="https://img.daisyui.com/images/stock/photo-1534528741775-53994a69daeb.webp"
      />
    </div>
  </div>

  <span class="badge">Badge</span>

  <div class="breadcrumbs text-sm">
    <ul>
      <li><a>Home</a></li>
      <li><a>Documents</a></li>
      <li>Add Document</li>
    </ul>
  </div>

  <button class="btn btn-xs">Xsmall</button>
  <button class="btn btn-sm">Small</button>
  <button class="btn">Medium</button>
  <button class="btn btn-lg">Large</button>
  <button class="btn btn-xl">Xlarge</button>

  <div class="card bg-base-100 w-96 shadow-sm">
    <figure>
      <img
        src="https://img.daisyui.com/images/stock/photo-1606107557195-0e29a4b5b4aa.webp"
        alt="Shoes"
      />
    </figure>
    <div class="card-body">
      <h2 class="card-title">Card Title</h2>
      <p>
        A card component has a figure, a body part, and inside body there are
        title and actions parts
      </p>
      <div class="card-actions justify-end">
        <button class="btn btn-primary">Buy Now</button>
      </div>
    </div>
  </div>

  <div class="carousel rounded-box">
    <div class="carousel-item">
      <img
        src="https://img.daisyui.com/images/stock/photo-1559703248-dcaaec9fab78.webp"
        alt="Burger"
      />
    </div>
    <div class="carousel-item">
      <img
        src="https://img.daisyui.com/images/stock/photo-1565098772267-60af42b81ef2.webp"
        alt="Burger"
      />
    </div>
    <div class="carousel-item">
      <img
        src="https://img.daisyui.com/images/stock/photo-1572635148818-ef6fd45eb394.webp"
        alt="Burger"
      />
    </div>
    <div class="carousel-item">
      <img
        src="https://img.daisyui.com/images/stock/photo-1494253109108-2e30c049369b.webp"
        alt="Burger"
      />
    </div>
    <div class="carousel-item">
      <img
        src="https://img.daisyui.com/images/stock/photo-1550258987-190a2d41a8ba.webp"
        alt="Burger"
      />
    </div>
    <div class="carousel-item">
      <img
        src="https://img.daisyui.com/images/stock/photo-1559181567-c3190ca9959b.webp"
        alt="Burger"
      />
    </div>
    <div class="carousel-item">
      <img
        src="https://img.daisyui.com/images/stock/photo-1601004890684-d8cbf643f5f2.webp"
        alt="Burger"
      />
    </div>
  </div>

  <div class="chat chat-start">
    <div class="chat-image avatar">
      <div class="w-10 rounded-full">
        <img
          alt="Tailwind CSS chat bubble component"
          src="https://img.daisyui.com/images/stock/photo-1534528741775-53994a69daeb.webp"
        />
      </div>
    </div>
    <div class="chat-bubble">
      It was said that you would, destroy the Sith, not join them.
    </div>
  </div>
  <div class="chat chat-start">
    <div class="chat-image avatar">
      <div class="w-10 rounded-full">
        <img
          alt="Tailwind CSS chat bubble component"
          src="https://img.daisyui.com/images/stock/photo-1534528741775-53994a69daeb.webp"
        />
      </div>
    </div>
    <div class="chat-bubble">
      It was you who would bring balance to the Force
    </div>
  </div>
  <div class="chat chat-start">
    <div class="chat-image avatar">
      <div class="w-10 rounded-full">
        <img
          alt="Tailwind CSS chat bubble component"
          src="https://img.daisyui.com/images/stock/photo-1534528741775-53994a69daeb.webp"
        />
      </div>
    </div>
    <div class="chat-bubble">Not leave it in Darkness</div>
  </div>

  <input type="checkbox" checked="checked" class="checkbox" />

  <div tabindex="0" class="collapse bg-base-100 border-base-300 border">
    <div class="collapse-title font-semibold">How do I create an account?</div>
    <div class="collapse-content text-sm">
      Click the "Sign Up" button in the top right corner and follow the
      registration process.
    </div>
  </div>

  <span class="countdown">
    <span style="--value: 59" aria-live="polite" aria-label="59">59</span>
  </span>

  <figure class="diff aspect-16/9" tabindex="0">
    <div class="diff-item-1" role="img">
      <img
        alt="daisy"
        src="https://img.daisyui.com/images/stock/photo-1560717789-0ac7c58ac90a.webp"
      />
    </div>
    <div class="diff-item-2" role="img" tabindex="0">
      <img
        alt="daisy"
        src="https://img.daisyui.com/images/stock/photo-1560717789-0ac7c58ac90a-blur.webp"
      />
    </div>
    <div class="diff-resizer"></div>
  </figure>

  <div class="flex w-full flex-col">
    <div class="card bg-base-300 rounded-box grid h-20 place-items-center">
      content
    </div>
    <div class="divider">OR</div>
    <div class="card bg-base-300 rounded-box grid h-20 place-items-center">
      content
    </div>
  </div>

  <details class="dropdown">
    <summary class="btn m-1">open or close</summary>
    <ul
      class="menu dropdown-content bg-base-100 rounded-box z-1 w-52 p-2 shadow-sm"
    >
      <li><a>Item 1</a></li>
      <li><a>Item 2</a></li>
    </ul>
  </details>

  <fieldset class="fieldset">
    <legend class="fieldset-legend">Page title</legend>
    <input type="text" class="input" placeholder="My awesome page" />
    <p class="fieldset-label">You can edit page title later on from settings</p>
  </fieldset>

  <input type="file" class="file-input" />

  <form class="filter">
    <input class="btn btn-square" type="reset" value="×" />
    <input class="btn" type="radio" name="frameworks" aria-label="Svelte" />
    <input class="btn" type="radio" name="frameworks" aria-label="Vue" />
    <input class="btn" type="radio" name="frameworks" aria-label="React" />
  </form>

  <div class="hero bg-base-200 min-h-screen">
    <div class="hero-content text-center">
      <div class="max-w-md">
        <h1 class="text-5xl font-bold">Hello there</h1>
        <p class="py-6">
          Provident cupiditate voluptatem et in. Quaerat fugiat ut assumenda
          excepturi exercitationem quasi. In deleniti eaque aut repudiandae et a
          id nisi.
        </p>
        <button class="btn btn-primary">Get Started</button>
      </div>
    </div>
  </div>

  <div class="indicator">
    <span class="indicator-item status status-success"></span>
    <div class="bg-base-300 grid h-32 w-32 place-items-center">content</div>
  </div>

  <input type="text" placeholder="Type here" class="input" />

  <div class="join">
    <button class="btn join-item">Button</button>
    <button class="btn join-item">Button</button>
    <button class="btn join-item">Button</button>
  </div>

  <kbd class="kbd">K</kbd>

  <label class="input">
    <span class="label">https://</span>
    <input type="text" placeholder="URL" />
  </label>

  <a class="link">Click me</a>

  <ul class="list bg-base-100 rounded-box shadow-md">
    <li class="p-4 pb-2 text-xs opacity-60 tracking-wide">
      Most played songs this week
    </li>

    <li class="list-row">
      <div>
        <img
          class="size-10 rounded-box"
          src="https://img.daisyui.com/images/profile/demo/1@94.webp"
        />
      </div>
      <div>
        <div>Dio Lupa</div>
        <div class="text-xs uppercase font-semibold opacity-60">
          Remaining Reason
        </div>
      </div>
      <button class="btn btn-square btn-ghost">
        <svg
          class="size-[1.2em]"
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 24 24"
        >
          <g
            stroke-linejoin="round"
            stroke-linecap="round"
            stroke-width="2"
            fill="none"
            stroke="currentColor"
          >
            <path d="M6 3L20 12 6 21 6 3z"></path>
          </g>
        </svg>
      </button>
      <button class="btn btn-square btn-ghost">
        <svg
          class="size-[1.2em]"
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 24 24"
        >
          <g
            stroke-linejoin="round"
            stroke-linecap="round"
            stroke-width="2"
            fill="none"
            stroke="currentColor"
          >
            <path
              d="M19 14c1.49-1.46 3-3.21 3-5.5A5.5 5.5 0 0 0 16.5 3c-1.76 0-3 .5-4.5 2-1.5-1.5-2.74-2-4.5-2A5.5 5.5 0 0 0 2 8.5c0 2.3 1.5 4.05 3 5.5l7 7Z"
            >
            </path>
          </g>
        </svg>
      </button>
    </li>

    <li class="list-row">
      <div>
        <img
          class="size-10 rounded-box"
          src="https://img.daisyui.com/images/profile/demo/4@94.webp"
        />
      </div>
      <div>
        <div>Ellie Beilish</div>
        <div class="text-xs uppercase font-semibold opacity-60">
          Bears of a fever
        </div>
      </div>
      <button class="btn btn-square btn-ghost">
        <svg
          class="size-[1.2em]"
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 24 24"
        >
          <g
            stroke-linejoin="round"
            stroke-linecap="round"
            stroke-width="2"
            fill="none"
            stroke="currentColor"
          >
            <path d="M6 3L20 12 6 21 6 3z"></path>
          </g>
        </svg>
      </button>
      <button class="btn btn-square btn-ghost">
        <svg
          class="size-[1.2em]"
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 24 24"
        >
          <g
            stroke-linejoin="round"
            stroke-linecap="round"
            stroke-width="2"
            fill="none"
            stroke="currentColor"
          >
            <path
              d="M19 14c1.49-1.46 3-3.21 3-5.5A5.5 5.5 0 0 0 16.5 3c-1.76 0-3 .5-4.5 2-1.5-1.5-2.74-2-4.5-2A5.5 5.5 0 0 0 2 8.5c0 2.3 1.5 4.05 3 5.5l7 7Z"
            >
            </path>
          </g>
        </svg>
      </button>
    </li>

    <li class="list-row">
      <div>
        <img
          class="size-10 rounded-box"
          src="https://img.daisyui.com/images/profile/demo/3@94.webp"
        />
      </div>
      <div>
        <div>Sabrino Gardener</div>
        <div class="text-xs uppercase font-semibold opacity-60">Cappuccino</div>
      </div>
      <button class="btn btn-square btn-ghost">
        <svg
          class="size-[1.2em]"
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 24 24"
        >
          <g
            stroke-linejoin="round"
            stroke-linecap="round"
            stroke-width="2"
            fill="none"
            stroke="currentColor"
          >
            <path d="M6 3L20 12 6 21 6 3z"></path>
          </g>
        </svg>
      </button>
      <button class="btn btn-square btn-ghost">
        <svg
          class="size-[1.2em]"
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 24 24"
        >
          <g
            stroke-linejoin="round"
            stroke-linecap="round"
            stroke-width="2"
            fill="none"
            stroke="currentColor"
          >
            <path
              d="M19 14c1.49-1.46 3-3.21 3-5.5A5.5 5.5 0 0 0 16.5 3c-1.76 0-3 .5-4.5 2-1.5-1.5-2.74-2-4.5-2A5.5 5.5 0 0 0 2 8.5c0 2.3 1.5 4.05 3 5.5l7 7Z"
            >
            </path>
          </g>
        </svg>
      </button>
    </li>
  </ul>

  <span class="loading loading-spinner loading-xs"></span>

  <img
    class="mask mask-squircle"
    src="https://img.daisyui.com/images/stock/photo-1567653418876-5bb0e566e1c2.webp"
  />

  <ul class="menu bg-base-200 rounded-box w-56">
    <li><a>Item 1</a></li>
    <li><a>Item 2</a></li>
    <li><a>Item 3</a></li>
  </ul>

  <div class="mockup-browser border-base-300 border w-full">
    <div class="mockup-browser-toolbar">
      <div class="input">https://daisyui.com</div>
    </div>
    <div class="grid place-content-center border-t border-base-300 h-80">
      Hello!
    </div>
  </div>

  <div class="mockup-code w-full">
    <pre data-prefix="$"><code>npm i daisyui</code></pre>
  </div>

  <div class="mockup-phone">
    <div class="mockup-phone-camera"></div>
    <div class="mockup-phone-display text-white grid place-content-center">
      It's Glowtime.
    </div>
  </div>

  <div class="mockup-window border border-base-300 w-full">
    <div class="grid place-content-center border-t border-base-300 h-80">
      Hello!
    </div>
  </div>

  <button class="btn" onclick={my_modal_1.showModal()}>open modal</button>
  <dialog id="my_modal_1" class="modal">
    <div class="modal-box">
      <h3 class="text-lg font-bold">Hello!</h3>
      <p class="py-4">Press ESC key or click the button below to close</p>
      <div class="modal-action">
        <form method="dialog">
          <!-- if there is a button in form, it will close the modal -->
          <button class="btn">Close</button>
        </form>
      </div>
    </div>
  </dialog>

  <div class="join">
    <button class="join-item btn">1</button>
    <button class="join-item btn btn-active">2</button>
    <button class="join-item btn">3</button>
    <button class="join-item btn">4</button>
  </div>

  <progress class="progress w-56" value="0" max="100"></progress>

  <div
    class="radial-progress"
    style="--value: 70"
    aria-valuenow="70"
    role="progressbar"
  >
    70%
  </div>

  <input type="radio" name="radio-1" class="radio" checked="checked" />

  <input type="range" min="0" max="100" value="40" class="range" />

  <div class="rating">
    <input
      type="radio"
      name="rating-1"
      class="mask mask-star"
      aria-label="1 star"
    />
    <input
      type="radio"
      name="rating-1"
      class="mask mask-star"
      aria-label="2 star"
      checked="checked"
    />
    <input
      type="radio"
      name="rating-1"
      class="mask mask-star"
      aria-label="3 star"
    />
    <input
      type="radio"
      name="rating-1"
      class="mask mask-star"
      aria-label="4 star"
    />
    <input
      type="radio"
      name="rating-1"
      class="mask mask-star"
      aria-label="5 star"
    />
  </div>

  <select class="select">
    <option disabled selected>Pick a color</option>
    <option>Crimson</option>
    <option>Amber</option>
    <option>Velvet</option>
  </select>

  <div class="skeleton h-32 w-32"></div>

  <div class="stack h-20 w-32">
    <div
      class="bg-primary text-primary-content grid place-content-center rounded-box"
    >
      1
    </div>
    <div
      class="bg-accent text-accent-content grid place-content-center rounded-box"
    >
      2
    </div>
    <div
      class="bg-secondary text-secondary-content grid place-content-center rounded-box"
    >
      3
    </div>
  </div>

  <div class="stats shadow">
    <div class="stat">
      <div class="stat-title">Total Page Views</div>
      <div class="stat-value">89,400</div>
      <div class="stat-desc">21% more than last month</div>
    </div>
  </div>

  <span class="status"></span>

  <ul class="steps">
    <li class="step step-primary">Register</li>
    <li class="step step-primary">Choose plan</li>
    <li class="step">Purchase</li>
    <li class="step">Receive Product</li>
  </ul>

  <label class="swap">
    <input type="checkbox" />
    <div class="swap-on">ON</div>
    <div class="swap-off">OFF</div>
  </label>

  <div role="tablist" class="tabs">
    <a role="tab" class="tab">Tab 1</a>
    <a role="tab" class="tab tab-active">Tab 2</a>
    <a role="tab" class="tab">Tab 3</a>
  </div>

  <div class="overflow-x-auto">
    <table class="table">
      <!-- head -->
      <thead>
        <tr>
          <th></th>
          <th>Name</th>
          <th>Job</th>
          <th>Favorite Color</th>
        </tr>
      </thead>
      <tbody>
        <!-- row 1 -->
        <tr>
          <th>1</th>
          <td>Cy Ganderton</td>
          <td>Quality Control Specialist</td>
          <td>Blue</td>
        </tr>
        <!-- row 2 -->
        <tr>
          <th>2</th>
          <td>Hart Hagerty</td>
          <td>Desktop Support Technician</td>
          <td>Purple</td>
        </tr>
        <!-- row 3 -->
        <tr>
          <th>3</th>
          <td>Brice Swyre</td>
          <td>Tax Accountant</td>
          <td>Red</td>
        </tr>
      </tbody>
    </table>
  </div>

  <textarea class="textarea" placeholder="Bio"></textarea>

  <input type="checkbox" value="synthwave" class="toggle theme-controller" />

  <ul class="timeline">
    <li>
      <div class="timeline-start">1984</div>
      <div class="timeline-middle">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 20 20"
          fill="currentColor"
          class="h-5 w-5"
        >
          <path
            fill-rule="evenodd"
            d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.857-9.809a.75.75 0 00-1.214-.882l-3.483 4.79-1.88-1.88a.75.75 0 10-1.06 1.061l2.5 2.5a.75.75 0 001.137-.089l4-5.5z"
            clip-rule="evenodd"
          />
        </svg>
      </div>
      <div class="timeline-end timeline-box">First Macintosh computer</div>
      <hr />
    </li>
    <li>
      <hr />
      <div class="timeline-start">1998</div>
      <div class="timeline-middle">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 20 20"
          fill="currentColor"
          class="h-5 w-5"
        >
          <path
            fill-rule="evenodd"
            d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.857-9.809a.75.75 0 00-1.214-.882l-3.483 4.79-1.88-1.88a.75.75 0 10-1.06 1.061l2.5 2.5a.75.75 0 001.137-.089l4-5.5z"
            clip-rule="evenodd"
          />
        </svg>
      </div>
      <div class="timeline-end timeline-box">iMac</div>
      <hr />
    </li>
    <li>
      <hr />
      <div class="timeline-start">2001</div>
      <div class="timeline-middle">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 20 20"
          fill="currentColor"
          class="h-5 w-5"
        >
          <path
            fill-rule="evenodd"
            d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.857-9.809a.75.75 0 00-1.214-.882l-3.483 4.79-1.88-1.88a.75.75 0 10-1.06 1.061l2.5 2.5a.75.75 0 001.137-.089l4-5.5z"
            clip-rule="evenodd"
          />
        </svg>
      </div>
      <div class="timeline-end timeline-box">iPod</div>
      <hr />
    </li>
    <li>
      <hr />
      <div class="timeline-start">2007</div>
      <div class="timeline-middle">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 20 20"
          fill="currentColor"
          class="h-5 w-5"
        >
          <path
            fill-rule="evenodd"
            d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.857-9.809a.75.75 0 00-1.214-.882l-3.483 4.79-1.88-1.88a.75.75 0 10-1.06 1.061l2.5 2.5a.75.75 0 001.137-.089l4-5.5z"
            clip-rule="evenodd"
          />
        </svg>
      </div>
      <div class="timeline-end timeline-box">iPhone</div>
      <hr />
    </li>
    <li>
      <hr />
      <div class="timeline-start">2015</div>
      <div class="timeline-middle">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 20 20"
          fill="currentColor"
          class="h-5 w-5"
        >
          <path
            fill-rule="evenodd"
            d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.857-9.809a.75.75 0 00-1.214-.882l-3.483 4.79-1.88-1.88a.75.75 0 10-1.06 1.061l2.5 2.5a.75.75 0 001.137-.089l4-5.5z"
            clip-rule="evenodd"
          />
        </svg>
      </div>
      <div class="timeline-end timeline-box">Apple Watch</div>
    </li>
  </ul>

  <div class="toast">
    <div class="alert alert-info">
      <span>New message arrived.</span>
    </div>
  </div>

  <input type="checkbox" checked="checked" class="toggle" />

  <div class="tooltip" data-tip="hello">
    <button class="btn">Hover me</button>
  </div>

  <input
    class="input validator"
    type="email"
    required
    placeholder="mail@site.com"
  />
</Layout>
