<script lang="ts">
import { GoRestClient } from "$lib/api/goFiber";
import Button from "$components/ui/Button.svelte";
import type {Printer} from "$lib/interfaces"
import DataTables from "$components/ui/dataTables.svelte";
import { onMount } from "svelte";
let api = new GoRestClient(true)

let healthCheckPromise;
let postPrinterPromise;
let getAllPrinterPromise;

let printerTestPayload: Printer = {
    user: "bob",
    brand: "Ratrig",
    name: "v-minion",
    powerSupply: "24v",
    probe: "pindaProbe",
    board: "btt",
    hotend: "phateus dragonfly BMO",
    extruder: "bmo",
    nozzle: ".4",
    buildVolume: "180cm3",
    mods: "nil"
}

onMount(() => {
    // Make the API call right when the component mounts
    getAllPrinter();
});

function healthCheck() {
    heatlhCheckPromise = api.get("ancillary/version/");
}

function postPrinter() {
    postPrinterPromise = api.post("api/printer", printerTestPayload);
}

function getAllPrinter() {
    getAllPrinterPromise = api.get("api/printer");
}


</script>


<Button callbackFn={() => console.log("hello there")}>test</Button>

<Button callbackFn={() => healthCheck()}>Get api version</Button>
<Button callbackFn={() => postPrinter()}>Add default printer</Button>
<Button callbackFn={() => getAllPrinter()}>Refresh printer list</Button>

<!-- <p>Use the onMount and whatnot</p> -->

{#await healthCheckPromise}
    <p>...waiting</p>
{:then status}
    {#if status === null || status === undefined}
        <p>Status is unavailable.</p>
    {:else}
        {#each Object.entries(status) as [key, value]}
            <p>{key}: {value}</p>
        {/each}
    {/if}
{:catch error}
    <p style="color: red">{error.message}</p>
{/await}

{#if getAllPrinterPromise}
    {#await getAllPrinterPromise then printers}
        <DataTables data={printers.data} />
    {:catch error}
        <p style="color: red">{error.message}</p>
    {/await}
{/if}
