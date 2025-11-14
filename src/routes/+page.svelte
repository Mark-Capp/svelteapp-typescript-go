<script lang="ts">
    import { type PageData } from "./$types";

    const { data } : { data: PageData } = $props();
</script>


<div class="container">

    {#each data.items as item}
    <div class="card mb-2 mt-2">
        <div class="card-body">
            <h5 class="card-title">{item.title}</h5>
            <p class="card-text"></p>
        </div>
    </div>
    {/each}

    <div class="card">
        <h5 class="card-header">Add Item</h5>
        <div class="card-body">
            <form method="POST" action="?/addItem">
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

    <div class="card">
        <h5 class="card-header">Add Tag</h5>
        <div class="card-body">
            <form method="POST" action="?/addTag">
                <div class="form-group row">
                    <label for="colFormLabel" class="col-sm-2 col-form-label">Tag</label>
                    <div class="col-sm-10">
                        <input type="text" class="form-control" name="tag" id="colFormLabel">
                    </div>
                </div>
                <input type="submit" class="btn btn-primary mt-3" value="Add Tag">
            </form>
        </div>
    </div>
</div>