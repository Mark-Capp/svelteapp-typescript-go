<script lang="ts">
    import Header from "$lib/components/Header.svelte";
    import { slide } from "svelte/transition";
    import { type PageData } from "./$types";

    let { data }: { data: PageData } = $props();

    let films = $state(data.films);
</script>


<div>
  <Header />
  <div class="container-fluid">
    <div class="row">
        {#each films as film}
            <div transition:slide={{ duration: 500 }}>
            
            <div class="card m-1 p-1 bg-light">
            <h4>
                {film.title} ({film.year}) - {film.genre} <br/>
                Directed by: {film.director}
                
            </h4>
            <div class="card-text bg-white p-1">
                {film.synopsis}
            </div>
            </div>
            </div>
        {/each}
    </div>
  </div>
</div>