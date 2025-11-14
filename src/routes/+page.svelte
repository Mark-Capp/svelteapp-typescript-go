<script lang="ts">
    import { type PageData } from "./$types";

    const { data } : { data: PageData } = $props();
</script>


<div class="container">

    {#each data.items as item}
    <div class="card mb-5">
        <div class="card-body">
            <h5 class="card-title">{item.name}</h5>
            <p class="card-text">With supporting text below as a natural lead-in to additional content.</p>
        </div>
    </div>
    {/each}

    <div class="card">
        <h5 class="card-header">Add Item</h5>
        <div class="card-body">
            <form method="POST">
                <div class="form-group row">
                    <label for="colFormLabel" class="col-sm-2 col-form-label">Item</label>
                    <div class="col-sm-10">
                        <input type="text" class="form-control" name="item" id="colFormLabel">
                    </div>
                </div>
                <input type="submit" class="btn btn-primary mt-3" value="Add Item">
            </form>
        </div>
    </div>
</div>