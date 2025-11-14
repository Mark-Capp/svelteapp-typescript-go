<script lang="ts">
    import { slide } from "svelte/transition";
    import { type PageData } from "./$types";

    let { data }: { data: PageData } = $props();

    let films = $state(data.films);
</script>


<div>
  <div class="container-fluid">
    <div class="row">
        {#each films as film}
        <div transition:slide={{ duration: 500 }}>
            <div class="card m-1 p-1 bg-light">
                <h4 class="card-title">
                    {film.title} 
                </h4>
                <h6 class="card-subtitle mb-2 text-muted">
                    ({film.year}) - {film.genre} <br/>
                    Directed by : {film.director}
                </h6>

                <div class="card-text bg-white p-1">
                    {film.synopsis}
                </div>
            </div>
        </div>
        {/each}
    </div>
  </div>
</div>