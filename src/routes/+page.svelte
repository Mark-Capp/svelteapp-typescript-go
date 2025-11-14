<script lang="ts">
    import { invalidateAll } from "$app/navigation";
    import { addItem, addTag } from "$lib/api/services";
    import type { PageData } from "./$types";

    let { data } : { data:PageData} = $props();

    let tagName = $state("");
    let itemTitle = $state("");

    async function submitTag(event: Event) {
        await addTag(tagName);
        invalidateAll();
    }

    async function submitItem(event: Event) {
        await addItem(itemTitle);
        invalidateAll();
    }
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
            <div class="form-group row">
                <label for="colFormLabel" class="col-sm-2 col-form-label">Item</label>
                <div class="col-sm-10">
                    <input type="text" class="form-control" name="item" id="colFormLabel" bind:value={itemTitle}>
                </div>
            </div>
            <input onclick={submitItem} value="Add Item" type="submit" class="btn btn-primary mt-3">
        </div>
    </div>

    <div class="card">
        <h5 class="card-header">Add Tag</h5>
        <div class="card-body">
            <div class="form-group row">
                <label for="colFormLabel" class="col-sm-2 col-form-label">Tag</label>
                <div class="col-sm-10">
                    <input type="text" class="form-control" name="tag" id="colFormLabel" bind:value={tagName}>
                </div>
            </div>
            <input onclick={submitTag} class="btn btn-primary mt-3" value="Add Tag">
        </div>
    </div>
</div>