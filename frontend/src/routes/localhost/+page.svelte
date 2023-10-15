<script>
import { GoRestClient } from "$lib/api/goFiber";
import Button from "$components/ui/Button.svelte";

let api = new GoRestClient(true)
let statusPromise; 

function healthCheck() {
    statusPromise = api.get("ancillary/version/");
}

</script>


<Button callbackFn={() => console.log("hello there")}>test</Button>
<Button callbackFn={() => healthCheck()}>api call</Button>

{#await statusPromise}
    <p>...waiting</p>
{:then status}
    {#if status === null || status === undefined}
        <p>Status is unavailable.</p>
    {:else}
        {#each Object.entries(status) as [key, value]}
            <p>{key}: {value}</p>
        {/each}
    {/if}
    <!-- <p>The api version is {status}</p> -->
{:catch error}
    <p style="color: red">{error.message}</p>
{/await}
