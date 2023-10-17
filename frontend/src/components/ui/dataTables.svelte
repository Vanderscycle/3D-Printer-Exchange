<script>
    export let data = [];
//     let columns = data.length > 0 ? Object.keys(data[0]) : [];
    $: columns = data.length > 0 ? Object.keys(data[0]) : [];

</script>

<div class="relative overflow-x-auto">
    <table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
        <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
            <tr>
                {#each columns as column}
                    <th scope="col" class="px-6 py-3">{column}</th>
                {/each}
            </tr>
        </thead>
        <tbody>
            {#each data as row}
                <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                    {#each columns as column}
                        <td>{row[column]}</td>
                    {/each}
                </tr>
            {/each}
        </tbody>
    </table>
</div>
