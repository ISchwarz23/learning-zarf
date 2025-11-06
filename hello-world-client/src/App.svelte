<script>
  import Counter from './lib/Counter.svelte'
  
  import { onMount } from "svelte";
  import { getGreeting } from "./lib/api.js";

  let greeting = "Loading...";

  onMount(async () => {
    try {
      greeting = await getGreeting();
    } catch (err) {
      greeting = "Error fetching greeting 😢";
      console.error(err);
    }
  });
</script>

<main>
  <h1>{greeting}</h1>
</main>
