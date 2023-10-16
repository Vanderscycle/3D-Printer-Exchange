<script>
    export let data = [];

    let columns = data.length > 0 ? Object.keys(data[0]) : [];
</script>

<table border="1">
    <thead>
        <tr>
            {#each columns as column}
                <th>{column}</th>
            {/each}
        </tr>
    </thead>
    <tbody>
        {#each data as row}
            <tr>
                {#each columns as column}
                    <td>{row[column]}</td>
                {/each}
            </tr>
        {/each}
    </tbody>
</table>
