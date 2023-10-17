<script lang="ts">
import { GoRestClient } from "$lib/api/goFiber";
import Button from "$components/ui/Button.svelte";
import type {Printer} from "$lib/interfaces"
import DataTables from "$components/ui/dataTables.svelte";
import { onMount } from "svelte";

let api = new GoRestClient(true)

let printersTable;
let usersTable;

onMount(() => {
    // Make the API call right when the component mounts
    getAllPrinter();
});

// live check
let healthCheckPromise;

function healthCheck() {
    healthCheckPromise = api.get("ancillary/version/");
}

//printer
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


function postPrinter() {
    postPrinterPromise = api.post("api/printer", printerTestPayload);
}

function getAllPrinter() {
    getAllPrinterPromise = api.get("api/printer");
}
//TODO: add the option, when clicking on a row, to fetch that data and show clearly

//user
let postNewUserPromise;
let getAllUserPromise;
let userTestPayload: any = {
    username: "bob",
    email: "bob@printer.com",
    password: "superDooberSecretLMAO"
}
function postNewUser() {
    postPrinterPromise = api.post("api/user", userTestPayload);
}

function getAllUser() {
    getAllPrinterPromise = api.get("api/user");
}
</script>


<Button callbackFn={() => healthCheck()}>Get api version</Button>
<Button callbackFn={() => postPrinter()}>Add default printer</Button>
<Button callbackFn={() => getAllPrinter()}>Refresh printer list</Button>

<Button callbackFn={() => postNewUser()}>Add default user</Button>
<Button callbackFn={() => getAllUser()}>Refresh user list</Button>

<!-- <p>Use the onMount and whatnot</p> -->

{#if healthCheckPromise}
    {#await healthCheckPromise then health}
        {#each Object.entries(health) as [key, value]}
            <p>{key}: {value}</p>
        {/each}
    {:catch error}
        <p style="color: red">{error.message}</p>
    {/await}
{:else}
    <p>...waiting</p>
{/if}


{#if getAllPrinterPromise}
    {#await getAllPrinterPromise then printers}
        <DataTables data={printers.data}/>
    {:catch error}
        <p style="color: red">{error.message}</p>
    {/await}
{/if}

<!-- it reuses the same table as above -->
{#if getAllUserPromise}
    {#await getAllUserPromise then users}
        <DataTables data={users.data}/>
    {:catch error}
        <p style="color: red">{error.message}</p>
    {/await}
{/if}

<hr class="border-t-2 border-red-500 w-1/2 mx-auto my-8">
